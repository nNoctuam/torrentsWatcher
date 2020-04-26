package main

import (
	"fmt"
	"torrentsWatcher/parsing"
)

// tracker parsing
// adding torrent for watching - db
// storing auth cookies
// notifications
// cron to check torrents
// authorization

func main() {
	fmt.Print("it works.\n")
	parsing.GetTorrentInfo("http://nnmclub.to/forum/viewtopic.php?p=10307544")
}
