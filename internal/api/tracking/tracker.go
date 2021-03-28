package tracking

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/utils/network"
	"torrentsWatcher/internal/storage"
)

var UnauthorizedError = errors.New("unauthorized")

type Credentials struct {
	Login    string
	Password string
}

type Tracker struct {
	Domain          string
	ForceHttps      bool
	Credentials     Credentials
	TorrentsStorage storage.Torrents
	CookiesStorage  storage.Cookies
	Impl            TrackerImpl
}

func (t *Tracker) GetInfo(url string) (*models.Torrent, error) {
	if t.ForceHttps {
		url = strings.Replace(url, "http://", "https://", 1)
	}

	torrent, err := t.loadAndParse(url)

	if err != nil && (err == UnauthorizedError || err.Error() == "record not found") {
		err = t.login()
		if err != nil {
			return nil, err
		}

		torrent, err = t.loadAndParse(url)
	}

	if torrent != nil {
		torrent.PageUrl = url
	}

	return torrent, err
}

func (t *Tracker) Download(url string) (string, []byte, error) {
	log.Println("downloading ", url)
	cookies, err := t.getCookies()
	if err != nil {
		return "", nil, err
	}

	headers, body, err := network.LoadBytes(url, cookies)
	if err != nil {
		return "", nil, err
	}
	var data bytes.Buffer
	_, err = io.Copy(&data, body)
	if err != nil {
		return "", nil, err
	}

	var fileName string
	_, params, err := mime.ParseMediaType(headers.Get("Content-Disposition"))
	if err != nil {
		log.Printf("cannot get media type: %v", err)
		log.Printf("%+v", headers)
		log.Printf("body size: %d", data.Len())
		fileName = fmt.Sprintf("download_%d.torrent", time.Now().Unix())
		err = nil
		return fileName, data.Bytes(), err
	}

	mediaType, _, _ := mime.ParseMediaType(headers.Get("Content-Type"))
	if mediaType != "application/x-bittorrent" {
		log.Printf("wrong media type for .torrent: %s", mediaType)
		return "", nil, errors.New("wrong media type for .torrent: " + mediaType)
	}

	fileName = params["filename"]

	return fileName, data.Bytes(), err
}

func (t *Tracker) Search(text string) (torrents []*models.Torrent, err error) {
	log.Println("Search " + text)
	cookies, err := t.getCookies()
	if err != nil {
		return
	}

	r, err := t.Impl.MakeSearchRequest(text)
	if err != nil {
		return
	}

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode == 302 && strings.Contains(res.Header.Get("Location"), "login") {
		err = t.login()
		if err != nil {
			return nil, err
		}
		cookies, err = t.getCookies()
		if err != nil {
			return
		}
		for _, cookie := range cookies {
			r.AddCookie(cookie)
		}

		res, err = client.Do(r)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	body, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	if err != nil {
		log.Printf("failed to search in %s code=%d len=%d\n%+v", t.Domain, res.StatusCode, res.ContentLength, res.Header)
		return
	}
	if body == nil {
		log.Printf("failed to search in %s %d\n%+v", t.Domain, res.StatusCode, res.Header)
		return
	}
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	torrents, _ = t.Impl.ParseSearch(document)

	return torrents, nil
}

func (t *Tracker) loadAndParse(url string) (*models.Torrent, error) {
	log.Println("loadAndParse " + url)
	cookies, err := t.getCookies()
	if err != nil {
		return nil, err
	}

	_, body, err := network.LoadHTML(url, cookies)
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	return t.Impl.Parse(document)
}

func (t *Tracker) getCookies() ([]*http.Cookie, error) {
	var cookies []*http.Cookie
	if t.Credentials != (Credentials{}) {
		var savedCookies []models.AuthCookie
		if err := t.CookiesStorage.Find(&savedCookies, &models.AuthCookie{Domain: t.Domain}); err != nil {
			return nil, err
		}
		for _, savedCookie := range savedCookies {
			cookies = append(cookies, &http.Cookie{
				Name:  savedCookie.Name,
				Value: savedCookie.Value,
			})
		}
	}
	return cookies, nil
}

func (t *Tracker) login() error {
	var cookies []*http.Cookie
	cookies, err := t.Impl.Login(t.Credentials)
	if err != nil {
		return err
	}

	if err := t.CookiesStorage.DeleteByDomain(t.Domain); err != nil {
		fmt.Println("error removing old cookies")
	}

	for _, cookie := range cookies {
		err = t.CookiesStorage.Save(&models.AuthCookie{
			Domain: t.Domain,
			Name:   cookie.Name,
			Value:  cookie.Value,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
