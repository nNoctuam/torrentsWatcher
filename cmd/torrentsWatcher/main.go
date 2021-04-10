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
	"torrentsWatcher/internal/api/torrentclient/impl"
	"torrentsWatcher/internal/api/watcher"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"

	"torrentsWatcher/config"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/tracking"
	trackingImpl "torrentsWatcher/internal/api/tracking/impl"
	"torrentsWatcher/internal/handlers"
	"torrentsWatcher/internal/storage"
	storageImpl "torrentsWatcher/internal/storage/impl"
)

// TODO:
//  no ignored errors
// 	log instead of fmt.Print
// 	unit tests
//  auto replace being monitored torrents

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
	db.AutoMigrate(&models.Torrent{}, &models.AuthCookie{}, &models.TransmissionTorrent{})

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

	transmissionClient, err := impl.NewTransmission(cfg.Transmission.RpcUrl, cfg.Transmission.Login, cfg.Transmission.Password)
	if err != nil {
		log.Fatal("err")
	}
	torrentClient := torrentclient.New(cfg.AutoDownloadDir, transmissionClient)

	serve(errorChan, cfg.Host, cfg.Port, trackers, torrentsStorage, torrentClient, cfg.Transmission.Folders)

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
	downloadFolders map[string]string,
) {
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

	router.MethodFunc("GET", "/download-folders", handlers.GetDownloadFolders(downloadFolders))
	router.MethodFunc("GET", "/torrents", handlers.GetTorrents(torrentsStorage))
	router.MethodFunc("GET", "/transmission-torrents", handlers.GetTransmissionTorrents(torrentsStorage, torrentClient))
	router.MethodFunc("POST", "/torrent", handlers.AddTorrent(trackers, torrentsStorage))
	router.MethodFunc("POST", "/search", handlers.Search(trackers))
	router.MethodFunc("POST", "/download", handlers.DownloadWithClient(trackers, torrentClient, torrentsStorage, downloadFolders))
	router.MethodFunc("DELETE", `/torrent/{id:\d+}`, handlers.DeleteTorrent(torrentsStorage))
	router.MethodFunc("POST", `/rename`, handlers.Rename(torrentClient))

	content, _ := fs.Sub(distContent, "dist")
	router.Handle("/*", http.FileServer(http.FS(content)))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	go func() {
		fmt.Printf("serving at http://%s\n", server.Addr)
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
