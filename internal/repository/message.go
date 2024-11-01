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
	CustomArgs CustomArgs
	Categories Categories
}

type Recipient struct {
	Name    string
	Address string
}

type Content struct {
	Html *string
	Text *string
}

type CustomArgs map[string]string

type Categories []string
