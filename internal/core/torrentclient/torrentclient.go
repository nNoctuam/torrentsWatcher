package torrentclient

type Client interface {
	AddTorrent(content []byte, dir string, paused bool) (Torrent, error)
	UpdateTorrent(url string, content []byte) error
	RemoveTorrents(ids []int, deleteLocalData bool) error
	GetTorrents() ([]Torrent, error)
	Rename(id int, oldPath string, newPath string) error
}
