package services

import (
	dto "greet-chat-srv/dto/chat"
)

func (s *chatService) GetAll(sender_id, receiver_id string) ([]dto.ChatResponse, error) {
	datas, err := s.ChatR.GetMessagesForUser(sender_id, receiver_id)
	if err != nil {
		return nil, err
	}

	var responses []dto.ChatResponse
	for _, msg := range datas {
		response := dto.ChatResponse{
			ID:           msg.ID,
			Content:      msg.Content,
			SenderID:     msg.SenderID,
			ReceiverID:   msg.ReceiverID,
			SenderName:   msg.SenderName,
			ReceiverName: msg.ReceiverName,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
