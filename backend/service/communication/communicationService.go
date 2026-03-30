package communication

import (
	"fmt"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// CommunicationService handles file and video session data.
type CommunicationService interface {
	CreateFile(file *models.FileAsset) (*models.FileAsset, error)
	ListFilesByOwner(ownerID uint) ([]*models.FileAsset, error)
	ListFilesByProject(projectID uint) ([]*models.FileAsset, error)
	ListFilesByChat(chatID string) ([]*models.FileAsset, error)
	UpdateFileTags(fileID uint, tags string) (*models.FileAsset, error)
	DeleteFile(fileID uint) (bool, error)
	GetFile(fileID uint) (*models.FileAsset, error)
	CreateVideoSession(session *models.VideoSession) (*models.VideoSession, error)
	UpdateVideoSessionStatus(id uint, status string) (*models.VideoSession, error)
	AddVideoRecording(id uint, recordingURL string) (*models.VideoSession, error)
	ListVideoSessionsForUser(userID uint) ([]*models.VideoSession, error)
	ListVideoSessionsForProject(projectID uint) ([]*models.VideoSession, error)
	InviteToVideoSession(sessionID uint, inviterID uint, inviteeID *uint, inviteeEmail string, message string) (*models.VideoSessionInvite, error)
	ListVideoInvitesForSession(sessionID uint) ([]*models.VideoSessionInvite, error)
	ListVideoInvitesForUser(userID uint) ([]*models.VideoSessionInvite, error)
	RespondToVideoInvite(inviteID uint, accept bool) (*models.VideoSessionInvite, error)
	ListParticipants(sessionID uint) ([]*models.VideoSessionParticipant, error)
	CountInvites(sessionID uint) (int64, error)
	CountParticipants(sessionID uint) (int64, error)
}

type communicationService struct {
	db *gorm.DB
}

// NewCommunicationService initializes the service.
func NewCommunicationService(db *gorm.DB) CommunicationService {
	return &communicationService{db}
}

func (s *communicationService) CreateFile(file *models.FileAsset) (*models.FileAsset, error) {
	if err := s.db.Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (s *communicationService) ListFilesByOwner(ownerID uint) ([]*models.FileAsset, error) {
	var files []*models.FileAsset
	if err := s.db.Where("owner_id = ?", ownerID).Order("created_at desc").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s *communicationService) ListFilesByProject(projectID uint) ([]*models.FileAsset, error) {
	var files []*models.FileAsset
	if err := s.db.Where("project_id = ?", projectID).Order("created_at desc").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s *communicationService) ListFilesByChat(chatID string) ([]*models.FileAsset, error) {
	var files []*models.FileAsset
	if err := s.db.Where("chat_id = ?", chatID).Order("created_at desc").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s *communicationService) UpdateFileTags(fileID uint, tags string) (*models.FileAsset, error) {
	var file models.FileAsset
	if err := s.db.First(&file, fileID).Error; err != nil {
		return nil, err
	}
	file.Tags = tags
	if err := s.db.Save(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (s *communicationService) DeleteFile(fileID uint) (bool, error) {
	if err := s.db.Delete(&models.FileAsset{}, fileID).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *communicationService) GetFile(fileID uint) (*models.FileAsset, error) {
	var file models.FileAsset
	if err := s.db.First(&file, fileID).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (s *communicationService) CreateVideoSession(session *models.VideoSession) (*models.VideoSession, error) {
	if session.Status == "" {
		session.Status = "SCHEDULED"
	}
	if err := s.db.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (s *communicationService) UpdateVideoSessionStatus(id uint, status string) (*models.VideoSession, error) {
	var session models.VideoSession
	if err := s.db.First(&session, id).Error; err != nil {
		return nil, err
	}
	if status != "" {
		session.Status = status
	}
	now := time.Now()
	if status == "LIVE" && session.StartedAt == nil {
		session.StartedAt = &now
	}
	if status == "ENDED" {
		session.EndedAt = &now
	}
	if err := s.db.Save(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *communicationService) AddVideoRecording(id uint, recordingURL string) (*models.VideoSession, error) {
	var session models.VideoSession
	if err := s.db.First(&session, id).Error; err != nil {
		return nil, err
	}
	if recordingURL == "" {
		return nil, fmt.Errorf("recording url required")
	}
	session.RecordingURL = recordingURL
	if err := s.db.Save(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *communicationService) ListVideoSessionsForUser(userID uint) ([]*models.VideoSession, error) {
	var sessions []*models.VideoSession
	if err := s.db.Where("host_id = ?", userID).Order("scheduled_at desc").Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *communicationService) ListVideoSessionsForProject(projectID uint) ([]*models.VideoSession, error) {
	var sessions []*models.VideoSession
	if err := s.db.Where("project_id = ?", projectID).Order("scheduled_at desc").Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *communicationService) InviteToVideoSession(sessionID uint, inviterID uint, inviteeID *uint, inviteeEmail string, message string) (*models.VideoSessionInvite, error) {
	invite := &models.VideoSessionInvite{
		SessionID:    sessionID,
		InviterID:    inviterID,
		InviteeID:    inviteeID,
		InviteeEmail: inviteeEmail,
		Status:       "PENDING",
		Message:      message,
	}
	if err := s.db.Create(invite).Error; err != nil {
		return nil, err
	}
	return invite, nil
}

func (s *communicationService) ListVideoInvitesForSession(sessionID uint) ([]*models.VideoSessionInvite, error) {
	var invites []*models.VideoSessionInvite
	if err := s.db.Where("session_id = ?", sessionID).Order("created_at desc").Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

func (s *communicationService) ListVideoInvitesForUser(userID uint) ([]*models.VideoSessionInvite, error) {
	var invites []*models.VideoSessionInvite
	if err := s.db.Where("invitee_id = ?", userID).Order("created_at desc").Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

func (s *communicationService) RespondToVideoInvite(inviteID uint, accept bool) (*models.VideoSessionInvite, error) {
	var invite models.VideoSessionInvite
	if err := s.db.First(&invite, inviteID).Error; err != nil {
		return nil, err
	}
	status := "DECLINED"
	if accept {
		status = "ACCEPTED"
	}
	invite.Status = status
	now := time.Now()
	invite.RespondedAt = &now
	if err := s.db.Save(&invite).Error; err != nil {
		return nil, err
	}
	if accept && invite.InviteeID != nil {
		participant := &models.VideoSessionParticipant{
			SessionID: invite.SessionID,
			UserID:    *invite.InviteeID,
			Role:      "ATTENDEE",
			Status:    "ACCEPTED",
			JoinedAt:  &now,
		}
		_ = s.db.Create(participant).Error
	}
	return &invite, nil
}

func (s *communicationService) ListParticipants(sessionID uint) ([]*models.VideoSessionParticipant, error) {
	var participants []*models.VideoSessionParticipant
	if err := s.db.Where("session_id = ?", sessionID).Find(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}

func (s *communicationService) CountInvites(sessionID uint) (int64, error) {
	var count int64
	if err := s.db.Model(&models.VideoSessionInvite{}).Where("session_id = ?", sessionID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (s *communicationService) CountParticipants(sessionID uint) (int64, error) {
	var count int64
	if err := s.db.Model(&models.VideoSessionParticipant{}).Where("session_id = ?", sessionID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
