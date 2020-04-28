package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"torrentsWatcher/config"
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
	cfg := getConfig()

	db.InitDB()
	defer db.CloseDB()

	fmt.Print("it works.\n")

	migrate()

	go watch.Watch(time.Duration(cfg.IntervalHours) * time.Hour)
	serve(cfg.Host, cfg.Port)
}

func getConfig() *config.AppConfig {
	cfg := &config.AppConfig{}

	dat, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	err = yaml.Unmarshal(dat, cfg)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return cfg
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
	router.Handle("/*", http.FileServer(http.Dir("./frontend/dist")))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	_ = server.ListenAndServe()
}

func migrate() {
	db.DB.AutoMigrate(&models.Torrent{})
}
