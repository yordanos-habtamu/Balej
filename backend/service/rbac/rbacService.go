package rbac

import (
	"fmt"

	"github.com/yordanos-habtamu/GormWithGraphql/models"
	"gorm.io/gorm"
)

// RBACService manages roles and permissions.
type RBACService interface {
	ListRoles() ([]*models.Role, error)
	CreateRole(role *models.Role) (*models.Role, error)
	UpdateRole(id uint, updates *models.Role) (*models.Role, error)
	DeleteRole(id uint) (bool, error)
	ListPermissions() ([]*models.Permission, error)
	CreatePermission(p *models.Permission) (*models.Permission, error)
	SetRolePermissions(roleID uint, permissionIDs []uint) (*models.Role, error)
}

type rbacService struct {
	db *gorm.DB
}

func NewRBACService(db *gorm.DB) RBACService {
	return &rbacService{db}
}

func (s *rbacService) ListRoles() ([]*models.Role, error) {
	var roles []*models.Role
	if err := s.db.Preload("Permissions").Order("created_at desc").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *rbacService) CreateRole(role *models.Role) (*models.Role, error) {
	if err := s.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (s *rbacService) UpdateRole(id uint, updates *models.Role) (*models.Role, error) {
	var role models.Role
	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	if updates.Name != "" {
		role.Name = updates.Name
	}
	if updates.Summary != "" {
		role.Summary = updates.Summary
	}
	role.IsSystem = updates.IsSystem
	if err := s.db.Save(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *rbacService) DeleteRole(id uint) (bool, error) {
	if err := s.db.Delete(&models.Role{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *rbacService) ListPermissions() ([]*models.Permission, error) {
	var perms []*models.Permission
	if err := s.db.Order("created_at desc").Find(&perms).Error; err != nil {
		return nil, err
	}
	return perms, nil
}

func (s *rbacService) CreatePermission(p *models.Permission) (*models.Permission, error) {
	if err := s.db.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (s *rbacService) SetRolePermissions(roleID uint, permissionIDs []uint) (*models.Role, error) {
	var role models.Role
	if err := s.db.First(&role, roleID).Error; err != nil {
		return nil, err
	}
	var perms []models.Permission
	if err := s.db.Where("id IN ?", permissionIDs).Find(&perms).Error; err != nil {
		return nil, fmt.Errorf("failed to load permissions: %v", err)
	}
	if err := s.db.Model(&role).Association("Permissions").Replace(&perms); err != nil {
		return nil, err
	}
	if err := s.db.Preload("Permissions").First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
