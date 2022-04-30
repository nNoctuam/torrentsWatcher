package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"torrentsWatcher/internal/core/models"
	"torrentsWatcher/internal/core/storage"
	"torrentsWatcher/internal/core/torrentclient"
	storageImpl "torrentsWatcher/internal/impl/storage"
	torrentClientImpl "torrentsWatcher/internal/impl/torrentclient"
	trackingImpl "torrentsWatcher/internal/impl/tracker"
	"torrentsWatcher/internal/services/tracking"
	"torrentsWatcher/internal/services/watcher"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"

	"torrentsWatcher/config"

	"go.uber.org/zap"

	grpc_server "torrentsWatcher/internal/interfaces/grpc/server"
)

const portHTTP = 10000

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	errorChan := make(chan error)
	wg := new(sync.WaitGroup)

	cfg, err := config.Load("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	logger, err := newLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("sqlite3", "./torrents.db")
	if err != nil {
		logger.Fatal("open db", zap.Error(err))
		os.Exit(1)
	}

	defer db.Close()
	db.AutoMigrate(&models.Torrent{}, &models.AuthCookie{}, &models.TransmissionTorrent{})

	torrentsStorage := storageImpl.NewTorrentsSqliteStorage(db)
	cookiesStorage := storageImpl.NewCookiesSqliteStorage(db)

	trackers := tracking.Trackers([]*tracking.Tracker{
		trackingImpl.NewNnmClub(logger, cfg.Credentials[trackingImpl.NnmClubDomain], torrentsStorage, cookiesStorage),
		trackingImpl.NewRutracker(logger, cfg.Credentials[trackingImpl.RutrackerDomain], torrentsStorage, cookiesStorage),
		trackingImpl.NewKinozal(logger, cfg.Credentials[trackingImpl.KinozalDomain], torrentsStorage, cookiesStorage),
	})
	for i, t := range trackers {
		if t.Credentials.Login == "" {
			trackers = append(trackers[:i], trackers[i+1:]...)
		}
	}

	transmissionClient, err := torrentClientImpl.NewTransmission(
		cfg.AutoDownloadDir,
		cfg.Transmission.RPCURL,
		cfg.Transmission.Login,
		cfg.Transmission.Password,
	)
	if err != nil {
		logger.Fatal("make transmission client", zap.Error(err))
		_ = db.Close()
		// nolint: gocritic
		os.Exit(1)
	}

	wg.Add(1)
	go watcher.New(ctx, wg, logger, cfg.Interval, trackers, transmissionClient, torrentsStorage).Run()

	httpServer := serveHTTP(errorChan, logger, portHTTP)
	go serveRPC(logger.Named("RPC"), httpServer, trackers, torrentsStorage, cfg.Transmission.Folders, transmissionClient)

	logger.Info("Service started")
	select {
	case err := <-errorChan:
		logger.Panic("Service crashed", zap.Error(err))
	case <-ctx.Done():
		logger.Info("Service context stopped")
	case <-waitExitSignal():
		logger.Info("Service stopped by signal")
	}

	ctxCancel()
	wg.Wait()
}

func serveRPC(
	logger *zap.Logger,
	mainHTTPServer *http.Server,
	trackers tracking.Trackers,
	torrentsStorage storage.Torrents,
	downloadFolders map[string]string,
	torrentClient torrentclient.Client,
) {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcServer.RegisterService(&grpc_server.BaseServiceDesc, grpc_server.NewRPCServer(
		logger,
		trackers,
		torrentsStorage,
		downloadFolders,
		torrentClient,
	))
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	mainHandler := mainHTTPServer.Handler
	mainHTTPServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			fmt.Println("got gRPC request: " + req.Method + " " + req.URL.String())
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}

		if wrappedGrpc.IsAcceptableGrpcCorsRequest(req) {
			fmt.Println("got AcceptableGrpcCorsRequest request: " + req.Method + " " + req.URL.String())
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			resp.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-Grpc-Web, X-User-Agent")
			resp.WriteHeader(http.StatusNoContent)
			return
		}

		mainHandler.ServeHTTP(resp, req)
	})
	fmt.Printf("Attached gRPC to the http server\n")
}

func serveHTTP(
	errorChan chan error,
	logger *zap.Logger,
	port int,
) *http.Server {
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

	router.Handle("/*", http.FileServer(http.Dir("dist")))

	router.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("test response"))
	}))

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: router,
	}

	go func() {
		logger.Info("Webserver start", zap.String("host", "http://"+server.Addr))
		errorChan <- server.ListenAndServe()
	}()
	return server
}

func waitExitSignal() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	return ch
}

func newLogger(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	atom := zap.NewAtomicLevel()
	err := atom.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}

	cfg.Level = atom

	return cfg.Build()
}
