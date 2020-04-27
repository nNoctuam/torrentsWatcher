package parsing

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parsing/implementations"
)

func GetTorrentInfo(url string) (models.Torrent, error) {
	body, err := loadHTML(url)
	if err != nil {
		return models.Torrent{}, err
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	parser := getParser(url)

	info, err := parser(doc)
	info.PageUrl = url
	fmt.Printf("parsed: %v (err = %v)\n", info, err)

	return info, err
}

func getParser(url string) func(document *goquery.Document) (models.Torrent, error) {
	return implementations.ParseNnmClub
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
