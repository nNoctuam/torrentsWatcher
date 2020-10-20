package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"

	"torrentsWatcher/config"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/parser"
	"torrentsWatcher/internal/api/parser/impl"
	"torrentsWatcher/internal/api/watch"
	"torrentsWatcher/internal/handlers"
	"torrentsWatcher/internal/storage"
	impl2 "torrentsWatcher/internal/storage/impl"
)

// TODO:
// 	notifications:
// 		browser
// 		messenger
// 		email
// 	docker build
// 	log instead of fmt.Print
// 	supervisor config
// 	unit tests
// 	DI

func main() {
	cfg := config.Load()
	notificator := getNotificator(cfg)

	db, err := gorm.Open("sqlite3", "./torrents.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.Torrent{}, &models.AuthCookie{})

	torrentsStorage := impl2.NewTorrentsSqliteStorage(db)
	cookiesStorage := impl2.NewCookiesSqliteStorage(db)

	parsers := []*parser.Tracker{
		impl.NewNnmClub(cfg.Credentials[impl.NnmClubDomain], torrentsStorage, cookiesStorage),
		impl.NewRutracker(cfg.Credentials[impl.RutrackerDomain], torrentsStorage, cookiesStorage),
	}

	go watch.Watch(
		cfg.Interval,
		parsers,
		notificator,
		torrentsStorage,
		cookiesStorage,
	)
	serve(cfg.Host, cfg.Port, parsers, torrentsStorage)
}

func serve(
	host string,
	port string,
	parsers []*parser.Tracker,
	torrentsStorage storage.Torrents,
) {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/torrents", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTorrents(w, r, torrentsStorage)
	})
	router.MethodFunc("POST", "/torrent", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTorrent(w, r, parsers, torrentsStorage)
	})
	router.MethodFunc("GET", `/torrent/{id:\d+}/download`, func(w http.ResponseWriter, r *http.Request) {
		handlers.DownloadTorrent(w, r, torrentsStorage)
	})
	router.Handle("/*", http.FileServer(http.Dir("./frontend/dist")))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	_ = server.ListenAndServe()
}

func getNotificator(cfg *config.AppConfig) notification.Notificator {
	switch runtime.GOOS {
	case "windows":
		return &notification.Windows{Config: notification.Config(cfg.Notifications)}
	case "linux":
		fallthrough
	default:
		return &notification.Linux{Config: notification.Config(cfg.Notifications)}
	}
}
