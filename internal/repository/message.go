package repository

import "time"

type Message struct {
	EventID    string
	MessageID  string
	ReceivedAt time.Time
	Subject    string
	From       Recipient
	To         Recipient
	Content    Content
}

type Recipient struct {
	Name    string
	Address string
}

type Content struct {
	Html *string
	Text *string
}
