package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"runtime"
	"time"
	"torrentsWatcher/config"
	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/watch"
	"torrentsWatcher/internal/handlers"
)

// tracker parsing
// adding torrent for watching - db
// storing auth cookies
// notifications
// cron to check torrents
// authorization

func main() {
	config.Load()
	//c, _ := json.Marshal(config.App)
	//fmt.Print(string(c))
	//return

	db.InitDB()
	defer db.CloseDB()

	fmt.Print("it works.\n")

	migrate()
	initNotifications()

	go watch.Watch(time.Duration(config.App.IntervalHours) * time.Hour)
	serve(config.App.Host, config.App.Port)
}

func serve(host string, port string) {
	router := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(corsMiddleware.Handler)

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
