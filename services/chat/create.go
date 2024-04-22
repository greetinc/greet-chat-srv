package services

import (
	dto "greet-chat-srv/dto/chat"
	"greet-chat-srv/entity"
)

func (h *chatService) SendMessage(msg *entity.Message) error {
	req := dto.ChatRequest{
		ID:         msg.ID,
		Content:    msg.Content,
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
	}

	_, err := h.ChatR.StoreMessage(req)
	return err
}
