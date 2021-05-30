package notifications

import (
	"torrentsWatcher/internal/core/models"
)

func NotifyAbout(torrent *models.Torrent, notificator Notificator) {
	if notificator.GetConfig().TrayMessage {
		notificator.SendMessage(MessageTypeTray, torrent.Title)
	}
	if notificator.GetConfig().KDEMessage {
		notificator.SendMessage(MessageTypeKDE, torrent.Title)
	}
	if notificator.GetConfig().OpenFile && torrent.File != nil {
		notificator.OpenFile(torrent.File, "tmp.torrent")
	} else {
		if notificator.GetConfig().OpenInBrowser {
			notificator.OpenInBrowser(torrent.PageUrl)
		}
	}
}
