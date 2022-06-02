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

	"torrentsWatcher/internal/models"
	"torrentsWatcher/internal/ports"
	"torrentsWatcher/internal/storage"
	"torrentsWatcher/internal/utils/network"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"golang.org/x/net/html/charset"
)

type Tracker struct {
	Logger          *zap.Logger
	Domain          string
	ForceHTTPS      bool
	Credentials     ports.Credentials
	TorrentsStorage storage.Torrents
	CookiesStorage  storage.Cookies
	Website         ports.WebsiteConnector
}

func (t *Tracker) GetInfo(url string) (*models.Torrent, error) {
	if t.ForceHTTPS {
		url = strings.Replace(url, "http://", "https://", 1)
	}

	torrent, err := t.loadAndParse(url)

	if err != nil && (err == ports.ErrUnauthorized || err.Error() == "record not found") {
		err = t.login()
		if err != nil {
			return nil, fmt.Errorf("login: %w", err)
		}

		torrent, err = t.loadAndParse(url)
	}

	if torrent != nil {
		torrent.PageURL = url
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
		t.Logger.Warn("cannot get media type", zap.Error(err), zap.Any("headers", headers), zap.Any("bodySize", data.Len()))
		fileName = fmt.Sprintf("download_%d.torrent", time.Now().Unix())
		err = nil
		return fileName, data.Bytes(), err
	}

	mediaType, _, _ := mime.ParseMediaType(headers.Get("Content-Type"))
	if mediaType != "application/x-bittorrent" {
		t.Logger.Warn("wrong media type for .torrent: %s", zap.String("mediaType", mediaType))
		return "", nil, errors.New("wrong media type for .torrent: " + mediaType)
	}

	fileName = params["filename"]

	return fileName, data.Bytes(), err
}

func (t *Tracker) Search(text string) (torrents []*models.Torrent, err error) {
	t.Logger.Info("Search", zap.String("text", text), zap.String("tracker", t.Domain))
	cookies, err := t.getCookies()
	if err != nil {
		return
	}

	r, err := t.Website.MakeSearchRequest(text)
	if err != nil {
		return
	}

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
		Transport: &http.Transport{
			// nolint: gosec
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(r)
	if err != nil {
		t.Logger.Info("Search failed", zap.Error(err), zap.String("text", text), zap.String("tracker", t.Domain))
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 302 && strings.Contains(res.Header.Get("Location"), "login") {
		err = t.login()
		if err != nil {
			return nil, fmt.Errorf("login: %w", err)
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
			t.Logger.Info("Search failed after login", zap.Error(err), zap.String("text", text), zap.String("tracker", t.Domain))
			return
		}
		defer res.Body.Close()
	}

	body, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	if err != nil {
		t.Logger.Warn(fmt.Sprintf(
			"failed to search in %s code=%d len=%d\n%+v",
			t.Domain,
			res.StatusCode,
			res.ContentLength,
			res.Header,
		))
		return
	}
	if body == nil {
		t.Logger.Warn(fmt.Sprintf("failed to search in %s %d\n%+v", t.Domain, res.StatusCode, res.Header))
		return
	}
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	torrents, _ = t.Website.ParseSearch(document)

	return torrents, nil
}

func (t *Tracker) loadAndParse(url string) (*models.Torrent, error) {
	t.Logger.Debug("loadAndParse", zap.String("url", url))
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

	return t.Website.Parse(document)
}

func (t *Tracker) getCookies() ([]*http.Cookie, error) {
	var cookies []*http.Cookie
	if t.Credentials != (ports.Credentials{}) {
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
	t.Logger.Info("login: " + t.Domain)
	cookies, err := t.Website.Login(t.Credentials)
	if err != nil {
		return err
	}

	if err := t.CookiesStorage.DeleteByDomain(t.Domain); err != nil {
		t.Logger.Error("removing old cookies", zap.Error(err))
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
