package messaging

type messagePlatform interface {
	SendMessage(message string, useriID string)
	ProcessMessage(message string)
}
