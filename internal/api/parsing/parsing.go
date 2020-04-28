package parsing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parsing/implementations"
)

var parsers = map[string]func(document *goquery.Document) (*models.Torrent, error){
	"nnmclub.to": implementations.ParseNnmClub,
}

func GetTorrentInfo(torrentUrl string) (*models.Torrent, error) {
	body, err := loadHTML(torrentUrl)
	if err != nil {
		return &models.Torrent{}, err
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	parser := getParser(torrentUrl)
	if parser == nil {
		return &models.Torrent{}, errors.New("Parser not found for url " + torrentUrl)
	}

	torrent, err := parser(doc)
	jsonView, _ := json.Marshal(torrent)
	fmt.Printf("parsed: %s (err = %v)\n", jsonView, err)

	return torrent, err
}

func getParser(torrentUrl string) func(document *goquery.Document) (*models.Torrent, error) {
	parsedUrl, err := url.Parse(torrentUrl)
	if err != nil {
		fmt.Printf("Couldn't parse url %s", torrentUrl)
		return nil
	}

	parser, exists := parsers[parsedUrl.Host]
	if !exists {
		return nil
	}
	return parser
}

func loadHTML(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utf8, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Encoding error:", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		fmt.Println("IO error:", err)
		return nil, err
	}
	return bytes.NewReader(body), nil
}
