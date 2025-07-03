package database

import (
	"context"
	"fmt"
	"go-one/internal/models"
	"gorm.io/gorm"
	"time"
)

// TransactionFunc 定义事务中要执行的函数类型
type TransactionFunc func(tx *gorm.DB) error

// TransactionPropagation 事务传播行为
type TransactionPropagation int

const (
	// PropagationRequired 如果当前没有事务，就新建一个事务；如果已经存在一个事务，加入到这个事务中
	PropagationRequired TransactionPropagation = iota
	// PropagationRequiresNew 新建事务，如果当前存在事务，把当前事务挂起
	PropagationRequiresNew
	// PropagationNested 如果当前存在事务，则在嵌套事务内执行；如果当前没有事务，则按PropagationRequired执行
	PropagationNested
)

// TransactionOption 事务选项
type TransactionOption struct {
	Propagation TransactionPropagation
	Timeout     time.Duration
}

// WithTransaction 事务管理器，类似 Spring 的 @Transactional
func WithTransaction(ctx context.Context, fn TransactionFunc) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// WithTransactionPropagation 支持事务传播行为的事务管理器
func WithTransactionPropagation(ctx context.Context, opt TransactionOption, fn TransactionFunc) error {
	// 获取当前上下文中的事务
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		switch opt.Propagation {
		case PropagationRequired:
			return fn(tx)
		case PropagationRequiresNew:
			return DB.WithContext(ctx).Transaction(fn)
		case PropagationNested:
			return tx.SavePoint("sp1").Transaction(fn)
		}
	}

	// 如果没有现有事务，创建新事务
	return DB.WithContext(ctx).Transaction(fn)
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
	// 获取当前上下文中的事务
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		// 使用现有事务
		return s.transferPointsWithTx(tx, fromUserID, toUserID, points)
	}

	// 如果上下文中没有事务，创建新事务
	return WithTransaction(ctx, func(tx *gorm.DB) error {
		return s.transferPointsWithTx(tx, fromUserID, toUserID, points)
	})
}

// transferPointsWithTx 使用指定事务执行积分转账
func (s *UserService) transferPointsWithTx(tx *gorm.DB, fromUserID, toUserID uint, points int) error {
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

// 使用示例：带传播行为的服务方法
func (s *UserService) ComplexOperation(ctx context.Context, userID uint) error {
	return WithTransactionPropagation(ctx, TransactionOption{
		Propagation: PropagationRequired,
		Timeout:     time.Second * 10,
	}, func(tx *gorm.DB) error {
		// 在事务中执行操作
		return nil
	})
}
