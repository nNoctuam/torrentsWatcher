package notifications

import (
	"torrentsWatcher/internal/core/models"
)

func NotifyAbout(torrent *models.Torrent, notificator Notificator) {
	notificator.SendMessage(notificator.GetConfig().Message, torrent.Title)
	if notificator.GetConfig().OpenFile && torrent.File != nil {
		notificator.OpenFile(torrent.File, "tmp.torrent")
	} else {
		if notificator.GetConfig().OpenInBrowser {
			notificator.OpenInBrowser(torrent.PageUrl)
		}
	}
}
