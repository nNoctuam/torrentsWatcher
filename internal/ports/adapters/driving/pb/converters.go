package pb

import (
	"torrentsWatcher/internal/domain/models"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TorrentToPB(t *models.Torrent) *Torrent {
	return &Torrent{
		Id:         uint32(t.ID),
		Title:      t.Title,
		PageUrl:    t.PageURL,
		FileUrl:    t.FileURL,
		Forum:      t.Forum,
		Author:     t.Author,
		Size:       t.Size,
		Seeders:    t.Seeders,
		CreatedAt:  &timestamp.Timestamp{Seconds: t.CreatedAt.Unix()},
		UpdatedAt:  &timestamp.Timestamp{Seconds: t.UpdatedAt.Unix()},
		UploadedAt: &timestamp.Timestamp{Seconds: t.UploadedAt.Unix()},
	}
}

func TorrentsToPB(torrents []*models.Torrent) (result []*Torrent) {
	for _, t := range torrents {
		result = append(result, TorrentToPB(t))
	}
	return result
}
