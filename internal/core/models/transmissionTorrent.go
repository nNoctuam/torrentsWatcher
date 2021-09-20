package models

import "time"

type TransmissionTorrent struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Hash      string    `json:"hash"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}
