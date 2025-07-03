package database

import (
	"context"
	"go-one/internal/models"
	"testing"
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
