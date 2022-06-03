package trackers

import (
	"fmt"
	"net/url"
	"sync"

	"torrentsWatcher/internal/domain/models"
)

type TrackerAdapter interface {
	Matches(url string) bool
	GetInfo(url string) (*models.Torrent, error)
	Download(url string) (string, []byte, error)
	Search(text string) (torrents []*models.Torrent, err error)
}

type Adapter struct {
	Trackers []TrackerAdapter
}

func (f Adapter) GetTorrentInfo(torrentURL string) (*models.Torrent, error) {
	parser, err := f.getTracker(torrentURL)
	if err != nil {
		return nil, err
	}

	torrent, err := parser.GetInfo(torrentURL)
	return torrent, err
}

func (f Adapter) DownloadTorrentFile(torrent *models.Torrent) (string, []byte, error) {
	parser, err := f.getTracker(torrent.FileURL)
	if err != nil {
		return "", nil, err
	}
	return parser.Download(torrent.FileURL)
}

func (f Adapter) SearchEverywhere(text string) (torrents []*models.Torrent) {
	wg := sync.WaitGroup{}
	tChan := make(chan []*models.Torrent)
	for _, p := range f.Trackers {
		wg.Add(1)
		go func(p TrackerAdapter) {
			found, _ := p.Search(text)
			tChan <- found
		}(p)
	}

	q := make(chan interface{})
	go func() {
		for {
			select {
			case t := <-tChan:
				torrents = append(torrents, t...)
				wg.Done()
			case <-q:
				return
			}
		}
	}()

	wg.Wait()
	q <- nil

	return torrents
}

func (f Adapter) getTracker(torrentURL string) (TrackerAdapter, error) {
	parsedURL, err := url.Parse(torrentURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse url %s", torrentURL)
	}

	for _, parser := range f.Trackers {
		if parser.Matches(parsedURL.Host) {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("tracker not found for %s", torrentURL)
}
