package database

import (
	"context"
	"fmt"
	"go-one/internal/models"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestTransactionManagement(t *testing.T) {
	// 初始化数据库连接
	if err := InitDB(); err != nil {
		t.Fatalf("数据库初始化失败: %v", err)
	}

	// 创建用户服务实例
	userService := NewUserService(DB)
	ctx := context.Background()

	// 测试积分转账
	t.Run("TransferPoints", func(t *testing.T) {
		// 准备测试数据：创建两个用户
		user1 := models.User{
			Name:     "用户1",
			Email:    "user1@example.com",
			Age:      25,
			IsActive: true,
			Points:   100,
		}
		user2 := models.User{
			Name:     "用户2",
			Email:    "user2@example.com",
			Age:      30,
			IsActive: true,
			Points:   50,
		}

		// 创建测试用户
		if err := DB.Create(&user1).Error; err != nil {
			t.Fatalf("创建用户1失败: %v", err)
		}
		if err := DB.Create(&user2).Error; err != nil {
			t.Fatalf("创建用户2失败: %v", err)
		}

		// 测试正常转账场景
		t.Run("Successful Transfer", func(t *testing.T) {
			err := userService.TransferPoints(ctx, user1.ID, user2.ID, 50)
			if err != nil {
				t.Errorf("转账失败: %v", err)
			}

			// 验证积分变动
			var updatedUser1, updatedUser2 models.User
			DB.First(&updatedUser1, user1.ID)
			DB.First(&updatedUser2, user2.ID)

			if updatedUser1.Points != 50 {
				t.Errorf("用户1积分错误，期望：50，实际：%d", updatedUser1.Points)
			}
			if updatedUser2.Points != 100 {
				t.Errorf("用户2积分错误，期望：100，实际：%d", updatedUser2.Points)
			}
		})

		// 测试积分不足场景
		t.Run("Insufficient Points", func(t *testing.T) {
			err := userService.TransferPoints(ctx, user1.ID, user2.ID, 1000)
			if err == nil {
				t.Error("期望获得积分不足错误，但是没有")
			}

			// 验证积分未变动
			var updatedUser1, updatedUser2 models.User
			DB.First(&updatedUser1, user1.ID)
			DB.First(&updatedUser2, user2.ID)

			if updatedUser1.Points != 50 {
				t.Errorf("用户1积分不应变动，期望：50，实际：%d", updatedUser1.Points)
			}
			if updatedUser2.Points != 100 {
				t.Errorf("用户2积分不应变动，期望：100，实际：%d", updatedUser2.Points)
			}
		})
	})

	// 测试批量创建用户
	t.Run("BatchCreateUsers", func(t *testing.T) {
		users := []models.User{
			{
				Name:     "批量用户1",
				Email:    "batch1@example.com",
				Age:      20,
				IsActive: true,
				Points:   10,
			},
			{
				Name:     "批量用户2",
				Email:    "batch2@example.com",
				Age:      21,
				IsActive: true,
				Points:   20,
			},
		}

		// 测试正常批量创建
		t.Run("Successful Batch Create", func(t *testing.T) {
			err := userService.BatchCreateUsers(ctx, users)
			if err != nil {
				t.Errorf("批量创建用户失败: %v", err)
			}

			// 验证用户是否创建成功
			var count int64
			DB.Model(&models.User{}).Where("email LIKE 'batch%@example.com'").Count(&count)
			if count != 2 {
				t.Errorf("期望创建2个用户，实际创建了%d个", count)
			}
		})

		// 测试批量创建失败场景（邮箱重复）
		t.Run("Batch Create with Duplicate Email", func(t *testing.T) {
			duplicateUsers := []models.User{
				{
					Name:     "重复用户1",
					Email:    "batch1@example.com", // 使用已存在的邮箱
					Age:      25,
					IsActive: true,
					Points:   30,
				},
				{
					Name:     "重复用户2",
					Email:    "new@example.com",
					Age:      26,
					IsActive: true,
					Points:   40,
				},
			}

			err := userService.BatchCreateUsers(ctx, duplicateUsers)
			if err == nil {
				t.Error("期望获得邮箱重复错误，但是没有")
			}

			// 验证没有新用户被创建
			var count int64
			DB.Model(&models.User{}).Where("name LIKE '重复用户%'").Count(&count)
			if count != 0 {
				t.Errorf("不应该创建任何用户，但是创建了%d个", count)
			}
		})
	})

	// 清理测试数据
	t.Cleanup(func() {
		DB.Exec("DELETE FROM user WHERE email LIKE '%@example.com'")
	})
}

func TestWithTransactionPropagation(t *testing.T) {
	// 初始化数据库连接
	if err := InitDB(); err != nil {
		t.Fatalf("数据库初始化失败: %v", err)
	}

	userService := NewUserService(DB)

	// 用于测试的辅助函数：创建用户
	createTestUser := func(name string, points int) (*models.User, error) {
		user := &models.User{
			Name:     name,
			Email:    name + "@example.com",
			Age:      25,
			IsActive: true,
			Points:   points,
		}
		return user, DB.Create(user).Error
	}

	// 测试 PropagationRequired
	t.Run("PropagationRequired", func(t *testing.T) {
		// 准备测试数据
		user1, err := createTestUser("required_user1", 100)
		if err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}
		user2, err := createTestUser("required_user2", 50)
		if err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}

		// 测试正常情况
		t.Run("Successful Transaction", func(t *testing.T) {
			ctx := context.Background()
			err := WithTransactionPropagation(ctx, TransactionOption{
				Propagation: PropagationRequired,
				Timeout:     time.Second * 5,
			}, func(tx *gorm.DB) error {
				// 转账操作
				return userService.TransferPoints(ctx, user1.ID, user2.ID, 30)
			})

			if err != nil {
				t.Errorf("事务执行失败: %v", err)
			}

			// 验证结果
			var updatedUser1, updatedUser2 models.User
			DB.First(&updatedUser1, user1.ID)
			DB.First(&updatedUser2, user2.ID)

			if updatedUser1.Points != 70 {
				t.Errorf("user1 积分错误，期望：70，实际：%d", updatedUser1.Points)
			}
			if updatedUser2.Points != 80 {
				t.Errorf("user2 积分错误，期望：80，实际：%d", updatedUser2.Points)
			}
		})

		// 测试回滚情况
		t.Run("Transaction Rollback", func(t *testing.T) {
			ctx := context.Background()
			initialPoints1 := user1.Points
			initialPoints2 := user2.Points

			err := WithTransactionPropagation(ctx, TransactionOption{
				Propagation: PropagationRequired,
				Timeout:     time.Second * 5,
			}, func(tx *gorm.DB) error {
				// 将事务对象放入上下文
				ctxWithTx := context.WithValue(ctx, "tx", tx)
				// 使用带有事务的上下文执行转账
				err := userService.TransferPoints(ctxWithTx, user1.ID, user2.ID, 1000)
				if err == nil {
					t.Error("应该返回积分不足错误")
				}
				return err // 返回错误以触发回滚
			})

			if err == nil {
				t.Error("期望获得错误，但是没有")
			}

			// 使用事务查询最新状态
			err = WithTransaction(ctx, func(tx *gorm.DB) error {
				var updatedUser1, updatedUser2 models.User
				if err := tx.First(&updatedUser1, user1.ID).Error; err != nil {
					return err
				}
				if err := tx.First(&updatedUser2, user2.ID).Error; err != nil {
					return err
				}

				if updatedUser1.Points != initialPoints1 {
					t.Errorf("user1 积分不应变动，期望：%d，实际：%d", initialPoints1, updatedUser1.Points)
				}
				if updatedUser2.Points != initialPoints2 {
					t.Errorf("user2 积分不应变动，期望：%d，实际：%d", initialPoints2, updatedUser2.Points)
				}
				return nil
			})

			if err != nil {
				t.Errorf("验证结果时发生错误: %v", err)
			}
		})
	})

	// 测试 PropagationRequiresNew
	t.Run("PropagationRequiresNew", func(t *testing.T) {
		// 准备测试数据
		user3, err := createTestUser("new_user1", 100)
		if err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}
		user4, err := createTestUser("new_user2", 50)
		if err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}

		ctx := context.Background()
		// 外部事务
		err = WithTransaction(ctx, func(tx *gorm.DB) error {
			// 在外部事务中执行一些操作
			if err := tx.Model(user3).Update("age", 30).Error; err != nil {
				return err
			}

			// 开启新事务（RequiresNew）
			err := WithTransactionPropagation(ctx, TransactionOption{
				Propagation: PropagationRequiresNew,
				Timeout:     time.Second * 5,
			}, func(newTx *gorm.DB) error {
				// 在新事务中执行积分转账
				return userService.TransferPoints(ctx, user3.ID, user4.ID, 30)
			})

			if err != nil {
				return err
			}

			// 外部事务回滚，但不应影响内部已提交的新事务
			return fmt.Errorf("强制外部事务回滚")
		})

		// 验证结果：外部事务回滚，但新事务的操作应该保留
		var updatedUser3, updatedUser4 models.User
		DB.First(&updatedUser3, user3.ID)
		DB.First(&updatedUser4, user4.ID)

		if updatedUser3.Age == 30 {
			t.Error("外部事务的操作应该回滚")
		}
		if updatedUser3.Points != 70 {
			t.Errorf("user3 积分错误，期望：70，实际：%d", updatedUser3.Points)
		}
		if updatedUser4.Points != 80 {
			t.Errorf("user4 积分错误，期望：80，实际：%d", updatedUser4.Points)
		}
	})

	// 测试 PropagationNested
	t.Run("PropagationNested", func(t *testing.T) {
		// 准备测试数据
		user5, err := createTestUser("nested_user1", 100)
		if err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}
		user6, err := createTestUser("nested_user2", 50)
		if err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}

		ctx := context.Background()
		// 外部事务
		err = WithTransaction(ctx, func(tx *gorm.DB) error {
			// 在外部事务中修改年龄
			if err := tx.Model(user5).Update("age", 35).Error; err != nil {
				return err
			}

			// 开启嵌套事务
			err := WithTransactionPropagation(ctx, TransactionOption{
				Propagation: PropagationNested,
				Timeout:     time.Second * 5,
			}, func(nestedTx *gorm.DB) error {
				// 在嵌套事务中执行积分转账
				if err := userService.TransferPoints(ctx, user5.ID, user6.ID, 30); err != nil {
					return err
				}
				// 模拟嵌套事务中的错误
				return fmt.Errorf("嵌套事务错误")
			})

			// 嵌套事务失败，但外部事务可以继续
			if err != nil {
				t.Logf("嵌套事务预期失败: %v", err)
			}

			// 外部事务继续执行并提交
			return nil
		})

		if err != nil {
			t.Errorf("外部事务不应失败: %v", err)
		}

		// 验证结果
		var updatedUser5, updatedUser6 models.User
		DB.First(&updatedUser5, user5.ID)
		DB.First(&updatedUser6, user6.ID)

		// 外部事务的修改应该保留
		if updatedUser5.Age != 35 {
			t.Errorf("外部事务的年龄修改应该保留，期望：35，实际：%d", updatedUser5.Age)
		}
		// 嵌套事务的修改应该回滚
		if updatedUser5.Points != 100 {
			t.Errorf("嵌套事务的积分修改应该回滚，期望：100，实际：%d", updatedUser5.Points)
		}
		if updatedUser6.Points != 50 {
			t.Errorf("嵌套事务的积分修改应该回滚，期望：50，实际：%d", updatedUser6.Points)
		}
	})

	// 清理测试数据
	t.Cleanup(func() {
		DB.Where("email LIKE '%@example.com'").Delete(&models.User{})
	})
}
