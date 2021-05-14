package torrentclient

import (
	"encoding/json"
	"os"
	"time"
)

type TorrentClient struct {
	autoDownloadDir string
	client          Client
}

type Client interface {
	AddTorrent(content []byte, dir string, paused bool) (Torrent, error)
	UpdateTorrent(url string, content []byte) error
	RemoveTorrents(ids []int, deleteLocalData bool) error
	GetTorrents() ([]Torrent, error)
	Rename(id int, oldPath string, newPath string) error
}

type Torrent struct {
	Id          int
	Name        string
	Hash        string
	Comment     string    `json:"comment"`
	DownloadDir string    `json:"downloadDir"`
	DateCreated time.Time `json:"dateCreated"`
	Labels      []string  `json:"labels"`
}

func (t *Torrent) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Id          int
		Name        string
		Comment     string `json:"comment"`
		DownloadDir string `json:"downloadDir"`
		Hash        string `json:"hashString"`
		DateCreated int64  `json:"dateCreated"`
	}
	torrent := &Alias{}
	err := json.Unmarshal(data, &torrent)
	if err != nil {
		return err
	}
	t.Id = torrent.Id
	t.Name = torrent.Name
	t.Hash = torrent.Hash
	t.Comment = torrent.Comment
	t.DownloadDir = torrent.DownloadDir
	t.DateCreated = time.Unix(torrent.DateCreated, 0)

	return nil
}

func New(autoDownloadDir string, client Client) *TorrentClient {
	return &TorrentClient{
		autoDownloadDir: autoDownloadDir,
		client:          client,
	}
}

func (c *TorrentClient) SaveToAutoDownloadFolder(name string, content []byte) error {
	return os.WriteFile(c.autoDownloadDir+"/"+name, content, 0660)
}

func (c *TorrentClient) AddTorrent(content []byte, dir string, paused bool) (Torrent, error) {
	return c.client.AddTorrent(content, dir, paused)
}

func (c *TorrentClient) UpdateTorrent(url string, content []byte) error {
	return c.client.UpdateTorrent(url, content)
}

func (c *TorrentClient) GetTorrents() ([]Torrent, error) {
	return c.client.GetTorrents()
}

func (c *TorrentClient) Rename(id int, oldPath string, newPath string) error {
	return c.client.Rename(id, oldPath, newPath)
}
