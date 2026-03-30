package admin

import (
	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// AdminService provides configuration and analytics.
type AdminService interface {
	ListFeatureFlags() ([]*models.FeatureFlag, error)
	UpsertFeatureFlag(flag *models.FeatureFlag) (*models.FeatureFlag, error)
	DeleteFeatureFlag(id uint) (bool, error)
	ListSkillCategories() ([]*models.SkillCategory, error)
	UpsertSkillCategory(cat *models.SkillCategory) (*models.SkillCategory, error)
	DeleteSkillCategory(id uint) (bool, error)
	ListCompliancePolicies() ([]*models.CompliancePolicy, error)
	UpsertCompliancePolicy(policy *models.CompliancePolicy) (*models.CompliancePolicy, error)
	DeleteCompliancePolicy(id uint) (bool, error)
	ListComplianceLogs() ([]*models.ComplianceLog, error)
	CreateComplianceLog(log *models.ComplianceLog) (*models.ComplianceLog, error)
	AdminStats() (*AdminStats, error)
	ListPublicTalents() ([]*models.TalentProfile, error)
	GetPublicTalentByUser(userID uint) (*models.TalentProfile, error)
	ListPublicProjects() ([]*models.Project, error)
	GetPublicProject(id uint) (*models.Project, error)
	ListPublicTalentsFiltered(skill string, skills []string, location string, title string, sort string, limit int, offset int) ([]*models.TalentProfile, error)
	ListPublicProjectsFiltered(category string, status string, ownerID *uint, sort string, limit int, offset int) ([]*models.Project, error)
}

type adminService struct {
	db *gorm.DB
}

type AdminStats struct {
	Users        int64
	Talents      int64
	Clients      int64
	Jobs         int64
	Projects     int64
	Applications int64
}

func NewAdminService(db *gorm.DB) AdminService {
	return &adminService{db}
}

func (s *adminService) ListFeatureFlags() ([]*models.FeatureFlag, error) {
	var flags []*models.FeatureFlag
	if err := s.db.Order("created_at desc").Find(&flags).Error; err != nil {
		return nil, err
	}
	return flags, nil
}

func (s *adminService) UpsertFeatureFlag(flag *models.FeatureFlag) (*models.FeatureFlag, error) {
	var existing models.FeatureFlag
	if err := s.db.Where("key = ?", flag.Key).First(&existing).Error; err != nil {
		if err := s.db.Create(flag).Error; err != nil {
			return nil, err
		}
		return flag, nil
	}
	existing.Enabled = flag.Enabled
	existing.Description = flag.Description
	if err := s.db.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (s *adminService) DeleteFeatureFlag(id uint) (bool, error) {
	if err := s.db.Delete(&models.FeatureFlag{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *adminService) ListSkillCategories() ([]*models.SkillCategory, error) {
	var cats []*models.SkillCategory
	if err := s.db.Order("name asc").Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (s *adminService) UpsertSkillCategory(cat *models.SkillCategory) (*models.SkillCategory, error) {
	var existing models.SkillCategory
	if err := s.db.Where("name = ?", cat.Name).First(&existing).Error; err != nil {
		if err := s.db.Create(cat).Error; err != nil {
			return nil, err
		}
		return cat, nil
	}
	existing.Description = cat.Description
	if err := s.db.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (s *adminService) DeleteSkillCategory(id uint) (bool, error) {
	if err := s.db.Delete(&models.SkillCategory{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *adminService) ListCompliancePolicies() ([]*models.CompliancePolicy, error) {
	var policies []*models.CompliancePolicy
	if err := s.db.Order("created_at desc").Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}

func (s *adminService) UpsertCompliancePolicy(policy *models.CompliancePolicy) (*models.CompliancePolicy, error) {
	if policy.ID != 0 {
		var existing models.CompliancePolicy
		if err := s.db.First(&existing, policy.ID).Error; err == nil {
			if policy.Title != "" {
				existing.Title = policy.Title
			}
			if policy.Content != "" {
				existing.Content = policy.Content
			}
			if policy.Status != "" {
				existing.Status = policy.Status
			}
			if err := s.db.Save(&existing).Error; err != nil {
				return nil, err
			}
			return &existing, nil
		}
	}
	if err := s.db.Create(policy).Error; err != nil {
		return nil, err
	}
	return policy, nil
}

func (s *adminService) DeleteCompliancePolicy(id uint) (bool, error) {
	if err := s.db.Delete(&models.CompliancePolicy{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *adminService) ListComplianceLogs() ([]*models.ComplianceLog, error) {
	var logs []*models.ComplianceLog
	if err := s.db.Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (s *adminService) CreateComplianceLog(log *models.ComplianceLog) (*models.ComplianceLog, error) {
	if err := s.db.Create(log).Error; err != nil {
		return nil, err
	}
	return log, nil
}

func (s *adminService) AdminStats() (*AdminStats, error) {
	stats := &AdminStats{}
	s.db.Model(&models.User{}).Count(&stats.Users)
	s.db.Model(&models.User{}).Where("role = ?", "talent").Count(&stats.Talents)
	s.db.Model(&models.User{}).Where("role = ?", "client").Count(&stats.Clients)
	s.db.Model(&models.JobOffer{}).Count(&stats.Jobs)
	s.db.Model(&models.Project{}).Count(&stats.Projects)
	s.db.Model(&models.JobApplicant{}).Count(&stats.Applications)
	return stats, nil
}

func (s *adminService) ListPublicTalents() ([]*models.TalentProfile, error) {
	var profiles []*models.TalentProfile
	if err := s.db.Preload("Skills").Preload("Experiences").Preload("WorkSamples").
		Where("is_public = ?", true).Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *adminService) ListPublicTalentsFiltered(skill string, skills []string, location string, title string, sort string, limit int, offset int) ([]*models.TalentProfile, error) {
	query := s.db.Preload("Skills").Preload("Experiences").Preload("WorkSamples").Where("is_public = ?", true)
	if location != "" {
		query = query.Where("location ILIKE ?", "%"+location+"%")
	}
	if title != "" {
		query = query.Where("professional_title ILIKE ?", "%"+title+"%")
	}
	if skill != "" {
		query = query.Joins("JOIN skills ON skills.profile_id = talent_profiles.id").Where("skills.skill_name ILIKE ?", "%"+skill+"%")
	}
	if len(skills) > 0 {
		query = query.Joins("JOIN skills AS skill_filter ON skill_filter.profile_id = talent_profiles.id").
			Where("skill_filter.skill_name IN ?", skills)
	}
	switch sort {
	case "trending":
		// weight views higher but keep recency in play
		query = query.Order("(view_count * 0.7 + EXTRACT(EPOCH FROM updated_at) * 0.3) desc")
	case "most_viewed":
		query = query.Order("view_count desc")
	default:
		query = query.Order("updated_at desc")
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	var profiles []*models.TalentProfile
	if err := query.Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *adminService) GetPublicTalentByUser(userID uint) (*models.TalentProfile, error) {
	var profile models.TalentProfile
	if err := s.db.Preload("Skills").Preload("Experiences").Preload("WorkSamples").
		Where("user_id = ? AND is_public = ?", userID, true).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *adminService) ListPublicProjects() ([]*models.Project, error) {
	var projects []*models.Project
	if err := s.db.Preload("Owner").Where("is_public = ?", true).Order("updated_at desc").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *adminService) ListPublicProjectsFiltered(category string, status string, ownerID *uint, sort string, limit int, offset int) ([]*models.Project, error) {
	query := s.db.Preload("Owner").Where("is_public = ?", true)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if ownerID != nil {
		query = query.Where("owner_id = ?", *ownerID)
	}
	if category != "" {
		query = query.Where("summary ILIKE ? OR title ILIKE ?", "%"+category+"%", "%"+category+"%")
	}
	switch sort {
	case "trending":
		query = query.Order("(view_count * 0.7 + EXTRACT(EPOCH FROM updated_at) * 0.3) desc")
	case "most_viewed":
		query = query.Order("view_count desc")
	default:
		query = query.Order("updated_at desc")
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	var projects []*models.Project
	if err := query.Order("updated_at desc").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *adminService) GetPublicProject(id uint) (*models.Project, error) {
	var project models.Project
	if err := s.db.Preload("Owner").Where("id = ? AND is_public = ?", id, true).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}
