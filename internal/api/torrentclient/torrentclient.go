package torrentclient

import "os"

type TorrentClient struct {
	autoDownloadDir string
}

func New(autoDownloadDir string) *TorrentClient {
	return &TorrentClient{autoDownloadDir: autoDownloadDir}
}

func (c *TorrentClient) StartDownload(name string, content []byte) error {
	return os.WriteFile(c.autoDownloadDir+"/"+name, content, 0660)
}
