package notifications

type Config struct {
	OpenInBrowser bool
	OpenFile      bool
	Message       map[string]bool
}

const (
	MessageTypeTray string = "Tray"
	MessageTypeKDE  string = "KDE"
)

type Notificator interface {
	OpenInBrowser(url string)
	SendMessage(messageTypes map[string]bool, text string)
	OpenFile(content []byte, name string)
	GetConfig() Config
}
