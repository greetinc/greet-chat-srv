package services

import (
	dto "greet-chat-srv/dto/chat"
	"greet-chat-srv/entity"
	repositories "greet-chat-srv/repositories/chat"

	m "github.com/greetinc/greet-middlewares/middlewares"
)

type ChatService interface {
	SendMessage(msg *entity.Message) error
	GetAll(sender_id, receiver_id string) ([]dto.ChatResponse, error)
}

type chatService struct {
	ChatR repositories.ChatRepository
	jwt   m.JWTService
}

func NewChatService(ChatR repositories.ChatRepository, jwtS m.JWTService) ChatService {
	return &chatService{
		ChatR: ChatR,
		jwt:   jwtS,
	}
}
