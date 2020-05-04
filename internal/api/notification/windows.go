package notification

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"torrentsWatcher/internal/api/utils/shell"
)

type Windows struct {
	Config Config
}

func (n *Windows) getConfig() Config {
	return n.Config
}

func (n *Windows) openInBrowser(url string) {
	output, err := shell.TryExec("start", `""`, fmt.Sprintf(`"%s"`, url))
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

	output, err := shell.TryExec("start", `""`, fmt.Sprintf(`"%s"`, filename))
	if err != nil {
		fmt.Print(err, output)
	}
}
