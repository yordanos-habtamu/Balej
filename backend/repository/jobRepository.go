package repository

import (
	"log"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	Create(job *models.JobOffer) error
	FindAll() ([]*models.JobOffer, error)
	FindAllApplicants() ([]*models.JobApplicant, error)
	FindByID(id uint) (*models.JobOffer, error)
	Update(job *models.JobOffer) error
	FindApplicantByID(id uint) (*models.JobApplicant, error)
	CreateApplication(job *models.JobApplicant) (*models.JobApplicant,error)
	ApproveApplication(job *models.JobApplicant) error
	RejectApplication(job *models.JobApplicant) error
	SearchJobOffers(filter *models.JobFilter) ([]*models.JobOffer, error)
	SearchJobApplicants(filter *models.JobApplicantFilter) ([]*models.JobApplicant, error)
}

type jobRepo struct {
	db *gorm.DB
}

func NewjobRepository(db *gorm.DB) JobRepository {
	return &jobRepo{db}
}

func (r *jobRepo) Create(job *models.JobOffer) error {
	return r.db.Create(job).Error
}
func (r *jobRepo) Update(job *models.JobOffer) error {
	return r.db.Save(job).Error
}
func (r *jobRepo) FindAll() ([]*models.JobOffer, error) {
	var jobs []*models.JobOffer
	err := r.db.Find(&jobs).Error
	return jobs, err
}

func (r *jobRepo) FindByID(id uint) (*models.JobOffer, error) {
	var jobOffer models.JobOffer
    if err := r.db.Preload("Applicants").First(&jobOffer, id).Error; err != nil {
        return nil, err
    }
    return &jobOffer, nil
}

func (r *jobRepo) CreateApplication(job *models.JobApplicant) (*models.JobApplicant,error) {
	log.Printf("Creating job application in service: %+v", job)
    if err := r.db.Create(job).Error; err != nil {
        log.Printf("Error creating job application: %v", err)
        return nil, err
    }
    return job, nil
}
func (r *jobRepo) FindApplicantByID(id uint) (*models.JobApplicant, error) {
	var applicant models.JobApplicant
	if err := r.db.Where("user_id = ?", id).First(&applicant).Error; err != nil {
		log.Printf("Error finding job applicant by ID %d: %v", id, err)
		return nil, err
	}
	return &applicant, nil
}
func (r *jobRepo) ApproveApplication(job *models.JobApplicant) error {
	log.Printf("Approving job application in service: %+v", job)
	if err := r.db.Model(&models.JobApplicant{}).Where("job_offer_id = ? AND user_id = ?", job.JobOfferID, job.UserID).Update("status", "accepted").Error; err != nil {
		log.Printf("Error approving job application: %v", err)
		return err
	}
	return nil
}
func (r *jobRepo) RejectApplication(job *models.JobApplicant) error {
	log.Printf("Rejecting job application in service: %+v", job)
	if err := r.db.Model(&models.JobApplicant{}).Where("job_offer_id = ? AND user_id = ?", job.JobOfferID, job.UserID).Update("status", "rejected").Error; err != nil {
		log.Printf("Error rejecting job application: %v", err)
		return err
	}
	return nil
}
func (r *jobRepo) FindAllApplicants() ([]*models.JobApplicant, error) {
	var applicants []*models.JobApplicant
	err := r.db.Find(&applicants).Error
	if err != nil {
		log.Printf("Error finding all applicants: %v", err)
		return nil, err
	}
	return applicants, nil
}

func (r *jobRepo) SearchJobOffers(filter *models.JobFilter) ([]*models.JobOffer, error) {
	var jobs []*models.JobOffer
	query := r.db.Model(&models.JobOffer{})

	if filter.JobOfferID != 0 {
		query = query.Where("id = ?", filter.JobOfferID)
	}
	if filter.UserID != 0 {
		query = query.Where("offered_by = ?", filter.UserID)
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}

	err := query.Find(&jobs).Error
	if err != nil {
		log.Printf("Error searching job offers: %v", err)
		return nil, err
	}
	return jobs, nil
}

func (r *jobRepo) SearchJobApplicants(filter *models.JobApplicantFilter) ([]*models.JobApplicant, error) {
	var applicants []*models.JobApplicant
	query := r.db.Model(&models.JobApplicant{})

	if filter.JobOfferID != 0 {
		query = query.Where("job_offer_id = ?", filter.JobOfferID)
	}
	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	err := query.Find(&applicants).Error
	if err != nil {
		log.Printf("Error searching job applicants: %v", err)
		return nil, err
	}
	return applicants, nil
}