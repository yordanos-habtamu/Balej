package moderation

import (
	"fmt"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// ModerationService orchestrates the oversight protocols for the Balej network.
type ModerationService interface {
	FlagTarget(targetType, targetID, reason string, reportedBy uint) error
	ResolveFlag(flagID uint, action string, notes string, adminID uint) error
	CreateResolution(flagID uint, action string, notes string, adminID uint) (*models.ModerationResolution, error)
	ListResolutions(flagID uint) ([]*models.ModerationResolution, error)
	GetModerationQueue() ([]*models.Flag, error)
	LogAction(action string, adminID uint, targetID string, metadata string) error
	GetAuditTrail() ([]*models.AuditLog, error)
	LogActionWithResult(action string, adminID uint, targetID string, metadata string) (*models.AuditLog, error)
}

type moderationService struct {
	db *gorm.DB
}

// NewModerationService initializes the oversight node
func NewModerationService(db *gorm.DB) ModerationService {
	return &moderationService{db}
}

// FlagTarget marks a user, job, or message with a moderation challenge.
func (s *moderationService) FlagTarget(targetType, targetID, reason string, reportedBy uint) error {
	flag := &models.Flag{
		TargetType: targetType,
		TargetID:   targetID,
		Reason:     reason,
		Status:     "OPEN",
		ReportedBy: reportedBy,
	}

	return s.db.Create(flag).Error
}

// ResolveFlag finalizes the challenge by applying an administrative directive.
func (s *moderationService) ResolveFlag(flagID uint, action string, notes string, adminID uint) error {
	var flag models.Flag
	if err := s.db.First(&flag, flagID).Error; err != nil {
		return fmt.Errorf("flag not found: %v", err)
	}

	flag.Status = "RESOLVED"
	if err := s.db.Save(&flag).Error; err != nil {
		return fmt.Errorf("failed to finalize flag: %v", err)
	}

	if _, err := s.CreateResolution(flagID, action, notes, adminID); err != nil {
		return err
	}

	// Log the resolution action for the audit trail
	metadata := fmt.Sprintf("Action: %s | FlagID: %d", action, flagID)
	return s.LogAction("FLAG_RESOLVED", adminID, flag.TargetID, metadata)
}

func (s *moderationService) CreateResolution(flagID uint, action string, notes string, adminID uint) (*models.ModerationResolution, error) {
	resolution := &models.ModerationResolution{
		FlagID:     flagID,
		Action:     action,
		Notes:      notes,
		ResolvedBy: adminID,
		ResolvedAt: time.Now(),
	}
	if err := s.db.Create(resolution).Error; err != nil {
		return nil, err
	}
	return resolution, nil
}

func (s *moderationService) ListResolutions(flagID uint) ([]*models.ModerationResolution, error) {
	var resolutions []*models.ModerationResolution
	if err := s.db.Where("flag_id = ?", flagID).Order("resolved_at desc").Find(&resolutions).Error; err != nil {
		return nil, err
	}
	return resolutions, nil
}

// GetModerationQueue returns all active challenges for the Balej admin HUD.
func (s *moderationService) GetModerationQueue() ([]*models.Flag, error) {
	var flags []*models.Flag
	// Preload the reporter identity for high-fidelity admin displays
	err := s.db.Preload("ReportedByUser").Where("status = ?", "OPEN").Find(&flags).Error
	return flags, err
}

// LogAction persists an administrative pulse to the system-wide audit record.
func (s *moderationService) LogAction(action string, adminID uint, targetID string, metadata string) error {
	log := &models.AuditLog{
		Action:      action,
		PerformedBy: adminID,
		TargetID:    targetID,
		Metadata:    metadata,
		Timestamp:   time.Now(),
	}

	return s.db.Create(log).Error
}

// LogActionWithResult persists and returns the audit log.
func (s *moderationService) LogActionWithResult(action string, adminID uint, targetID string, metadata string) (*models.AuditLog, error) {
	log := &models.AuditLog{
		Action:      action,
		PerformedBy: adminID,
		TargetID:    targetID,
		Metadata:    metadata,
		Timestamp:   time.Now(),
	}
	if err := s.db.Create(log).Error; err != nil {
		return nil, err
	}
	return log, nil
}

// GetAuditTrail returns the historical administrative pulse of the platform.
func (s *moderationService) GetAuditTrail() ([]*models.AuditLog, error) {
	var logs []*models.AuditLog
	err := s.db.Preload("AdminUser").Order("timestamp desc").Find(&logs).Error
	return logs, err
}
