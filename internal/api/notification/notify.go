package notification

import (
	"torrentsWatcher/internal/api/models"
)

type Config struct {
	TrayMessage   bool
	KDEMessage    bool
	OpenInBrowser bool
	OpenFile      bool
}

type Notificator interface {
	openInBrowser(url string)
	sendTrayMessage(text string)
	sendKDEMessage(text string)
	openFile(content []byte, name string)
	getConfig() Config
}

func NotifyAbout(torrent *models.Torrent, notificator Notificator) {
	if notificator.getConfig().TrayMessage {
		notificator.sendTrayMessage(torrent.Title)
	}
	if notificator.getConfig().KDEMessage {
		notificator.sendKDEMessage(torrent.Title)
	}
	if notificator.getConfig().OpenFile && torrent.File != nil {
		notificator.openFile(torrent.File, "tmp.torrent")
	} else {
		if notificator.getConfig().OpenInBrowser {
			notificator.openInBrowser(torrent.PageUrl)
		}
	}
}
