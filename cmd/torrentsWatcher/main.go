package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/go-chi/chi"

	"torrentsWatcher/config"
	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/parser"
	"torrentsWatcher/internal/api/parser/impl"
	"torrentsWatcher/internal/api/watch"
	"torrentsWatcher/internal/handlers"
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
	parsers := []*parser.Tracker{
		impl.NewNnmClub(cfg.Credentials[impl.NnmClubDomain]),
		impl.NewRutracker(cfg.Credentials[impl.RutrackerDomain]),
	}

	db.InitDB()
	defer db.CloseDB()
	migrate()

	fmt.Print("it works.\n")

	go watch.Watch(time.Duration(cfg.IntervalHours)*time.Hour, parsers, notificator)
	serve(cfg.Host, cfg.Port, parsers)
}

func serve(host string, port string, parsers []*parser.Tracker) {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/torrents", handlers.GetTorrents)
	router.MethodFunc("POST", "/torrent", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTorrent(w, r, parsers)
	})
	router.MethodFunc("GET", `/torrent/{id:\d+}/download`, handlers.DownloadTorrent)
	router.Handle("/*", http.FileServer(http.Dir("./frontend/dist")))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	_ = server.ListenAndServe()
}

func migrate() {
	db.DB.AutoMigrate(&models.Torrent{}, &models.AuthCookie{})
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
