package tracking

import (
	"net/http"

	"torrentsWatcher/internal/models"

	"github.com/PuerkitoBio/goquery"
)

type Credentials struct {
	Login    string
	Password string
}

type WebsiteConnector interface {
	Login(credentials Credentials) ([]*http.Cookie, error)
	Parse(document *goquery.Document) (*models.Torrent, error)
	MakeSearchRequest(text string) (r *http.Request, err error)
	ParseSearch(document *goquery.Document) (torrents []*models.Torrent, err error)
}
