package repository

import "github.com/BoneCrusher9/openmessenger/backend/internal/domain"

type MessageRepository interface {
	Create(msg *domain.Message) error
	GetByID(id string) (*domain.Message, error)
	GetChatMessages(chatID string, limit, offset int) ([]*domain.Message, error)
	Delete(id string) error
}
