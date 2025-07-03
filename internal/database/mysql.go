package database

import (
	"fmt"
	"go-one/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// 方式1：使用 mysql_native_password
	dsn := "root:admin@tcp(127.0.0.1:3306)/go_one?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"

	// 方式2：如果数据库是全新安装，使用默认的 caching_sha2_password
	// dsn := "root:password@tcp(127.0.0.1:3306)/go_one?charset=utf8mb4&parseTime=True&loc=Local"

	// 方式3：明确指定认证插件
	// dsn := "root:password@tcp(127.0.0.1:3306)/go_one?charset=utf8mb4&parseTime=True&loc=Local&auth=caching_sha2_password"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %v", err)
	}

	DB = db
	return nil
}

// CreateUser 创建用户
func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

// GetUserByID 根据ID获取用户
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *models.User) error {
	return DB.Save(user).Error
}

// DeleteUser 删除用户
func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}

// ListUsers 获取用户列表
func ListUsers(page, pageSize int) ([]models.User, error) {
	var users []models.User
	offset := (page - 1) * pageSize
	err := DB.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, err
}
