package compliance

import (
	"fmt"
	"time"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// ComplianceService orchestrates export/deletion request flows.
type ComplianceService interface {
	CreateRequest(userID uint, reqType models.ComplianceRequestType, notes string) (*models.ComplianceRequest, error)
	ListRequests(reqType *models.ComplianceRequestType) ([]*models.ComplianceRequest, error)
	UpdateRequestStatus(id uint, status string, notes string) (*models.ComplianceRequest, error)
	EnqueueJob(requestID uint, reqType models.ComplianceRequestType) (*models.ComplianceJob, error)
	ListJobs(status *string) ([]*models.ComplianceJob, error)
	UpdateJobStatus(id uint, status string, result string) (*models.ComplianceJob, error)
	StartAutoCompleteWorker(stop <-chan struct{}, interval time.Duration)
}

type complianceService struct {
	db *gorm.DB
}

// NewComplianceService initializes the compliance node.
func NewComplianceService(db *gorm.DB) ComplianceService {
	return &complianceService{db}
}

// CreateRequest persists a compliance request.
func (s *complianceService) CreateRequest(userID uint, reqType models.ComplianceRequestType, notes string) (*models.ComplianceRequest, error) {
	req := &models.ComplianceRequest{
		UserID: userID,
		Type:   reqType,
		Status: "PENDING",
		Notes:  notes,
	}
	if err := s.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

// ListRequests returns compliance requests.
func (s *complianceService) ListRequests(reqType *models.ComplianceRequestType) ([]*models.ComplianceRequest, error) {
	var reqs []*models.ComplianceRequest
	query := s.db.Order("created_at desc")
	if reqType != nil {
		query = query.Where("type = ?", string(*reqType))
	}
	if err := query.Find(&reqs).Error; err != nil {
		return nil, err
	}
	return reqs, nil
}

// UpdateRequestStatus updates a request status.
func (s *complianceService) UpdateRequestStatus(id uint, status string, notes string) (*models.ComplianceRequest, error) {
	var req models.ComplianceRequest
	if err := s.db.First(&req, id).Error; err != nil {
		return nil, fmt.Errorf("request not found: %v", err)
	}
	req.Status = status
	if notes != "" {
		req.Notes = notes
	}
	if err := s.db.Save(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

// EnqueueJob creates a processing job for a compliance request.
func (s *complianceService) EnqueueJob(requestID uint, reqType models.ComplianceRequestType) (*models.ComplianceJob, error) {
	job := &models.ComplianceJob{
		RequestID: requestID,
		Type:      reqType,
		Status:    "QUEUED",
	}
	if err := s.db.Create(job).Error; err != nil {
		return nil, err
	}
	return job, nil
}

// ListJobs returns jobs by status.
func (s *complianceService) ListJobs(status *string) ([]*models.ComplianceJob, error) {
	var jobs []*models.ComplianceJob
	query := s.db.Order("created_at desc")
	if status != nil && *status != "" {
		query = query.Where("status = ?", *status)
	}
	if err := query.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

// UpdateJobStatus updates a compliance job status and result.
func (s *complianceService) UpdateJobStatus(id uint, status string, result string) (*models.ComplianceJob, error) {
	var job models.ComplianceJob
	if err := s.db.First(&job, id).Error; err != nil {
		return nil, err
	}
	if status != "" {
		job.Status = status
	}
	if result != "" {
		job.Result = result
	}
	if status == "COMPLETED" {
		now := time.Now()
		job.CompletedAt = &now
	}
	if err := s.db.Save(&job).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

// StartAutoCompleteWorker auto-completes queued jobs periodically (mock processing).
func (s *complianceService) StartAutoCompleteWorker(stop <-chan struct{}, interval time.Duration) {
	if interval <= 0 {
		interval = 30 * time.Second
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			var jobs []models.ComplianceJob
			if err := s.db.Where("status IN ?", []string{"QUEUED", "PROCESSING"}).Limit(10).Find(&jobs).Error; err != nil {
				continue
			}
			for _, job := range jobs {
				if job.Type == models.ComplianceDeletion {
					continue
				}
				_, _ = s.UpdateJobStatus(job.ID, "COMPLETED", "AUTO_COMPLETED")
			}
		case <-stop:
			return
		}
	}
}
