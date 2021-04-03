package impl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Transmission struct {
	login     string
	password  string
	rpcUrl    *url.URL
	csrfToken string
}

type Torrent struct {
	Id          int
	Name        string
	DateCreated time.Time `json:"dateCreated"`
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

func (t *Transmission) AddTorrent(content []byte, dir string) error {
	var responseModel struct {
		Arguments struct {
			TorrentAdded struct {
				Name string `json:"name"`
			} `json:"torrent-added"`
		} `json:"arguments"`
		Result string `json:"result"`
	}
	err := t.call("torrent-add", map[string]string{
		"download-dir": dir,
		"metainfo":     base64.StdEncoding.EncodeToString(content),
	}, &responseModel)
	if err != nil {
		return err
	}

	if responseModel.Result != "success" {
		return errors.New("torrent-add result: " + responseModel.Result)
	}

	return nil
}

func (t *Transmission) GetTorrents() ([]Torrent, error) {
	var responseModel struct {
		Arguments struct {
			Torrents []Torrent
		}
	}
	err := t.call("torrent-get", map[string]interface{}{
		"fields": []string{"id", "name", "dateCreated"},
	}, &responseModel)

	return responseModel.Arguments.Torrents, err
}

func (t *Transmission) call(method string, arguments interface{}, responseModel interface{}) error {
	body, err := json.Marshal(struct {
		Method    string      `json:"method"`
		Arguments interface{} `json:"arguments"`
	}{Method: method, Arguments: arguments})
	if err != nil {
		return err
	}

	responseBytes, err := t.rpcRequest(body)
	if err != nil {
		return err
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
