package domain

import "time"

type ChatType string

const (
	ChatTypePrivate ChatType = "private"
	ChatTypeGroup   ChatType = "group"
)

type Chat struct {
	ID        string
	Type      ChatType
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
