package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"sync"
	"syscall"

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
	ctx, ctxCancel := context.WithCancel(context.Background())
	errorChan := make(chan error)
	wg := new(sync.WaitGroup)

	basePath := path.Dir(os.Args[0])

	cfg := config.Load(basePath + "/config.yml")
	notificator := getNotificator(cfg)
	parsers := []*parser.Tracker{
		impl.NewNnmClub(cfg.Credentials[impl.NnmClubDomain]),
		impl.NewRutracker(cfg.Credentials[impl.RutrackerDomain]),
	}

	db.InitDB(basePath + "/torrents.db")
	defer db.CloseDB()
	migrate()

	fmt.Println("Service started")

	wg.Add(1)
	go watch.Run(ctx, wg, cfg.Period, parsers, notificator)
	serve(errorChan, cfg.Host, cfg.Port, basePath, parsers)

	select {
	case err := <-errorChan:
		fmt.Println(err)
	case <-ctx.Done():
		fmt.Println("Service context stopped")
	case <-waitExitSignal():
		fmt.Println("Service stopped by signal")
	}

	ctxCancel()
	wg.Wait()
}

func serve(errorChan chan error, host string, port string, basePath string, parsers []*parser.Tracker) {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/torrents", handlers.GetTorrents)
	router.MethodFunc("POST", "/torrent", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTorrent(w, r, parsers)
	})
	router.MethodFunc("GET", `/torrent/{id:\d+}/download`, handlers.DownloadTorrent)
	router.MethodFunc("DELETE", `/torrent/{id:\d+}`, handlers.DeleteTorrent)
	router.Handle("/*", http.FileServer(http.Dir(basePath+"/frontend/dist")))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	go func() {
		errorChan <- server.ListenAndServe()
	}()
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

func waitExitSignal() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	return ch
}
