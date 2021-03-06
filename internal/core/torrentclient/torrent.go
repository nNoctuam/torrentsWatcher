package torrentclient

import (
	"encoding/json"
	"time"
)

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
