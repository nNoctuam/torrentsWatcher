package tracker

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/models"
)

type NnmClub struct {
	Tracker
}

func NewNnmClub() *Tracker {
	return &Tracker{
		Domain:   "nnmclub.to",
		iTracker: &NnmClub{},
	}
}

func (t *NnmClub) doesRequireLogin() bool {
	return false
}

func (t *NnmClub) parse(document *goquery.Document) (*models.Torrent, error) {
	var info models.Torrent
	var err error

	info.Title = document.Find(".maintitle").First().Text()
	info.UploadedAt, err = parseNnmClubUploadedAt(document)

	return &info, err
}

func (t *NnmClub) login() (*http.Cookie, error) {
	return nil, nil
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
