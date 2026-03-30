package models

import (
	"time"
	"gorm.io/gorm"
)

// Flag represents a moderation report against a job, user, or message
type Flag struct {
	gorm.Model
	TargetType string    `json:"target_type" gorm:"not null"` // JOB, USER, MESSAGE
	TargetID   string    `json:"target_id" gorm:"not null"`
	Reason     string    `json:"reason" gorm:"not null;type:text"`
	Status     string    `json:"status" gorm:"default:'OPEN'"` // OPEN, RESOLVED, DISMISSED
	ReportedBy uint      `json:"reported_by" gorm:"not null"`
	ReportedByUser User  `json:"reported_by_user" gorm:"foreignKey:ReportedBy"`
}

// ModerationResolution captures resolution details for a flag.
type ModerationResolution struct {
	gorm.Model
	FlagID     uint      `json:"flag_id" gorm:"not null"`
	Action     string    `json:"action" gorm:"not null"`
	Notes      string    `json:"notes" gorm:"type:text"`
	ResolvedBy uint      `json:"resolved_by" gorm:"not null"`
	ResolvedAt time.Time `json:"resolved_at" gorm:"autoCreateTime"`
}

// AuditLog captures system-wide administrative actions
type AuditLog struct {
	gorm.Model
	Action      string    `json:"action" gorm:"not null"` // e.g. "USER_BANNED", "JOB_DELETED"
	PerformedBy uint      `json:"performed_by" gorm:"not null"`
	AdminUser   User      `json:"admin_user" gorm:"foreignKey:PerformedBy"`
	TargetID    string    `json:"target_id"`
	Metadata    string    `json:"metadata" gorm:"type:json"` // Contextual data
	Timestamp   time.Time `json:"timestamp" gorm:"autoCreateTime"`
}
