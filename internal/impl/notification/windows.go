package notification

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"torrentsWatcher/internal/core/notifications"
	"torrentsWatcher/internal/utils/shell"
)

type Windows struct {
	Config notifications.Config
}

func (n *Windows) GetConfig() notifications.Config {
	return n.Config
}

func (n *Windows) OpenInBrowser(url string) {
	output, err := shell.TryExec("start", `""`, fmt.Sprintf(`"%s"`, url))
	if err != nil {
		fmt.Print(err, output)
	}
}

func (n *Windows) SendMessage(messageType notifications.MessageType, text string) {
	log.Fatal("notification.Windows.SendMessage() is not implemented yet.")
}

func (n *Windows) OpenFile(content []byte, name string) {
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

	output, err := shell.TryExec("start", `""`, fmt.Sprintf(`"%s"`, filename))
	if err != nil {
		fmt.Print(err, output)
	}
}
