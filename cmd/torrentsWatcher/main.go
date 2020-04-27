package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
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
	db.InitDB()
	defer db.CloseDB()

	fmt.Print("it works.\n")

	migrate()

	go watch.Watch(1 * time.Hour)
	serve()
}

func serve() {
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
	router.Handle("/*", http.FileServer(http.Dir("./frontend/dist")))

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	server.ListenAndServe()
}

func migrate() {
	db.DB.AutoMigrate(&models.Torrent{})
}
