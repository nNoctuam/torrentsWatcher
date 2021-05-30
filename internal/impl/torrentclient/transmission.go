package torrentclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"torrentsWatcher/internal/core/torrentclient"
)

type Transmission struct {
	autoDownloadDir string
	login           string
	password        string
	rpcUrl          *url.URL
	csrfToken       string
}

func NewTransmission(autoDownloadDir string, rpcUrl string, login string, password string) (*Transmission, error) {
	urlParsed, err := url.Parse(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &Transmission{
		autoDownloadDir: autoDownloadDir,
		login:           login,
		password:        password,
		rpcUrl:          urlParsed,
	}, nil
}

func (t *Transmission) SaveToAutoDownloadFolder(name string, content []byte) error {
	return os.WriteFile(t.autoDownloadDir+"/"+name, content, 0660)
}

func (t *Transmission) AddTorrent(content []byte, dir string, paused bool) (torrent torrentclient.Torrent, err error) {
	var responseModel struct {
		Arguments struct {
			TorrentAdded     torrentclient.Torrent `json:"torrent-added"`
			TorrentDuplicate torrentclient.Torrent `json:"torrent-duplicate"`
		} `json:"arguments"`
		Result string `json:"result"`
	}
	err = t.call("torrent-add", map[string]interface{}{
		"download-dir": dir,
		"metainfo":     base64.StdEncoding.EncodeToString(content),
		"paused":       paused,
	}, &responseModel)
	if err != nil {
		return torrentclient.Torrent{}, err
	}

	if responseModel.Result != "success" {
		return torrentclient.Torrent{}, errors.New("torrent-add result: " + responseModel.Result)
	}

	torrent = responseModel.Arguments.TorrentAdded
	if torrent.Hash == "" {
		torrent.Hash = responseModel.Arguments.TorrentDuplicate.Hash
	}
	if torrent.Name == "" {
		torrent.Name = responseModel.Arguments.TorrentDuplicate.Name
	}

	return torrent, nil
}

func (t *Transmission) GetTorrents() ([]torrentclient.Torrent, error) {
	var responseModel struct {
		Arguments struct {
			Torrents []torrentclient.Torrent `json:"torrents"`
		} `json:"arguments"`
	}
	err := t.call("torrent-get", map[string]interface{}{
		"fields": []string{"id", "name", "hashString", "dateCreated", "comment", "labels", "downloadDir"},
	}, &responseModel)

	return responseModel.Arguments.Torrents, err
}

func (t *Transmission) RemoveTorrents(ids []int, deleteLocalData bool) error {
	var responseModel struct {
		Arguments struct{} `json:"arguments"`
		Result    string   `json:"result"`
	}
	err := t.call("torrent-remove", map[string]interface{}{
		"ids":               ids,
		"delete-local-data": deleteLocalData,
	}, &responseModel)
	if err != nil {
		return err
	}

	if responseModel.Result != "success" {
		return errors.New("torrent-remove result: " + responseModel.Result)
	}
	return err
}

func (t *Transmission) UpdateTorrent(url string, content []byte) error {
	torrents, err := t.GetTorrents()
	if err != nil {
		return err
	}

	for _, oldTorrent := range torrents {
		if oldTorrent.Comment == url ||
			oldTorrent.Comment == strings.Replace(url, "http:", "https:", 1) ||
			oldTorrent.Comment == strings.Replace(url, "https:", "http:", 1) {
			newTorrent, err := t.AddTorrent(content, oldTorrent.DownloadDir, true)
			if err != nil {
				return fmt.Errorf("replace torrent: %w", err)
			}
			err = t.Rename(newTorrent.Id, newTorrent.Name, oldTorrent.Name)
			if err != nil {
				return fmt.Errorf("rename torrent: %w", err)
			}
			err = t.Start([]int{newTorrent.Id})
			if err != nil {
				return fmt.Errorf("start torrent: %w", err)
			}

			_ = t.Verify([]int{newTorrent.Id})

			err = t.RemoveTorrents([]int{oldTorrent.Id}, false)
			if err != nil {
				return fmt.Errorf("delete old torrent: %w", err)
			}

			return nil
		}
	}

	return errors.New("couldn't find old torrent")
}

func (t *Transmission) Rename(id int, oldPath string, newPath string) error {
	var responseModel struct {
		Arguments struct {
			ID   int    `json:"id"`
			Path string `json:"path"`
			Name string `json:"name"`
		} `json:"arguments"`
		Result string `json:"result"`
	}
	err := t.call("torrent-rename-path", map[string]interface{}{
		"ids":  []int{id},
		"path": oldPath,
		"name": newPath,
	}, &responseModel)
	if err != nil {
		return err
	}

	if responseModel.Result != "success" {
		return errors.New("torrent-rename result: " + responseModel.Result)
	}
	return nil
}

func (t *Transmission) Start(ids []int) error {
	var responseModel struct {
		Arguments struct{} `json:"arguments"`
		Result    string   `json:"result"`
	}
	err := t.call("torrent-start", map[string]interface{}{
		"ids": ids,
	}, &responseModel)
	if err != nil {
		return err
	}

	if responseModel.Result != "success" {
		return errors.New("torrent-start result: " + responseModel.Result)
	}
	return nil
}

func (t *Transmission) Verify(ids []int) error {
	var responseModel struct {
		Arguments struct{} `json:"arguments"`
		Result    string   `json:"result"`
	}
	err := t.call("torrent-verify", map[string]interface{}{
		"ids": ids,
	}, &responseModel)
	if err != nil {
		return err
	}

	if responseModel.Result != "success" {
		return errors.New("torrent-verify result: " + responseModel.Result)
	}
	return nil
}

func (t *Transmission) call(method string, arguments interface{}, responseModel interface{}) error {
	body, err := json.Marshal(struct {
		Method    string      `json:"method"`
		Arguments interface{} `json:"arguments"`
	}{Method: method, Arguments: arguments})
	if err != nil {
		return fmt.Errorf("transmission request marshalling: %w", err)
	}

	responseBytes, err := t.rpcRequest(body)
	if err != nil {
		return fmt.Errorf("transmission rpc request: %w", err)
	}
	return json.Unmarshal(responseBytes, responseModel)
}

func (t *Transmission) rpcRequest(body []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", t.rpcUrl.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	auth := make([]byte, base64.StdEncoding.EncodedLen(len([]byte(t.login+":"+t.password))))
	base64.StdEncoding.Encode(auth, []byte(t.login+":"+t.password))
	request.Header.Add("Authorization", "Basic "+string(auth))
	request.Header.Add("Content-Type", "application/json")
	if t.csrfToken != "" {
		request.Header.Add("X-Transmission-Session-Id", t.csrfToken)
	}

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusConflict {
		t.csrfToken = response.Header.Get("X-Transmission-Session-Id")
		request.Header.Add("X-Transmission-Session-Id", t.csrfToken)
		request.Body = io.NopCloser(bytes.NewReader(body))
		response, err = client.Do(request)
		if err != nil {
			return nil, err
		}
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
