package user

import (
	"github.com/gin-gonic/gin"
	"go-wire-architecture/internal/domain"
	"go-wire-architecture/internal/domain/types"
	"go-wire-architecture/pkg"
	"gorm.io/gorm"
)

// Service layer
type Service struct {
	logger          pkg.Logger
	repository      Repository
	paginationScope *gorm.DB
}

// NewService creates a new user service
func NewService(logger pkg.Logger, userRepository Repository) *Service {
	return &Service{
		logger:     logger,
		repository: userRepository,
	}
}

// WithTrx delegates transaction to repository database
func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// SetPaginationScope sets
func (s Service) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) Service {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

// Create creates the user
func (s Service) Create(user *domain.User) error {
	return s.repository.Create(&user).Error
}

// GetOneUser gets one user
func (s Service) GetOneUser(userID types.BinaryUUID) (user domain.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

// GetAllUser get all the user
func (s Service) GetAllUser() (response map[string]interface{}, err error) {
	var users []domain.User
	var count int64

	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": users, "count": count}, nil
}

// UpdateUser updates the user
func (s Service) UpdateUser(user *domain.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s Service) DeleteUser(uuid types.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&domain.User{}).Error
}
