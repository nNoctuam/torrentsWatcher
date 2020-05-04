package parser

import (
	"encoding/json"
	"fmt"
	"net/url"

	"torrentsWatcher/internal/api/models"
)

func GetTorrentInfo(torrentUrl string, parsers []*Tracker) (*models.Torrent, error) {
	parser, err := getParser(torrentUrl, parsers)
	if err != nil {
		return nil, err
	}

	torrent, err := (*parser).GetInfo(torrentUrl)
	jsonView, _ := json.Marshal(torrent)
	fmt.Printf("parsed: %s (err = %v)\n", jsonView, err)

	return torrent, err
}

func DownloadTorrentFile(torrent *models.Torrent, parsers []*Tracker) ([]byte, error) {
	parser, err := getParser(torrent.FileUrl, parsers)
	if err != nil {
		return nil, err
	}
	return (*parser).Download(torrent.FileUrl)
}

func getParser(torrentUrl string, parsers []*Tracker) (*Tracker, error) {
	parsedUrl, err := url.Parse(torrentUrl)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse url %s", torrentUrl)
	}

	for _, parser := range parsers {
		if parser.Domain == parsedUrl.Host {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("parser not found for %s", torrentUrl)
}
