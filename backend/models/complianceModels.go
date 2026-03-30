package models

import (
	"time"

	"gorm.io/gorm"
)

// ComplianceRequestType indicates the request category.
type ComplianceRequestType string

const (
	ComplianceExport  ComplianceRequestType = "EXPORT"
	ComplianceDeletion ComplianceRequestType = "DELETION"
)

// ComplianceRequest tracks export/deletion requests.
type ComplianceRequest struct {
	gorm.Model
	UserID    uint                  `json:"user_id" gorm:"not null"`
	Type      ComplianceRequestType `json:"type" gorm:"type:text;not null"`
	Status    string                `json:"status" gorm:"default:'PENDING'"`
	Notes     string                `json:"notes" gorm:"type:text"`
}

// ComplianceJob represents an export/deletion processing job.
type ComplianceJob struct {
	gorm.Model
	RequestID   uint                  `json:"request_id" gorm:"not null"`
	Type        ComplianceRequestType `json:"type" gorm:"type:text;not null"`
	Status      string                `json:"status" gorm:"default:'QUEUED'"`
	Result      string                `json:"result" gorm:"type:text"`
	CompletedAt *time.Time            `json:"completed_at"`
}
