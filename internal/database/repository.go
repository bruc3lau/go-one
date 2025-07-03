package database

import (
	"go-one/internal/models"
	"gorm.io/gorm"
)

// UserRepository 类似 Spring JPA Repository
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByNameAndAgeGreaterThan 通过方法名表达查询意图
func (r *UserRepository) FindByNameAndAgeGreaterThan(name string, age int) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("name = ? AND age > ?", name, age).Find(&users).Error
	return users, err
}

// FindByIsActiveAndPointsGreaterThanOrderByPointsDesc 复杂查询示例
func (r *UserRepository) FindByIsActiveAndPointsGreaterThanOrderByPointsDesc(isActive bool, points int) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("is_active = ? AND points > ?", isActive, points).
		Order("points DESC").
		Find(&users).Error
	return users, err
}

// ExistsByEmail 检查邮箱是否存在
func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// CountByAgeGreaterThan 统计查询
func (r *UserRepository) CountByAgeGreaterThan(age int) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("age > ?", age).Count(&count).Error
	return count, err
}
