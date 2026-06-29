package repository

import "github.com/BoneCrusher9/openmessenger/backend/internal/domain"

type ChatRepository interface {
	Create(chat *domain.Chat) error
	GetByID(id string) (*domain.Chat, error)
	GetUserChats(userID string) ([]*domain.Chat, error)
	AddUserToChat(chatID, userID string) error
	RemoveUserFromChat(chatID, userID string) error
}
