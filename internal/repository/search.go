package repository

import (
	"context"
	_ "embed"
	"encoding/json"
)

var (
	//go:embed sql/search.sql
	searchSQL string

	//go:embed sql/search.count.sql
	countSearchSQL string
)

type SearchResult struct {
	Messages []Message `json:"messages"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
}

func (s *Service) Search(ctx context.Context, to *string, subject *string, page int, rows int) (*SearchResult, error) {
	messages, err := s.searchMessages(ctx, to, subject, page, rows)
	if err != nil {
		return nil, err
	}

	total := len(messages)
	if total > 0 {
		result, err := s.countSearchMessages(ctx, to, subject)
		if err != nil {
			return nil, err
		}

		total = result
	}

	return &SearchResult{
		Messages: messages,
		Total:    total,
	}, nil
}

func (s *Service) countSearchMessages(ctx context.Context, to *string, subject *string) (int, error) {
	total := 0
	queryContext, err := s.conn.QueryContext(ctx, countSearchSQL, to, subject)
	if err != nil {
		return total, err
	}

	defer queryContext.Close()

	if queryContext.Next() {
		err = queryContext.Scan(&total)
		if err != nil {
			return total, err
		}

	}

	return total, nil
}

func (s *Service) searchMessages(
	ctx context.Context, to *string, subject *string, page int, rows int) ([]Message, error) {
	messages := make([]Message, 0)

	queryContext, err := s.conn.QueryContext(ctx, searchSQL, to, subject, page, rows, rows)
	if err != nil {
		return messages, err
	}

	defer queryContext.Close()

	for queryContext.Next() {
		var (
			message    Message
			customArgs string
			categories string
		)

		err := queryContext.Scan(
			&message.EventID,
			&message.MessageID,
			&message.ReceivedAt,
			&message.Subject,
			&message.From.Name,
			&message.From.Address,
			&message.To.Name,
			&message.To.Address,
			&message.Content.Html,
			&message.Content.Text,
			&customArgs,
			&categories,
		)
		if err != nil {
			return messages, err
		}

		err = json.Unmarshal([]byte(customArgs), &message.CustomArgs)
		if err != nil {
			return messages, err
		}

		err = json.Unmarshal([]byte(categories), &message.Categories)
		if err != nil {
			return messages, err
		}

		messages = append(messages, message)
	}

	return messages, err
}
