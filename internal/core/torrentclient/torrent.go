package torrentclient

import (
	"encoding/json"
	"time"
)

type Torrent struct {
	ID          int
	Name        string        `json:"name"`
	Hash        string        `json:"hash"`
	Comment     string        `json:"comment"`
	DownloadDir string        `json:"downloadDir"`
	DateCreated time.Time     `json:"dateCreated"`
	Labels      []string      `json:"labels"`
	Files       []TorrentFile `json:"files"`
}

type TorrentFile struct {
	BytesCompleted int    `json:"bytesCompleted"`
	Length         int    `json:"length"`
	Name           string `json:"name"`
}

func (t *Torrent) UnmarshalJSON(data []byte) error {
	type Alias struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		Comment     string        `json:"comment"`
		DownloadDir string        `json:"downloadDir"`
		Hash        string        `json:"hashString"`
		DateCreated int64         `json:"dateCreated"`
		Files       []TorrentFile `json:"files"`
	}
	torrent := &Alias{}
	err := json.Unmarshal(data, &torrent)
	if err != nil {
		return err
	}
	t.ID = torrent.ID
	t.Name = torrent.Name
	t.Hash = torrent.Hash
	t.Comment = torrent.Comment
	t.DownloadDir = torrent.DownloadDir
	t.DateCreated = time.Unix(torrent.DateCreated, 0)
	t.Files = torrent.Files

	return nil
}
