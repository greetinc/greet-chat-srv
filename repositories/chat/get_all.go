package repositories

import (
	dto "greet-chat-srv/dto/chat"
	"greet-chat-srv/entity"
)

func (r *chatRepository) GetMessagesForUser(sender_id, receiver_id string) ([]dto.ChatResponse, error) {
	var chatResponses []dto.ChatResponse

	// Mengambil pesan yang dikirim oleh sender ke receiver dan sebaliknya
	// Anda mungkin perlu menyesuaikan query di bawah ini sesuai dengan struktur tabel Anda
	err := r.DB.Model(&entity.Message{}).
		Select("messages.*, sender.full_name as sender_name, receiver.full_name as receiver_name").
		Joins("JOIN user_details as sender ON sender.user_id = messages.sender_id").
		Joins("JOIN user_details as receiver ON receiver.user_id = messages.receiver_id").
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", sender_id, receiver_id, receiver_id, sender_id).
		Scan(&chatResponses).Error
	if err != nil {
		return nil, err
	}

	return chatResponses, nil
}
