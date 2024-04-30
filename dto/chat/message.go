package dto

type ChatRequest struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	SenderID     string `json:"sender_id"`
	ReceiverID   string `json:"receiver_id"`
	SenderName   string `json:"sender_name"`
	ReceiverName string `json:"receiver_name"`
}

type ChatResponse struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	SenderID     string `json:"sender_id"`
	ReceiverID   string `json:"receiver_id"`
	SenderName   string `json:"sender_name"`
	ReceiverName string `json:"receiver_name"`
}
