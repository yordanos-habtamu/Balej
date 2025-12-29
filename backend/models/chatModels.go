package models

import "time"

type Chat struct {
	ID        string     `json:"id" gorm:"primaryKey;"`
	SenderID  uint        `json:"sender_id" gorm:"not null"`
	ReceiverID uint        `json:"reciver_id" gorm:"not null"`
	Messages  []Message  `json:"messages" gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type Message struct {
	ID        string    `json:"id" gorm:"primaryKey;"`
	ChatID    string    `json:"chat_id" gorm:"type:uuid;not null"` // Foreign key
	SenderID  int       `json:"sender_id" gorm:"not null"`         // Optional: for identifying who sent
	Content   string    `json:"content" gorm:"type:text;not null"`
	SendAt    time.Time `json:"send_at" gorm:"autoCreateTime"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
}
