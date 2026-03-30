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
  Delete(id uint) error
  CreateApplication(job *models.JobApplicant) (*models.JobApplicant,error)
  ApproveApplication(job *models.JobApplicant) error
  RejectApplication(job *models.JobApplicant) error
  FindApplicantByID(id uint) (*models.JobApplicant, error)
  FindApplicant(jobOfferID uint, userID uint) (*models.JobApplicant, error)
  FindAllApplicants() ([]*models.JobApplicant, error)
  SearchJobApplicants(filter *models.JobApplicantFilter) ([]*models.JobApplicant, error)
  SearchJobOffers(filter *models.JobFilter) ([]*models.JobOffer, error)
  UpdateApplicationStatus(jobOfferID uint, userID uint, status string) error
  CreateApplicationNote(note *models.JobApplicationNote) (*models.JobApplicationNote, error)
  ListApplicationNotes(jobOfferID uint, userID uint) ([]*models.JobApplicationNote, error)
  CreateApplicationEvent(event *models.JobApplicationEvent) (*models.JobApplicationEvent, error)
  ListApplicationEvents(jobOfferID uint, userID uint) ([]*models.JobApplicationEvent, error)
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

func (s *jobService) Delete(id uint) error {
	return s.repo.Delete(id)
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

func (s *jobService) FindApplicant(jobOfferID uint, userID uint) (*models.JobApplicant, error) {
	return s.repo.FindApplicant(jobOfferID, userID)
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

func (s *jobService) UpdateApplicationStatus(jobOfferID uint, userID uint, status string) error {
	return s.repo.UpdateApplicationStatus(jobOfferID, userID, status)
}

func (s *jobService) CreateApplicationNote(note *models.JobApplicationNote) (*models.JobApplicationNote, error) {
	return s.repo.CreateApplicationNote(note)
}

func (s *jobService) ListApplicationNotes(jobOfferID uint, userID uint) ([]*models.JobApplicationNote, error) {
	return s.repo.ListApplicationNotes(jobOfferID, userID)
}

func (s *jobService) CreateApplicationEvent(event *models.JobApplicationEvent) (*models.JobApplicationEvent, error) {
	return s.repo.CreateApplicationEvent(event)
}

func (s *jobService) ListApplicationEvents(jobOfferID uint, userID uint) ([]*models.JobApplicationEvent, error) {
	return s.repo.ListApplicationEvents(jobOfferID, userID)
}
