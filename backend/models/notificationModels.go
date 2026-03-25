package models

import (
	"time"
	"gorm.io/gorm"
)

// NotificationType categorizes the delivery protocol
type NotificationType string

const (
	TypeTransactional NotificationType = "TRANSACTIONAL" // e.g., Verification, Auth
	TypeNews          NotificationType = "BROADCAST"     // e.g., System updates
	TypeNewsletter    NotificationType = "NEWSLETTER"    // e.g., Marketing, Sector news
)

// NotificationPriority defines the urgency level for the HUD
type NotificationPriority string

const (
	PriorityLow      NotificationPriority = "LOW"
	PriorityNormal   NotificationPriority = "NORMAL"
	PriorityCritical NotificationPriority = "CRITICAL"
)

// Notification represents a persistent alert node for users
type Notification struct {
	gorm.Model
	UserID    uint                 `json:"user_id" gorm:"not null"`
	Title     string               `json:"title" gorm:"not null"`
	Content   string               `json:"content" gorm:"type:text;not null"`
	Type      NotificationType     `json:"type" gorm:"type:notification_type;default:'TRANSACTIONAL'"`
	Priority  NotificationPriority `json:"priority" gorm:"type:notification_priority;default:'NORMAL'"`
	IsRead    bool                 `json:"is_read" gorm:"default:false"`
	ActionURL string               `json:"action_url"` // Deep link for the frontend
	ExpiresAt *time.Time           `json:"expires_at"`
}

// NewsletterSubscription tracks users opting into sector broadcasts
type NewsletterSubscription struct {
	gorm.Model
	Email      string `json:"email" gorm:"unique;not null"`
	Sectors    string `json:"sectors" gorm:"type:text"` // Comma-separated or JSON list
	IsActive   bool   `json:"is_active" gorm:"default:true"`
}
