package torrentclient

import "os"

type TorrentClient struct {
	autoDownloadDir string
	client          Client
}

type Client interface {
	AddTorrent(content []byte, dir string) error
}

func New(autoDownloadDir string, client Client) *TorrentClient {
	return &TorrentClient{
		autoDownloadDir: autoDownloadDir,
		client:          client,
	}
}

func (c *TorrentClient) SaveToAutoDownloadFolder(name string, content []byte) error {
	return os.WriteFile(c.autoDownloadDir+"/"+name, content, 0660)
}

func (c *TorrentClient) AddTorrent(content []byte, dir string) error {
	return c.client.AddTorrent(content, dir)
}
