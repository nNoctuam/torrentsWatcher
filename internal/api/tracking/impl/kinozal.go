package impl

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"

	"github.com/PuerkitoBio/goquery"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tracking"
	"torrentsWatcher/internal/storage"
)

type Kinozal struct{}

const KinozalDomain = "kinozal.tv"

func NewKinozal(credentials tracking.Credentials, torrentsStorage storage.Torrents, cookiesStorage storage.Cookies) *tracking.Tracker {
	return &tracking.Tracker{
		Domain:          KinozalDomain,
		ForceHttps:      false,
		Credentials:     credentials,
		TorrentsStorage: torrentsStorage,
		CookiesStorage:  cookiesStorage,
		Impl:            &Kinozal{},
	}
}

func (t *Kinozal) Parse(document *goquery.Document) (*models.Torrent, error) {
	var info models.Torrent
	var err error

	authLink := document.Find(".menu form ul.men a")
	if authLink.Size() > 0 {
		a := authLink.Get(0)
		log.Printf("%+v", a)
	}
	if authLink.Size() > 0 && authLink.Get(0).FirstChild != nil && authLink.Get(0).FirstChild.Data == "Гость! ( Зарегистрируйтесь )" {
		log.Println(KinozalDomain+" unauthorized:", authLink.Get(0).FirstChild.Data)
		return &info, tracking.UnauthorizedError
	}

	info.Title = document.Find("h1 a.r0").First().Text()
	//info.UploadedAt, err = parseKinozalUploadedAt(document) todo
	info.FileUrl, _ = document.Find(".mn1_content td.nw a").First().Attr("href")
	info.FileUrl = "http:" + info.FileUrl
	info.FileUrl = strings.Replace(info.FileUrl, "dl."+KinozalDomain, KinozalDomain, 1)

	return &info, err
}

func (t *Kinozal) ParseSearch(document *goquery.Document) (torrents []*models.Torrent, err error) {
	headers := document.Find("table.t_peer tr.mn td").Nodes
	columns := map[string]int{}

	for i, th := range headers {
		if th.FirstChild == nil {
			continue
		}
		columns[th.FirstChild.Data] = i
	}

	document.Find("table.t_peer tr.bg").Each(func(i int, row *goquery.Selection) {
		torrent := &models.Torrent{}
		tds := row.Find("td")
		forumTD := tds.Get(0)
		titleTD := tds.Get(1)
		authorTD := tds.Get(columns["Раздает"])
		sizeTD := tds.Get(columns["Размер"])
		seedersTD := tds.Get(columns["Сидов"])
		//addedTD := tds.Get(columns["Залит"])

		for _, attr := range titleTD.FirstChild.Attr {
			if attr.Key == "href" {
				torrent.PageUrl = "http://" + KinozalDomain + "/" + attr.Val
				break
			}
		}

		torrent.Forum = forumTD.FirstChild.Data
		torrent.Title = titleTD.FirstChild.FirstChild.Data
		torrent.Seeders, _ = strconv.ParseUint(seedersTD.FirstChild.Data, 10, 64)

		sizes := map[string]float64{"КБ": 1, "МБ": 2, "ГБ": 3, "ТВ": 4}

		size := strings.Split(sizeTD.FirstChild.Data, " ")
		sizeBytes, _ := strconv.ParseFloat(size[0], 64)

		torrent.Size = uint64(math.Round(sizeBytes * math.Pow(1024, sizes[size[1]])))
		torrent.Author = authorTD.FirstChild.FirstChild.Data
		// todo
		//date := addedTD.FirstChild.NextSibling.Data + " " + addedTD.FirstChild.NextSibling.NextSibling.NextSibling.Data
		//torrent.UpdatedAt, _ = time.Parse("02-01-2006 15:04", strings.Trim(date, " "))

		torrents = append(torrents, torrent)
	})

	return torrents, err
}

func (t *Kinozal) MakeSearchRequest(text string) (r *http.Request, err error) {

	encoder := charmap.Windows1251.NewEncoder()
	//nolint:ineffassign
	text, err = encoder.String(text)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("s", text)
	params.Set("g", "0")
	params.Set("c", "0")
	params.Set("v", "0")
	params.Set("d", "0")
	params.Set("t", "1") // sort by seeders
	params.Set("f", "0") // sort desc
	r, err = http.NewRequest("GET", "http://"+KinozalDomain+"/browse.php?"+params.Encode(), nil)
	if err != nil {
		return
	}
	r.Header.Set("Content-Type", "text/html; charset=windows-1251")
	r.Header.Set("Referer", "http://"+KinozalDomain+"/browse.php")

	return
}

func (t *Kinozal) Login(credentials tracking.Credentials) ([]*http.Cookie, error) {
	fmt.Println(KinozalDomain + " login")

	params := url.Values{}
	params.Set("username", credentials.Login)
	params.Set("password", credentials.Password)
	params.Set("returnto", "")

	r, err := http.NewRequest("POST", "http://"+KinozalDomain+"/takelogin.php", strings.NewReader(params.Encode()))
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
