package impl

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parser"
	"torrentsWatcher/internal/storage"
)

type NnmClub struct{}

const NnmClubDomain = "nnmclub.to"

func NewNnmClub(credentials parser.Credentials, torrentsStorage storage.Torrents, cookiesStorage storage.Cookies) *parser.Tracker {
	return &parser.Tracker{
		Domain:          NnmClubDomain,
		ForceHttps:      true,
		Credentials:     credentials,
		TorrentsStorage: torrentsStorage,
		CookiesStorage:  cookiesStorage,
		Impl:            &NnmClub{},
	}
}

func (t *NnmClub) Parse(document *goquery.Document) (*models.Torrent, error) {
	var info models.Torrent
	var err error

	if document.Find("table.btTbl tr.row1 td.gensmall span b a").First().Text() != "Скачать" {
		return &info, errors.New(parser.UnauthorizedError)
	}

	info.Title = document.Find(".maintitle").First().Text()
	info.UploadedAt, err = parseNnmClubUploadedAt(document)
	info.FileUrl, _ = document.Find("table.btTbl tr.row1 td.gensmall span b a").First().Attr("href")
	if info.FileUrl[:8] == "download" {
		info.FileUrl = "https://nnmclub.to/forum/" + info.FileUrl
	}

	return &info, err
}

func (t *NnmClub) Login(credentials parser.Credentials) ([]*http.Cookie, error) {
	fmt.Println("login", t)
	code, err := getLoginCode()

	params := url.Values{}
	params.Set("username", credentials.Login)
	params.Set("password", credentials.Password)
	params.Set("autologin", "on")
	params.Set("redirect", "")
	params.Set("code", code)
	params.Set("login", "%C2%F5%EE%E4")

	r, err := http.NewRequest("POST", "https://nnmclub.to/forum/login.php", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Content-Length", strconv.Itoa(len(params.Encode())))

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res.Cookies(), nil
}

func getLoginCode() (string, error) {
	r, err := http.NewRequest("GET", "https://nnmclub.to/forum/login.php", nil)
	if err != nil {
		return "", err
	}

	r.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linu…) Gecko/20100101 Firefox/75.0")
	r.Header.Set("Origin", "http://nnmclub.to")

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login page response is %d %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	code, _ := doc.Find(`input[name="code"]`).First().Attr("value")
	if code == "" {
		return "", errors.New("hidden input with code not found")
	}
	err = res.Body.Close()
	if err != nil {
		return "", err
	}

	return code, nil
}

func parseNnmClubUploadedAt(document *goquery.Document) (time.Time, error) {
	previous := " Зарегистрирован: "
	var updatedAtNodeId int
	var uploadedAt string

	document.Find("table.btTbl tr.row1 td.genmed").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if i == updatedAtNodeId {
			uploadedAt = s.Text()
		}
		if text == previous {
			updatedAtNodeId = i + 1
		}
	})

	if uploadedAt == "" {
		return time.Time{}, errors.New("couldn't detect updated at")
	}

	r := strings.NewReplacer(
		"Янв", "Jan",
		"Фев", "Feb",
		"Мар", "Mar",
		"Апр", "Apr",
		"Май", "May",
		"Июн", "Jun",
		"Июл", "Jul",
		"Авг", "Aug",
		"Сен", "Sep",
		"Окт", "Oct",
		"Ноя", "Nov",
		"Дек", "Dec",
	)
	location, _ := time.LoadLocation("Local")
	return time.ParseInLocation("02 Jan 2006 15:04:05", strings.Trim(r.Replace(uploadedAt), " "), location)
}
