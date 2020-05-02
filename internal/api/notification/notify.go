package notification

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"torrentsWatcher/config"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tools"
)

var Notificator notificator

type notificator interface {
	openInBrowser(url string)
	sendTrayMessage(text string)
	sendKDEMessage(text string)
	openFile(content []byte, name string)
}

type Linux struct{}
type Windows struct{}

func NotifyAbout(torrent *models.Torrent) {
	if config.App.Notifications.TrayMessage {
		Notificator.sendTrayMessage(torrent.Title)
	}
	if config.App.Notifications.KDEMessage {
		Notificator.sendKDEMessage(torrent.Title)
	}
	if config.App.Notifications.OpenFile && torrent.File != nil {
		Notificator.openFile(torrent.File, "tmp.torrent")
	} else {
		if config.App.Notifications.OpenInBrowser {
			Notificator.openInBrowser(torrent.PageUrl)
		}
	}
}

func (n *Linux) openInBrowser(url string) {
	output, err := tools.TryExecShell("xdg-open", url)
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Linux) sendTrayMessage(text string) {
	output, err := tools.TryExecShell("notify-send", "-a", "torrentsWatcher", text)
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Linux) sendKDEMessage(text string) {
	output, err := tools.TryExecShell("kdialog", "--passivepopup", text, "--title", "Torrent was updated", "300")
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Linux) openFile(content []byte, name string) {
	filename := os.TempDir() + string(filepath.Separator) + name
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.Write(content)
	if err != nil {
		return
	}
	defer f.Close()

	output, err := tools.TryExecShell("xdg-open", filename)
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Windows) openInBrowser(url string) {
	output, err := tools.TryExecShell("start", `""`, fmt.Sprintf(`"%s"`, url))
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Windows) sendTrayMessage(text string) {
	log.Fatal("'sendTrayMessage' is not available for windows")
}

func (n *Windows) sendKDEMessage(text string) {
	log.Fatal("'sendKDEMessage' is not available for windows")
}

func (n *Windows) openFile(content []byte, name string) {
	filename := os.TempDir() + string(filepath.Separator) + name
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.Write(content)
	if err != nil {
		return
	}
	defer f.Close()

	output, err := tools.TryExecShell("start", `""`, fmt.Sprintf(`"%s"`, filename))
	if err != nil {
		fmt.Print(err, output)
	}
}
