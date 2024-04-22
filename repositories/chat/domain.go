package repositories

import (
	dto "greet-chat-srv/dto/chat"

	"gorm.io/gorm"
)

type ChatRepository interface {
	StoreMessage(req dto.ChatRequest) (*dto.ChatResponse, error)
	GetMessagesForUser(sender_id, receiver_id string) ([]dto.ChatResponse, error)
}

type chatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(DB *gorm.DB) ChatRepository {
	return &chatRepository{
		DB: DB,
	}
}
