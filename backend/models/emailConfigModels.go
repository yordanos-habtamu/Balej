package models

import "gorm.io/gorm"

// EmailSenderProfile stores credentials for an outbound email provider.
type EmailSenderProfile struct {
	gorm.Model
	Label       string `json:"label" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Password    string `json:"password" gorm:"type:text"` // token or password (store encrypted in production)
	APIKey      string `json:"api_key" gorm:"type:text"`
	Provider    string `json:"provider" gorm:"not null"` // SMTP, SendGrid, Mailgun, AWS_SES
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

// EmailRoutingConfig maps use-cases to sender profiles.
type EmailRoutingConfig struct {
	gorm.Model
	NotificationsSenderID *uint `json:"notifications_sender_id"`
	AnnouncementsSenderID *uint `json:"announcements_sender_id"`
	NewslettersSenderID   *uint `json:"newsletters_sender_id"`
}

// EmailTemplate defines reusable templates for transactional, broadcast, or newsletter emails.
type EmailTemplate struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Trigger string `json:"trigger" gorm:"not null"`
	Subject string `json:"subject" gorm:"type:text;not null"`
	Body    string `json:"body" gorm:"type:text;not null"`
	Channel string `json:"channel" gorm:"default:'EMAIL'"`
	Status  string `json:"status" gorm:"default:'DRAFT'"`
}

// Announcement represents a broadcast message in the admin hub.
type Announcement struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Content  string `json:"content" gorm:"type:text;not null"`
	Category string `json:"category" gorm:"not null"`
	Priority string `json:"priority" gorm:"default:'standard'"`
	Author   string `json:"author" gorm:"default:'ADMIN_CONSOLE'"`
}
