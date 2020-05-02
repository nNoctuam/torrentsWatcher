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
	"torrentsWatcher/internal/api/watch"
	"torrentsWatcher/internal/handlers"
)

func main() {
	config.Load()

	db.InitDB()
	migrate()
	defer db.CloseDB()

	initNotifications()

	fmt.Print("it works.\n")

	go watch.Watch(time.Duration(config.App.IntervalHours) * time.Hour)
	serve(config.App.Host, config.App.Port)
}

func serve(host string, port string) {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/torrents", handlers.GetTorrents)
	router.MethodFunc("POST", "/torrent", handlers.AddTorrent)
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

func initNotifications() {
	switch runtime.GOOS {
	case "windows":
		notification.Notificator = &notification.Windows{}
	case "linux":
		fallthrough
	default:
		notification.Notificator = &notification.Linux{}
	}
}
