package models

import (
	"encoding/json"
	"time"
)

type Torrent struct {
	Id         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UploadedAt time.Time  `json:"uploaded_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	Title      string     `json:"title"`
	PageUrl    string     `gorm:"unique" json:"page_url"`
	FileUrl    string     `json:"file_url"`
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
