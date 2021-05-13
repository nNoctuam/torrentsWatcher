package impl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	"torrentsWatcher/internal/api/torrentclient"
)

type Transmission struct {
	login     string
	password  string
	rpcUrl    *url.URL
	csrfToken string
}

func NewTransmission(rpcUrl string, login string, password string) (*Transmission, error) {
	urlParsed, err := url.Parse(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &Transmission{
		login:    login,
		password: password,
		rpcUrl:   urlParsed,
	}, nil
}

func (t *Transmission) AddTorrent(content []byte, dir string) (torrent torrentclient.Torrent, err error) {
	var responseModel struct {
		Arguments struct {
			TorrentAdded     torrentclient.Torrent `json:"torrent-added"`
			TorrentDuplicate torrentclient.Torrent `json:"torrent-duplicate"`
		} `json:"arguments"`
		Result string `json:"result"`
	}
	err = t.call("torrent-add", map[string]string{
		"download-dir": dir,
		"metainfo":     base64.StdEncoding.EncodeToString(content),
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
		"fields": []string{"id", "name", "hashString", "dateCreated", "comment"},
	}, &responseModel)

	return responseModel.Arguments.Torrents, err
}

func (t *Transmission) RemoveTorrents(ids []int, deleteLocalData bool) error {
	err := t.call("torrent-get", map[string]interface{}{
		"ids":               ids,
		"delete-local-data": deleteLocalData,
	}, nil)

	return err
}

func (c *Transmission) UpdateTorrent(url string, content []byte) error {
	torrents, err := c.GetTorrents()
	if err != nil {
		return err
	}

	for _, oldTorrent := range torrents {
		if oldTorrent.Comment == url {
			err = c.RemoveTorrents([]int{oldTorrent.Id}, false)
			if err != nil {
				return fmt.Errorf("delete old torrent: %w", err)
			}
			newTorrent, err := c.AddTorrent(content, oldTorrent.DownloadDir)
			if err != nil {
				return fmt.Errorf("replace torrent: %w", err)
			}
			err = c.Rename(newTorrent.Id, newTorrent.Name, oldTorrent.Name)
			if err != nil {
				return fmt.Errorf("rename torrent: %w", err)
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

func (c *Transmission) rpcRequest(body []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", c.rpcUrl.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	auth := make([]byte, base64.StdEncoding.EncodedLen(len([]byte(c.login+":"+c.password))))
	base64.StdEncoding.Encode(auth, []byte(c.login+":"+c.password))
	request.Header.Add("Authorization", "Basic "+string(auth))
	request.Header.Add("Content-Type", "application/json")
	if c.csrfToken != "" {
		request.Header.Add("X-Transmission-Session-Id", c.csrfToken)
	}

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusConflict {
		c.csrfToken = response.Header.Get("X-Transmission-Session-Id")
		request.Header.Add("X-Transmission-Session-Id", c.csrfToken)
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
