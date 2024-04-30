package entity

type Message struct {
	ID         string `gorm:"primary_key" json:"id"`
	Content    string `gorm:"content" json:"content"`
	SenderID   string `gorm:"sender_id" json:"sender_id"`
	ReceiverID string `gorm:"receiver_id" json:"receiver_id"`
	RoomID     string `gorm:"type:varchar(36);index" json:"room_id"`
}
