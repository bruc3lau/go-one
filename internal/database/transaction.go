package database

import (
	"context"
	"fmt"
	"go-one/internal/models"
	"gorm.io/gorm"
)

// TransactionFunc 定义事务中要执行的函数类型
type TransactionFunc func(tx *gorm.DB) error

// WithTransaction 事务管理器，类似 Spring 的 @Transactional
func WithTransaction(ctx context.Context, fn TransactionFunc) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// 示例：带事务的用户服务方法
type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// TransferPoints 用户积分转账示例，演示事务使用
func (s *UserService) TransferPoints(ctx context.Context, fromUserID, toUserID uint, points int) error {
	return WithTransaction(ctx, func(tx *gorm.DB) error {
		// 1. 检查并扣减支出方积分
		var fromUser models.User
		if err := tx.First(&fromUser, fromUserID).Error; err != nil {
			return err
		}
		if fromUser.Points < points {
			return fmt.Errorf("用户积分不足")
		}
		if err := tx.Model(&fromUser).Update("points", fromUser.Points-points).Error; err != nil {
			return err
		}

		// 2. 增加接收方积分
		var toUser models.User
		if err := tx.First(&toUser, toUserID).Error; err != nil {
			return err
		}
		if err := tx.Model(&toUser).Update("points", toUser.Points+points).Error; err != nil {
			return err
		}

		return nil
	})
}

// BatchCreateUsers 批量创建用户示例
func (s *UserService) BatchCreateUsers(ctx context.Context, users []models.User) error {
	return WithTransaction(ctx, func(tx *gorm.DB) error {
		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return err // 任何错误都会触发回滚
			}
		}
		return nil
	})
}
