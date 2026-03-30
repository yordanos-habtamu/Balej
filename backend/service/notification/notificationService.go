package notification

import (
	"fmt"
	"os"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/utils"
	"gorm.io/gorm"
)

// NotificationService defines the orchestration layer for all platform alerts.
type NotificationService interface {
	SendVerificationEmail(user *models.User, token string) error
	SendTransactionalAlert(userID uint, title, content string) error
	BroadcastNewsletter(subject, content string) error
	GetUserNotifications(userID uint) ([]*models.Notification, error)
	MarkNotificationRead(id uint, read bool) (*models.Notification, error)
	CreateNotification(notification *models.Notification) (*models.Notification, error)
	// Email relay config
	ListSenderProfiles() ([]*models.EmailSenderProfile, error)
	CreateSenderProfile(profile *models.EmailSenderProfile) (*models.EmailSenderProfile, error)
	UpdateSenderProfile(id uint, updates *models.EmailSenderProfile) (*models.EmailSenderProfile, error)
	DeleteSenderProfile(id uint) (bool, error)
	GetRoutingConfig() (*models.EmailRoutingConfig, error)
	UpdateRoutingConfig(config *models.EmailRoutingConfig) (*models.EmailRoutingConfig, error)
	// Templates
	ListEmailTemplates() ([]*models.EmailTemplate, error)
	CreateEmailTemplate(tpl *models.EmailTemplate) (*models.EmailTemplate, error)
	UpdateEmailTemplate(id uint, updates *models.EmailTemplate) (*models.EmailTemplate, error)
	DeleteEmailTemplate(id uint) (bool, error)
	// Announcements
	ListAnnouncements() ([]*models.Announcement, error)
	CreateAnnouncement(announcement *models.Announcement) (*models.Announcement, error)
	// Newsletters
	ListNewsletters() ([]*models.Newsletter, error)
	CreateNewsletter(newsletter *models.Newsletter) (*models.Newsletter, error)
	UpdateNewsletter(id uint, updates *models.Newsletter) (*models.Newsletter, error)
	ScheduleNewsletter(id uint, when time.Time) (*models.Newsletter, error)
	MarkNewsletterSent(id uint) (*models.Newsletter, error)
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

// GetUserNotifications returns notifications for a user, newest first.
func (s *notificationService) GetUserNotifications(userID uint) ([]*models.Notification, error) {
	var notifications []*models.Notification
	if err := s.db.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

// MarkNotificationRead toggles read state for a notification.
func (s *notificationService) MarkNotificationRead(id uint, read bool) (*models.Notification, error) {
	var notification models.Notification
	if err := s.db.First(&notification, id).Error; err != nil {
		return nil, err
	}
	notification.IsRead = read
	if err := s.db.Save(&notification).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

// CreateNotification persists a new notification.
func (s *notificationService) CreateNotification(notification *models.Notification) (*models.Notification, error) {
	if err := s.db.Create(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

// ListSenderProfiles returns all sender profiles.
func (s *notificationService) ListSenderProfiles() ([]*models.EmailSenderProfile, error) {
	var profiles []*models.EmailSenderProfile
	if err := s.db.Order("created_at desc").Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

// CreateSenderProfile inserts a new sender profile.
func (s *notificationService) CreateSenderProfile(profile *models.EmailSenderProfile) (*models.EmailSenderProfile, error) {
	if err := s.db.Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

// UpdateSenderProfile updates an existing sender profile.
func (s *notificationService) UpdateSenderProfile(id uint, updates *models.EmailSenderProfile) (*models.EmailSenderProfile, error) {
	var profile models.EmailSenderProfile
	if err := s.db.First(&profile, id).Error; err != nil {
		return nil, err
	}
	if updates.Label != "" {
		profile.Label = updates.Label
	}
	if updates.Email != "" {
		profile.Email = updates.Email
	}
	if updates.Password != "" {
		profile.Password = updates.Password
	}
	if updates.APIKey != "" {
		profile.APIKey = updates.APIKey
	}
	if updates.Provider != "" {
		profile.Provider = updates.Provider
	}
	profile.IsActive = updates.IsActive
	if err := s.db.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// DeleteSenderProfile deletes a sender profile.
func (s *notificationService) DeleteSenderProfile(id uint) (bool, error) {
	if err := s.db.Delete(&models.EmailSenderProfile{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetRoutingConfig returns the singleton routing configuration.
func (s *notificationService) GetRoutingConfig() (*models.EmailRoutingConfig, error) {
	var config models.EmailRoutingConfig
	err := s.db.First(&config).Error
	if err != nil {
		// create default config
		config = models.EmailRoutingConfig{}
		if err := s.db.Create(&config).Error; err != nil {
			return nil, err
		}
	}
	return &config, nil
}

// UpdateRoutingConfig updates routing configuration.
func (s *notificationService) UpdateRoutingConfig(config *models.EmailRoutingConfig) (*models.EmailRoutingConfig, error) {
	var existing models.EmailRoutingConfig
	if err := s.db.First(&existing).Error; err != nil {
		if err := s.db.Create(config).Error; err != nil {
			return nil, err
		}
		return config, nil
	}
	existing.NotificationsSenderID = config.NotificationsSenderID
	existing.AnnouncementsSenderID = config.AnnouncementsSenderID
	existing.NewslettersSenderID = config.NewslettersSenderID
	if err := s.db.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

// ListEmailTemplates returns all templates.
func (s *notificationService) ListEmailTemplates() ([]*models.EmailTemplate, error) {
	var templates []*models.EmailTemplate
	if err := s.db.Order("created_at desc").Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

// CreateEmailTemplate inserts a template.
func (s *notificationService) CreateEmailTemplate(tpl *models.EmailTemplate) (*models.EmailTemplate, error) {
	if err := s.db.Create(tpl).Error; err != nil {
		return nil, err
	}
	return tpl, nil
}

// UpdateEmailTemplate updates a template.
func (s *notificationService) UpdateEmailTemplate(id uint, updates *models.EmailTemplate) (*models.EmailTemplate, error) {
	var tpl models.EmailTemplate
	if err := s.db.First(&tpl, id).Error; err != nil {
		return nil, err
	}
	if updates.Name != "" {
		tpl.Name = updates.Name
	}
	if updates.Trigger != "" {
		tpl.Trigger = updates.Trigger
	}
	if updates.Subject != "" {
		tpl.Subject = updates.Subject
	}
	if updates.Body != "" {
		tpl.Body = updates.Body
	}
	if updates.Channel != "" {
		tpl.Channel = updates.Channel
	}
	if updates.Status != "" {
		tpl.Status = updates.Status
	}
	if err := s.db.Save(&tpl).Error; err != nil {
		return nil, err
	}
	return &tpl, nil
}

// DeleteEmailTemplate deletes a template.
func (s *notificationService) DeleteEmailTemplate(id uint) (bool, error) {
	if err := s.db.Delete(&models.EmailTemplate{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// ListAnnouncements returns all announcements.
func (s *notificationService) ListAnnouncements() ([]*models.Announcement, error) {
	var list []*models.Announcement
	if err := s.db.Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// CreateAnnouncement inserts a new announcement.
func (s *notificationService) CreateAnnouncement(announcement *models.Announcement) (*models.Announcement, error) {
	if err := s.db.Create(announcement).Error; err != nil {
		return nil, err
	}
	return announcement, nil
}

// ListNewsletters returns all newsletters.
func (s *notificationService) ListNewsletters() ([]*models.Newsletter, error) {
	var list []*models.Newsletter
	if err := s.db.Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// CreateNewsletter inserts a newsletter draft.
func (s *notificationService) CreateNewsletter(newsletter *models.Newsletter) (*models.Newsletter, error) {
	if err := s.db.Create(newsletter).Error; err != nil {
		return nil, err
	}
	return newsletter, nil
}

// UpdateNewsletter updates a newsletter draft.
func (s *notificationService) UpdateNewsletter(id uint, updates *models.Newsletter) (*models.Newsletter, error) {
	var n models.Newsletter
	if err := s.db.First(&n, id).Error; err != nil {
		return nil, err
	}
	if updates.Subject != "" {
		n.Subject = updates.Subject
	}
	if updates.Preheader != "" {
		n.Preheader = updates.Preheader
	}
	if updates.Content != "" {
		n.Content = updates.Content
	}
	if updates.Segment != "" {
		n.Segment = updates.Segment
	}
	if updates.Status != "" {
		n.Status = updates.Status
	}
	if err := s.db.Save(&n).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

// ScheduleNewsletter sets schedule time and status.
func (s *notificationService) ScheduleNewsletter(id uint, when time.Time) (*models.Newsletter, error) {
	var n models.Newsletter
	if err := s.db.First(&n, id).Error; err != nil {
		return nil, err
	}
	n.ScheduledAt = &when
	n.Status = models.NewsletterScheduled
	if err := s.db.Save(&n).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

// MarkNewsletterSent marks a newsletter as sent.
func (s *notificationService) MarkNewsletterSent(id uint) (*models.Newsletter, error) {
	var n models.Newsletter
	if err := s.db.First(&n, id).Error; err != nil {
		return nil, err
	}
	now := time.Now()
	n.SentAt = &now
	n.Status = models.NewsletterSent
	if err := s.db.Save(&n).Error; err != nil {
		return nil, err
	}
	return &n, nil
}
