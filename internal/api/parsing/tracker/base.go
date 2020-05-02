package tracker

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tools"
)

type iTracker interface {
	login() ([]*http.Cookie, error)
	doesRequireLogin() bool
	parse(document *goquery.Document) (*models.Torrent, error)
}

const UnauthorizedError = "unauthorized"

type Tracker struct {
	Domain     string
	ForceHttps bool
	iTracker   iTracker
}

func (t *Tracker) GetInfo(url string) (*models.Torrent, error) {
	if t.ForceHttps {
		url = strings.Replace(url, "http://", "https://", 1)
	}

	torrent, err := t.LoadAndParse(url)

	if err != nil && (err.Error() == UnauthorizedError || err.Error() == "record not found") {
		var cookies []*http.Cookie
		cookies, err = t.iTracker.login()
		if err != nil {
			return nil, err
		}

		if err := db.DB.Where(&models.AuthCookie{Domain: t.Domain}).Delete(models.AuthCookie{}).Error; err != nil {
			fmt.Println("error removing old cookies")
		}

		for _, cookie := range cookies {
			err = db.DB.Save(&models.AuthCookie{
				Domain: t.Domain,
				Name:   cookie.Name,
				Value:  cookie.Value,
			}).Error
			if err != nil {
				return nil, err
			}
		}

		torrent, err = t.LoadAndParse(url)
	}

	if torrent != nil {
		torrent.PageUrl = url
	}

	return torrent, err
}

func (t *Tracker) LoadAndParse(url string) (*models.Torrent, error) {
	cookies, err := t.getCookies()
	if err != nil {
		return nil, err
	}

	body, err := tools.LoadHTML(url, cookies)
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	return t.iTracker.parse(document)
}

func (t *Tracker) getCookies() ([]*http.Cookie, error) {
	var cookies []*http.Cookie
	if t.iTracker.doesRequireLogin() {
		var savedCookies []models.AuthCookie
		if err := db.DB.Where(&models.AuthCookie{Domain: t.Domain}).Find(&savedCookies).Error; err != nil {
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

func (t *Tracker) Download(url string) ([]byte, error) {
	cookies, err := t.getCookies()
	if err != nil {
		return nil, err
	}

	body, err := tools.LoadBytes(url, cookies)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, body)

	return data.Bytes(), err
}
