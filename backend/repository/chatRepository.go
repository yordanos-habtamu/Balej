package repository

import (
	"errors"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/utils"
	"gorm.io/gorm"
)

type ChatRepository interface {
	Create(senderID uint, receiverID uint) (*models.Chat, error)
	GetByID(chatID string) (*models.Chat, error)
    SendMessage(chatID string, message *models.Message , recieverId int, senderId int, isFile bool, fileUrl string, fileName string) error
	GetChatsByUserID(userID uint) ([]*models.Chat, error)
	GetMessagesByChatID(chatID string) ([]*models.Message, error)
}

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepo{db}
}

// Create creates a new chat only if it doesn't already exist
func (r *chatRepo) Create(senderID uint, receiverID uint) (*models.Chat, error) {
	chatID := utils.GenerateChatID(int(senderID), int(receiverID))

	// Check if the chat already exists
	var existing models.Chat
	if err := r.db.First(&existing, "id = ?", chatID).Error; err == nil {
		return &existing, nil // Already exists
	}

	// Doesn't exist, so create
	chat := &models.Chat{
		ID:         chatID,
		SenderID:   uint(senderID),
		ReceiverID: uint(receiverID),
		Messages:[] models.Message{},
	}

	if err := r.db.Create(chat).Error; err != nil {
		return nil, err
	}

	return chat, nil
}

func (r *chatRepo) GetByID(chatID string) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.Preload("Messages").First(&chat, "id = ?", chatID).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *chatRepo) SendMessage(chatID string, message *models.Message, senderId int, recieverId int, isFile bool, fileUrl string, fileName string) error {
	message.IsFile = isFile
	message.FileUrl = fileUrl
	message.FileName = fileName

	// Find the chat by ID
	var chat models.Chat

	// Try to find chat
	err := r.db.Preload("Messages").First(&chat, "id = ?", chatID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new chat
		chat = models.Chat{
			ID:         chatID,
			SenderID:   uint(message.SenderID),
			ReceiverID: uint(recieverId),
			Messages:   []models.Message{*message},
			CreatedAt:  time.Now(),
		}
		return r.db.Create(&chat).Error
	} else if err != nil {
		return err
	}

	// Append message to existing chat
	chat.Messages = append(chat.Messages, *message)
	return r.db.Save(&chat).Error
}

func (r *chatRepo) GetChatsByUserID(userID uint) ([]*models.Chat, error) {
	var chats []*models.Chat
	// Find chats where the user is either the sender or the receiver
	err := r.db.Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Order("updated_at desc").
		Find(&chats).Error
	return chats, err
}

func (r *chatRepo) GetMessagesByChatID(chatID string) ([]*models.Message, error) {
	var messages []*models.Message
	err := r.db.Where("chat_id = ?", chatID).Order("send_at asc").Find(&messages).Error
	return messages, err
}

