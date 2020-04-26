package models

import "time"

type TorrentInfo struct {
	Title     string
	PageUrl   string
	FileUrl   string
	UpdatedAt time.Time
}
