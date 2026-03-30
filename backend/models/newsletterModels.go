package models

import (
	"time"
	"gorm.io/gorm"
)

// NewsletterStatus tracks lifecycle of a newsletter.
type NewsletterStatus string

const (
	NewsletterDraft     NewsletterStatus = "DRAFT"
	NewsletterScheduled NewsletterStatus = "SCHEDULED"
	NewsletterSent      NewsletterStatus = "SENT"
)

// Newsletter represents a broadcast message.
type Newsletter struct {
	gorm.Model
	Subject     string          `json:"subject" gorm:"not null"`
	Preheader   string          `json:"preheader" gorm:"type:text"`
	Content     string          `json:"content" gorm:"type:text;not null"`
	Segment     string          `json:"segment" gorm:"default:'all'"`
	Status      NewsletterStatus `json:"status" gorm:"type:text;default:'DRAFT'"`
	ScheduledAt *time.Time      `json:"scheduled_at"`
	SentAt      *time.Time      `json:"sent_at"`
}
