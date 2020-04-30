package tracker

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/config"
	"torrentsWatcher/internal/api/models"
)

type Rutracker struct{}

func NewRutracker() *Tracker {
	return &Tracker{
		Domain:   "rutracker.org",
		iTracker: &Rutracker{},
	}
}

func (t *Rutracker) doesRequireLogin() bool {
	return true
}

func (t *Rutracker) parse(document *goquery.Document) (*models.Torrent, error) {
	var info models.Torrent
	var err error

	info.Title = document.Find(".maintitle").First().Text()

	if info.Title != "" && document.Find("#logged-in-username").Size() == 0 {
		fmt.Println("Unauthorized")
		return nil, errors.New(UnauthorizedError)
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

	uploadedAt := document.Find(".attach.bordered.med .row1 td li").First().Text()

	location, _ := time.LoadLocation("Local")
	info.UploadedAt, err = time.ParseInLocation("02-Jan-06 15:04", r.Replace(uploadedAt), location)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse date '%s'", uploadedAt)
	}

	return &info, err
}

func (t *Rutracker) login() (*http.Cookie, error) {
	data := url.Values{}
	data.Set("login_username", config.App.Credentials["rutracker.org"].Login)
	data.Set("login_password", config.App.Credentials["rutracker.org"].Password)
	data.Set("login", "%E2%F5%EE%E4")

	fmt.Println("login...")

	request, err := http.NewRequest("POST", "https://rutracker.org/forum/login.php", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == "bb_session" {
			return cookie, nil
		}
	}
	return nil, errors.New("cookie not found in response")
}
