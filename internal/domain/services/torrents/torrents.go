package torrents

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"torrentsWatcher/internal/domain/models"
	"torrentsWatcher/internal/ports/adapters/driven"
	"torrentsWatcher/internal/ports/adapters/driven/storage"

	"go.uber.org/zap"
)

type PartToRename struct {
	OldName string
	NewName string
}

type Torrents struct {
	logger          *zap.Logger
	trackers        driven.TrackersAdapter
	torrentsStorage storage.Torrents
	torrentClient   driven.TorrentClient
	downloadFolders map[string]string
	blockViewList   []string
}

func New(
	logger *zap.Logger,
	trackers driven.TrackersAdapter,
	torrentsStorage storage.Torrents,
	torrentClient driven.TorrentClient,
	downloadFolders map[string]string,
	blockViewList []string,
) *Torrents {
	return &Torrents{
		logger:          logger,
		trackers:        trackers,
		torrentsStorage: torrentsStorage,
		torrentClient:   torrentClient,
		downloadFolders: downloadFolders,
		blockViewList:   blockViewList,
	}
}

func (t *Torrents) Add(url string) (*models.Torrent, error) {
	torrent, err := t.trackers.GetTorrentInfo(url)
	if err != nil {
		return nil, err
	}

	_, file, err := t.trackers.DownloadTorrentFile(torrent)
	if err != nil {
		t.logger.Error("Failed to load torrent file", zap.Error(err), zap.String("url", torrent.FileURL))
		return nil, err
	}

	torrent.File = file
	err = t.torrentsStorage.Save(torrent)
	if err != nil {
		t.logger.Error("Failed to save torrent to storage", zap.Error(err))
		return nil, err
	}

	return torrent, nil
}

func (t *Torrents) Delete(id uint) error {
	var torrents []models.Torrent

	err := t.torrentsStorage.Find(&torrents, models.Torrent{
		ID: id,
	})
	if err != nil {
		return err
	}
	if len(torrents) == 0 {
		return errors.New("torrent not found")
	}

	torrent := torrents[0]
	now := time.Now()
	torrent.DeletedAt = &now

	if err = t.torrentsStorage.Save(&torrent); err != nil {
		return fmt.Errorf("failed to update torrent: %w", err)
	}

	return nil
}

func (t *Torrents) Download(url string, folderAlias string) (*driven.Torrent, error) {
	folder, ok := t.downloadFolders[folderAlias]
	t.logger.Debug(
		"folders matching",
		zap.String("folderName", folder),
		zap.String("path", folder),
		zap.Bool("found", ok),
	)

	torrent, err := t.trackers.GetTorrentInfo(url)
	if err != nil || torrent.FileURL == "" {
		return nil, fmt.Errorf("cannot get link to .torrent file: %w", err)
	}

	_, content, err := t.trackers.DownloadTorrentFile(torrent)
	if err != nil {
		return nil, fmt.Errorf("cannot download .torrent file: %w", err)
	}

	addedTorrent, err := t.torrentClient.AddTorrent(content, folder, false)
	if err != nil {
		return nil, fmt.Errorf("cannot add torrent: %w", err)
	}
	transmissionTorrent := &models.TransmissionTorrent{
		Hash: addedTorrent.Hash,
	}
	err = t.torrentsStorage.SaveTransmission(transmissionTorrent)
	if err != nil {
		return nil, fmt.Errorf("cannot save transmissionTorrent: %w", err)
	}

	return &addedTorrent, nil
}

func (t *Torrents) GetActive(onlyRegistered bool) ([]*driven.Torrent, error) {
	var torrents []models.TransmissionTorrent
	err := t.torrentsStorage.GetAllTransmission(&torrents)
	if err != nil {
		return nil, fmt.Errorf("get registered: %w", err)
	}

	activeTorrents, err := t.torrentClient.GetTorrents()
	if err != nil {
		return nil, fmt.Errorf("get from client: %w", err)
	}

	var result []*driven.Torrent
	for _, torrent := range activeTorrents {
		blocked := false
		for _, path := range t.blockViewList {
			if strings.Contains(torrent.DownloadDir+torrent.Name, path) {
				blocked = true
			}
		}
		if blocked {
			continue
		}

		found := false
		if onlyRegistered {
			for _, registeredTorrent := range torrents {
				if registeredTorrent.Hash == torrent.Hash {
					found = true
					break
				}
			}
		}
		if found || !onlyRegistered {
			result = append(result, &torrent)
		}
	}

	return result, nil
}

func (t *Torrents) GetParts(id uint) ([]driven.TorrentFile, error) {
	files, err := t.torrentClient.GetTorrentFiles([]int{int(id)})
	if err != nil {
		return nil, fmt.Errorf("get files: %w", err)
	}
	return files, nil
}

func (t *Torrents) GetDownloadFolders() []string {
	folders := make([]string, len(t.downloadFolders))
	i := 0
	for folder := range t.downloadFolders {
		folders[i] = folder
		i++
	}

	return folders
}

func (t *Torrents) GetMonitored() ([]models.Torrent, error) {
	var torrents []models.Torrent
	err := t.torrentsStorage.Find(&torrents, "")
	if err != nil {
		return nil, err
	}
	return torrents, nil
}

func (t *Torrents) RenameParts(id uint, parts []*PartToRename) error {
	var err error
	for _, pair := range parts {
		err = t.torrentClient.Rename(int(id), pair.OldName, pair.NewName)
		if err != nil {
			return fmt.Errorf("cannot rename part ('%s' => '%s' for #%d): %w", pair.OldName, pair.NewName, id, err)
		}
	}
	return nil
}

func (t *Torrents) Search(text string) []*models.Torrent {
	torrents := t.trackers.SearchEverywhere(text)

	sort.Slice(torrents, func(i, j int) bool {
		return torrents[i].Seeders > torrents[j].Seeders
	})

	return torrents
}
