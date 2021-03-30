package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"torrentsWatcher/internal/api/torrentclient"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"

	"torrentsWatcher/config"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/tracking"
	trackingImpl "torrentsWatcher/internal/api/tracking/impl"
	"torrentsWatcher/internal/api/watcher"
	"torrentsWatcher/internal/handlers"
	"torrentsWatcher/internal/storage"
	storageImpl "torrentsWatcher/internal/storage/impl"
)

// TODO:
//  login for search
//  no ignored errors
// 	log instead of fmt.Print
// 	DI
// 	docker build
// 	supervisor config
// 	unit tests
// 	notifications:
// 		browser
// 		messenger
// 		email

//go:embed dist/*
var distContent embed.FS

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	errorChan := make(chan error)
	wg := new(sync.WaitGroup)

	cfg := config.Load("./config.yml")
	notificator := getNotificator(cfg)

	db, err := gorm.Open("sqlite3", "./torrents.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.Torrent{}, &models.AuthCookie{})

	torrentsStorage := storageImpl.NewTorrentsSqliteStorage(db)
	cookiesStorage := storageImpl.NewCookiesSqliteStorage(db)

	trackers := tracking.Trackers([]*tracking.Tracker{
		trackingImpl.NewNnmClub(cfg.Credentials[trackingImpl.NnmClubDomain], torrentsStorage, cookiesStorage),
		trackingImpl.NewRutracker(cfg.Credentials[trackingImpl.RutrackerDomain], torrentsStorage, cookiesStorage),
		trackingImpl.NewKinozal(cfg.Credentials[trackingImpl.KinozalDomain], torrentsStorage, cookiesStorage),
	})
	for i, t := range trackers {
		if t.Credentials.Login == "" {
			trackers = append(trackers[:i], trackers[i+1:]...)
		}
	}

	wg.Add(1)
	go watcher.New(ctx, wg, cfg.Interval, trackers, notificator, torrentsStorage).Run()

	torrentClient := torrentclient.New(cfg.AutoDownloadDir)
	serve(errorChan, cfg.Host, cfg.Port, trackers, torrentsStorage, torrentClient)

	fmt.Println("Service started")
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

func serve(
	errorChan chan error,
	host string,
	port string,
	trackers tracking.Trackers,
	torrentsStorage storage.Torrents,
	torrentClient *torrentclient.TorrentClient,
) {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/torrents", handlers.GetTorrents(torrentsStorage))
	router.MethodFunc("POST", "/torrent", handlers.AddTorrent(trackers, torrentsStorage))
	router.MethodFunc("POST", "/search", handlers.Search(trackers))
	router.MethodFunc("POST", "/download", handlers.DownloadWithClient(trackers, torrentClient))
	router.MethodFunc("DELETE", `/torrent/{id:\d+}`, handlers.DeleteTorrent(torrentsStorage))

	content, _ := fs.Sub(distContent, "dist")
	router.Handle("/*", http.FileServer(http.FS(content)))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	go func() {
		errorChan <- server.ListenAndServe()
	}()
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
