package notifications

type Config struct {
	TrayMessage   bool
	KDEMessage    bool
	OpenInBrowser bool
	OpenFile      bool
}

type MessageType int

const (
	MessageTypeTray MessageType = iota
	MessageTypeKDE
)

type Notificator interface {
	OpenInBrowser(url string)
	SendMessage(messageType MessageType, text string)
	OpenFile(content []byte, name string)
	GetConfig() Config
}
