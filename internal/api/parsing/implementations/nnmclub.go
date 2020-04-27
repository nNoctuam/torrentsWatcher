package implementations

import (
	"errors"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/models"
)

func ParseNnmClub(document *goquery.Document) (models.Torrent, error) {
	var info models.Torrent
	var err error

	info.Title = document.Find(".maintitle").First().Text()
	info.UpdatedAt, err = parseNnmClubUpdatedAt(document)

	return info, err
}

func parseNnmClubUpdatedAt(document *goquery.Document) (time.Time, error) {
	previous := " Зарегистрирован: "
	var updatedAtNodeId int
	var updatedAt string

	document.Find("table.btTbl tr.row1 td.genmed").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if i == updatedAtNodeId {
			updatedAt = s.Text()
		}
		if text == previous {
			updatedAtNodeId = i + 1
		}
	})

	if updatedAt == "" {
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

	return time.Parse("02 Jan 2006 15:04:05", strings.Trim(r.Replace(updatedAt), " "))
}
