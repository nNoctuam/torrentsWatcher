package notification

import (
	"fmt"
	"os"
	"path/filepath"
	"torrentsWatcher/internal/core/notifications"
	"torrentsWatcher/internal/utils/shell"
)

type Linux struct {
	Config notifications.Config
}

func (n *Linux) GetConfig() notifications.Config {
	return n.Config
}

func (n *Linux) OpenInBrowser(url string) {
	output, err := shell.TryExec("xdg-open", url)
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Linux) SendMessage(messageType notifications.MessageType, text string) {
	switch messageType {
	case notifications.MessageTypeTray:
		n.sendTrayMessage(text)
	case notifications.MessageTypeKDE:
		n.sendKDEMessage(text)
	}
}

func (n *Linux) sendTrayMessage(text string) {
	output, err := shell.TryExec("notify-send", "-a", "torrentsWatcher", text)
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Linux) sendKDEMessage(text string) {
	output, err := shell.TryExec("kdialog", "--passivepopup", text, "--title", "Torrent was updated", "300")
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Linux) OpenFile(content []byte, name string) {
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

	output, err := shell.TryExec("xdg-open", filename)
	if err != nil {
		fmt.Print(err, output)
	}
}
