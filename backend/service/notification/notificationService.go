package notification

import (
	"fmt"
	"os"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/utils"
	"gorm.io/gorm"
)

// NotificationService defines the orchestration layer for all platform alerts.
type NotificationService interface {
	SendVerificationEmail(user *models.User, token string) error
	SendTransactionalAlert(userID uint, title, content string) error
	BroadcastNewsletter(subject, content string) error
}

type notificationService struct {
	db *gorm.DB
}

// NewNotificationService initializes the notification uplink
func NewNotificationService(db *gorm.DB) NotificationService {
	return &notificationService{db}
}

// SendVerificationEmail triggers the transactional identity verification flow.
func (s *notificationService) SendVerificationEmail(user *models.User, token string) error {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000" // Fallback to local sector
	}

	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", frontendURL, token)
	
	// Generate the industrial HTML template
	body := utils.GenerateVerificationTemplate(user.FirstName, verificationLink)
	
	// Dispatch the relay
	return utils.SendMail([]string{user.Email}, "Confirm Identity Node", body)
}

// SendTransactionalAlert persists an in-app notification for the user.
func (s *notificationService) SendTransactionalAlert(userID uint, title, content string) error {
	notification := &models.Notification{
		UserID:   userID,
		Title:    title,
		Content:  content,
		Type:     models.TypeTransactional,
		Priority: models.PriorityNormal,
	}

	return s.db.Create(notification).Error
}

// BroadcastNewsletter sends a high-volume email to all active subscribers.
func (s *notificationService) BroadcastNewsletter(subject, content string) error {
	var subscribers []models.NewsletterSubscription
	if err := s.db.Where("is_active = ?", true).Find(&subscribers).Error; err != nil {
		return fmt.Errorf("failed to fetch sector subscribers: %v", err)
	}

	// Iterate through the subscribers and dispatch the broadcast
	// Note: For large volumes, we would use a worker-pool concurrency pattern
	for _, sub := range subscribers {
		utils.SendMail([]string{sub.Email}, subject, content)
	}

	return nil
}
