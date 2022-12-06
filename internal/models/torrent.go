package models

import (
	"encoding/json"
	"time"
)

type Torrent struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	Title      string     `json:"title"`
	PageURL    string     `json:"page_url" gorm:"unique"`
	FileURL    string     `json:"file_url"`
	File       []byte     `json:"-"`
	Forum      string     `gorm:"-"`
	Author     string     `gorm:"-"`
	Size       uint64     `gorm:"-"`
	Seeders    uint64     `gorm:"-"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UploadedAt time.Time  `json:"uploaded_at"`
	DeletedAt  *time.Time `json:"deleted_at" sql:"index"`
}

func (t *Torrent) MarshalJSON() ([]byte, error) {
	type Alias Torrent
	return json.Marshal(&struct {
		*Alias
		CreatedAt  int64 `json:"created_at"`
		UpdatedAt  int64 `json:"updated_at"`
		UploadedAt int64 `json:"uploaded_at"`
		DeletedAt  int64 `json:"deleted_at"`
	}{
		Alias:      (*Alias)(t),
		CreatedAt:  t.CreatedAt.Unix(),
		UpdatedAt:  t.UpdatedAt.Unix(),
		UploadedAt: t.UploadedAt.Unix(),
		DeletedAt:  t.GetDeletedAt(),
	})
}

func (t *Torrent) GetDeletedAt() int64 {
	if t.DeletedAt == nil {
		return 0
	}
	return t.DeletedAt.Unix()
}

func (t *Torrent) UpdateFrom(updatedTorrent *Torrent) {
	t.Title = updatedTorrent.Title
	t.UploadedAt = updatedTorrent.UploadedAt
	t.FileURL = updatedTorrent.FileURL
}
