package repository

import "time"

type Message struct {
	EventID    string     `json:"event_id"`
	MessageID  string     `json:"message_id"`
	ReceivedAt time.Time  `json:"received_at"`
	Subject    string     `json:"subject"`
	From       Recipient  `json:"from"`
	To         Recipient  `json:"to"`
	Content    Content    `json:"content"`
	CustomArgs CustomArgs `json:"custom_args"`
	Categories Categories `json:"categories"`
}

type Recipient struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Content struct {
	Html *string `json:"html"`
	Text *string `json:"text"`
}

type CustomArgs map[string]string

type Categories []string
