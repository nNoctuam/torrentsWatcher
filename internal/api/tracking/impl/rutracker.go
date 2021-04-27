package impl

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"torrentsWatcher/internal/storage"

	"go.uber.org/zap"

	"golang.org/x/text/encoding/charmap"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tracking"
)

type Rutracker struct {
	//nolint:structcheck,unused
	logger *zap.Logger
}

const RutrackerDomain = "rutracker.org"

func NewRutracker(logger *zap.Logger, credentials tracking.Credentials, torrents storage.Torrents, cookies storage.Cookies) *tracking.Tracker {
	return &tracking.Tracker{
		Logger:          logger,
		Domain:          RutrackerDomain,
		ForceHttps:      true,
		Credentials:     credentials,
		TorrentsStorage: torrents,
		CookiesStorage:  cookies,
		Impl:            &Rutracker{},
	}
}

func (t *Rutracker) Parse(document *goquery.Document) (*models.Torrent, error) {
	var info models.Torrent
	var err error

	info.Title = document.Find(".maintitle").First().Text()
	info.Title = strings.Trim(info.Title, " \t\n")
	info.FileUrl, _ = document.Find(".dl-stub.dl-link.dl-topic").First().Attr("href")

	if len(info.FileUrl) > 6 && info.FileUrl[:6] == "dl.php" {
		info.FileUrl = "https://rutracker.org/forum/" + info.FileUrl
	}

	if info.Title != "" && document.Find("#logged-in-username").Size() == 0 {
		fmt.Println("Unauthorized")
		return nil, tracking.UnauthorizedError
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

func (t *Rutracker) Login(credentials tracking.Credentials) ([]*http.Cookie, error) {
	data := url.Values{}
	data.Set("login_username", credentials.Login)
	data.Set("login_password", credentials.Password)
	data.Set("login", "%E2%F5%EE%E4")

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

	return response.Cookies(), nil
}

func (t *Rutracker) ParseSearch(document *goquery.Document) (torrents []*models.Torrent, err error) {
	headers := document.Find("#tor-tbl thead th").Nodes
	columns := map[string]int{}

	for i, th := range headers {
		if th.FirstChild == nil {
			continue
		}
		columns[th.FirstChild.Data] = i
	}

	document.Find("#tor-tbl tbody tr").Each(func(i int, row *goquery.Selection) {
		torrent := &models.Torrent{}
		tds := row.Find("td")
		forumTD := tds.Get(columns["Форум"])
		titleTD := tds.Get(columns["Тема"])
		authorTD := tds.Get(columns["Автор"])
		sizeTD := tds.Get(columns["Размер"])
		seedersTD := tds.Get(columns["S"])
		addedTD := tds.Get(columns["Добавлен"])

		if strings.Contains(titleTD.FirstChild.Data, "Не найдено") {
			return
		}

		for _, attr := range titleTD.FirstChild.NextSibling.FirstChild.NextSibling.Attr {
			if attr.Key == "href" {
				torrent.PageUrl = "https://" + RutrackerDomain + "/forum/" + attr.Val
				break
			}
		}

		torrent.Forum = forumTD.FirstChild.NextSibling.FirstChild.FirstChild.Data
		torrent.Title = titleTD.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.Data
		torrent.Title = titleTD.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.Data
		torrent.Author = authorTD.FirstChild.NextSibling.FirstChild.FirstChild.Data
		torrent.Seeders, _ = strconv.ParseUint(seedersTD.FirstChild.NextSibling.FirstChild.Data, 10, 32)

		var sizeData []string
		reg, _ := regexp.Compile(`^\s*([\d.]+).+([KMG])B`)
		if sizeTD.FirstChild.NextSibling != nil {
			sizeData = reg.FindStringSubmatch(sizeTD.FirstChild.NextSibling.FirstChild.Data)
		} else {
			sizeData = reg.FindStringSubmatch(sizeTD.FirstChild.Data)
		}
		size, _ := strconv.ParseFloat(sizeData[1], 10)
		switch sizeData[2] {
		case "K":
			size *= 1000
		case "M":
			size *= 1000 * 1000
		case "G":
			size *= 1000 * 1000 * 1000
		}
		torrent.Size = uint64(size)

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
		torrent.UpdatedAt, _ = time.ParseInLocation("2-Jan-06", r.Replace(addedTD.FirstChild.NextSibling.FirstChild.Data), location)

		torrents = append(torrents, torrent)
	})

	return torrents, err
}

func (t *Rutracker) MakeSearchRequest(text string) (r *http.Request, err error) {

	encoder := charmap.Windows1251.NewEncoder()
	text, _ = encoder.String(text)

	params := url.Values{}
	params.Set("nm", text)
	params.Set("o", "10") // sort by seeders
	r, err = http.NewRequest("POST", "https://"+RutrackerDomain+"/forum/tracker.php", strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Content-Length", strconv.Itoa(len(params.Encode())))

	return
}
