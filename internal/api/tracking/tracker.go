package tracking

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

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
		var cookies []*http.Cookie
		cookies, err = t.Impl.Login(t.Credentials)
		if err != nil {
			return nil, err
		}

		if err := t.CookiesStorage.Delete(&models.AuthCookie{Domain: t.Domain}); err != nil {
			fmt.Println("error removing old cookies")
		}

		for _, cookie := range cookies {
			err = t.CookiesStorage.Save(&models.AuthCookie{
				Domain: t.Domain,
				Name:   cookie.Name,
				Value:  cookie.Value,
			})
			if err != nil {
				return nil, err
			}
		}

		torrent, err = t.loadAndParse(url)
	}

	if torrent != nil {
		torrent.PageUrl = url
	}

	return torrent, err
}

func (t *Tracker) Download(url string) ([]byte, error) {
	cookies, err := t.getCookies()
	if err != nil {
		return nil, err
	}

	body, err := network.LoadBytes(url, cookies)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, body)

	return data.Bytes(), err
}

func (t *Tracker) Search(text string) (torrents []*models.Torrent, err error) {
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
	body, _ := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	torrents, _ = t.Impl.ParseSearch(document)

	return torrents, nil
}

func (t *Tracker) loadAndParse(url string) (*models.Torrent, error) {
	cookies, err := t.getCookies()
	if err != nil {
		return nil, err
	}

	body, err := network.LoadHTML(url, cookies)
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