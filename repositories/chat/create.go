package repositories

import (
	dto "greet-chat-srv/dto/chat"
	"greet-chat-srv/entity"
)

func (b *chatRepository) StoreMessage(req dto.ChatRequest) (*dto.ChatResponse, error) {
	chat := entity.Message{
		ID:         req.ID,
		Content:    req.Content,
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
	}

	if err := b.DB.Create(&chat).Error; err != nil {
		return nil, err
	}

	response := &dto.ChatResponse{
		ID:         chat.ID,
		Content:    chat.Content,
		SenderID:   chat.SenderID,
		ReceiverID: chat.ReceiverID,
	}

	return response, nil

}
