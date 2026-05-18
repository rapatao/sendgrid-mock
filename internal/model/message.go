package model

import "time"

type Message struct {
	EventID     string       `json:"event_id"`
	MessageID   string       `json:"message_id"`
	ReceivedAt  time.Time    `json:"received_at"`
	Subject     string       `json:"subject"`
	From        Recipient    `json:"from"`
	To          Recipient    `json:"to"`
	Content     Content      `json:"content"`
	CustomArgs  CustomArgs   `json:"custom_args"`
	Categories  Categories   `json:"categories"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Content     string `json:"content"`
	Filename    string `json:"filename"`
	Type        string `json:"type"`
	Disposition string `json:"disposition"`
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
