package job

import (
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"github.com/yordanos-habtamu/GormWithGraphql/repository"
)

type JobService interface {
	Create(job *models.JobOffer) error
	GetAll() ([]*models.JobOffer, error)
   FindByID(id uint) (*models.JobOffer, error)
   Update(job *models.JobOffer) error
  CreateApplication(job *models.JobApplicant) (*models.JobApplicant,error)
  ApproveApplication(job *models.JobApplicant) error
  RejectApplication(job *models.JobApplicant) error
  FindApplicantByID(id uint) (*models.JobApplicant, error)
  FindAllApplicants() ([]*models.JobApplicant, error)
  SearchJobApplicants(filter *models.JobApplicantFilter) ([]*models.JobApplicant, error)
  SearchJobOffers(filter *models.JobFilter) ([]*models.JobOffer, error)
}

type jobService struct {
	repo repository.JobRepository
}

func NewJobService(r repository.JobRepository) JobService {
	return &jobService{r}
}

func (s *jobService) Create(job *models.JobOffer) error {
	return s.repo.Create(job)
}

func (s *jobService) GetAll() ([]*models.JobOffer, error) {
	return s.repo.FindAll()
}

func (s *jobService) FindByID(id uint) (*models.JobOffer, error) {
	return s.repo.FindByID(id)
}

func (s *jobService) Update(job *models.JobOffer) error {
	return s.repo.Update(job)
}

func (s *jobService) CreateApplication(job *models.JobApplicant) (*models.JobApplicant,error) {
	return s.repo.CreateApplication(job)
}
func (s *jobService) ApproveApplication(job *models.JobApplicant) error {
	return s.repo.ApproveApplication(job)
}
func (s *jobService) RejectApplication(job *models.JobApplicant) error {
	return s.repo.RejectApplication(job)
}
func (s *jobService) FindApplicantByID(id uint) (*models.JobApplicant, error) {
	return s.repo.FindApplicantByID(id)
}
func (s *jobService) FindAllApplicants() ([]*models.JobApplicant, error) {
	return s.repo.FindAllApplicants()
}

func (s *jobService) SearchJobApplicants(filter *models.JobApplicantFilter) ([]*models.JobApplicant, error) {
	return s.repo.SearchJobApplicants(filter)
}
func (s *jobService) SearchJobOffers(filter *models.JobFilter) ([]*models.JobOffer, error) {
	return s.repo.SearchJobOffers(filter)
}