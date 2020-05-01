package parsing

import (
	"encoding/json"
	"fmt"
	"net/url"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parsing/tracker"
)

var parsers = []*tracker.Tracker{
	tracker.NewNnmClub(),
	tracker.NewRutracker(),
}

func GetTorrentInfo(torrentUrl string) (*models.Torrent, error) {
	parser, err := getParser(torrentUrl)
	if err != nil {
		return nil, err
	}

	torrent, err := (*parser).GetInfo(torrentUrl)
	jsonView, _ := json.Marshal(torrent)
	fmt.Printf("parsed: %s (err = %v)\n", jsonView, err)

	return torrent, err
}

func DownloadTorrentFile(torrent *models.Torrent) ([]byte, error) {
	parser, err := getParser(torrent.FileUrl)
	if err != nil {
		return nil, err
	}
	return (*parser).Download(torrent.FileUrl)
}

func getParser(torrentUrl string) (*tracker.Tracker, error) {
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
