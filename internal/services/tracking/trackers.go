package tracking

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
	"torrentsWatcher/internal/core/models"
)

type Trackers []*Tracker

func (f Trackers) GetTorrentInfo(torrentURL string) (*models.Torrent, error) {
	parser, err := f.getTracker(torrentURL)
	if err != nil {
		return nil, err
	}

	torrent, err := parser.GetInfo(torrentURL)
	return torrent, err
}

func (f Trackers) DownloadTorrentFile(torrent *models.Torrent) (string, []byte, error) {
	parser, err := f.getTracker(torrent.FileURL)
	if err != nil {
		return "", nil, err
	}
	return parser.Download(torrent.FileURL)
}

func (f Trackers) SearchEverywhere(text string) (torrents []*models.Torrent) {
	wg := sync.WaitGroup{}
	tChan := make(chan []*models.Torrent)
	for _, p := range f {
		wg.Add(1)
		go func(p *Tracker) {
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

func (f Trackers) getTracker(torrentURL string) (*Tracker, error) {
	parsedURL, err := url.Parse(torrentURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse url %s", torrentURL)
	}

	for _, parser := range f {
		if parser.Domain == parsedURL.Host {
			return parser, nil
		}
	}
	for _, parser := range f {
		if strings.Contains(parsedURL.Host, parser.Domain) {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("tracker not found for %s", torrentURL)
}
