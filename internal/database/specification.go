package database

import (
	"go-one/internal/models"
	"gorm.io/gorm"
)

// Specification 定义查询规范接口
type Specification interface {
	Apply(db *gorm.DB) *gorm.DB
}

// UserSpecification 用户查询规范
type UserSpecification struct {
	Name     *string
	AgeStart *int
	AgeEnd   *int
	IsActive *bool
}

// Apply 实现查询规范
func (s UserSpecification) Apply(db *gorm.DB) *gorm.DB {
	if s.Name != nil {
		db = db.Where("name LIKE ?", "%"+*s.Name+"%")
	}
	if s.AgeStart != nil {
		db = db.Where("age >= ?", *s.AgeStart)
	}
	if s.AgeEnd != nil {
		db = db.Where("age <= ?", *s.AgeEnd)
	}
	if s.IsActive != nil {
		db = db.Where("is_active = ?", *s.IsActive)
	}
	return db
}

// FindAll 使用规范查询
func (s *UserService) FindAll(spec Specification) ([]models.User, error) {
	var users []models.User
	err := spec.Apply(s.db).Find(&users).Error
	return users, err
}

// Page 分页结果
type Page struct {
	Content    interface{} `json:"content"`
	TotalPages int64       `json:"totalPages"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
	Total      int64       `json:"total"`
}

// FindByPage 分页查询
func (s *UserService) FindByPage(spec Specification, page, size int) (*Page, error) {
	var users []models.User
	var total int64

	db := spec.Apply(s.db)

	// 获取总数
	if err := db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	err := db.Offset((page - 1) * size).Limit(size).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &Page{
		Content:    users,
		TotalPages: (total + int64(size) - 1) / int64(size),
		PageNumber: page,
		PageSize:   size,
		Total:      total,
	}, nil
}
