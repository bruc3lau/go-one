package database

import (
	"fmt"
	"go-one/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "root:admin@tcp(127.0.0.1:3306)/go_one?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置详细的日志级别
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 设置日志级别为 Info
				IgnoreRecordNotFoundError: false,       // 不忽略 ErrRecordNotFound 错误
				Colorful:                  true,        // 彩色输出
			},
		),
		// 添加命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	// 打印迁移开始信息
	fmt.Println("开始执行数据库迁移...")

	// 删除已存在的表（仅开发环境使用）
	//err = db.Migrator().DropTable(&models.User{})
	//if err != nil {
	//	fmt.Printf("删除表失败: %v\n", err)
	//}

	// 自动迁移
	err = db.AutoMigrate(
		&models.User{}, // 可以添加更多模型
		// &models.Order{},
		// &models.Product{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %v", err)
	}

	// 打印迁移完成信息
	fmt.Println("数据库迁移完成")

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
