package domain

import "time"

type User struct {
	ID           string
	Username     string
	DisplayName  string
	Email        string
	PasswordHash string

	AvatarURL string
	About     string

	CreatedAt time.Time
	UpdatedAt time.Time
}
