package notification

import (
	"fmt"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tools"
)

func NotifyAbout(torrent *models.Torrent) {
	sendKDENotification(torrent.Title)
	openInBrowser(torrent.PageUrl)
}

func openInBrowser(url string) {
	output, err := tools.TryExecShell("xdg-open", url)
	if err != nil {
		fmt.Print(err, output)
	}
}

func sendNotification(text string) {
	output, err := tools.TryExecShell("notify-send", "-a", "torrentsWatcher", text)
	if err != nil {
		fmt.Print(err, output)
	}
}

func sendKDENotification(text string) {
	output, err := tools.TryExecShell("kdialog", "--passivepopup", text, "--title", "Torrent was updated", "300")
	if err != nil {
		fmt.Print(err, output)
	}
}
