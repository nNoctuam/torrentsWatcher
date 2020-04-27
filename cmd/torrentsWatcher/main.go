package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"torrentsWatcher/internal/api"
	"torrentsWatcher/internal/handlers"
)

// tracker parsing
// adding torrent for watching - db
// storing auth cookies
// notifications
// cron to check torrents
// authorization

func main() {
	api.InitDB()
	defer api.CloseDB()

	fmt.Print("it works.\n")

	router := chi.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello")
	})
	router.HandleFunc("/torrents", handlers.GetTorrents)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	server.ListenAndServe()
	//torrent, err := parsing.GetTorrentInfo("http://nnmclub.to/forum/viewtopic.php?p=10307544")
	//err = DB.Create(&torrent).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
}
