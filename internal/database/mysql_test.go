package database

import (
	"fmt"
	"go-one/internal/models"
	"testing"
)

func TestUserCRUD(t *testing.T) {
	// 初始化数据库连接
	if err := InitDB(); err != nil {
		t.Fatalf("数据库初始化失败: %v", err)
	}

	// 测试创建用户
	t.Run("Create User", func(t *testing.T) {
		user := &models.User{
			Name:     "张三",
			Email:    "zhangsan@example.com",
			Age:      25,
			IsActive: true,
		}

		if err := CreateUser(user); err != nil {
			t.Errorf("创建用户失败: %v", err)
		}

		if user.ID == 0 {
			t.Error("创建用户后ID应该不为0")
		}
	})

	// 测试查询用户
	t.Run("Get User", func(t *testing.T) {
		// 先创建一个用户
		user := &models.User{
			Name:     "李四",
			Email:    "lisi@example.com",
			Age:      30,
			IsActive: true,
		}
		if err := CreateUser(user); err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}

		// 测试查询
		foundUser, err := GetUserByID(user.ID)
		fmt.Printf("foundUser: %+v\n", foundUser)
		if err != nil {
			t.Errorf("查询用户失败: %v", err)
		}
		if foundUser.Name != user.Name {
			t.Errorf("期望用户名为 %s，实际为 %s", user.Name, foundUser.Name)
		}
	})

	// 测试更新用户
	t.Run("Update User", func(t *testing.T) {
		// 先创建一个用户
		user := &models.User{
			Name:     "王五",
			Email:    "wangwu@example.com",
			Age:      35,
			IsActive: true,
		}
		if err := CreateUser(user); err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}

		// 更新年龄
		user.Age = 36
		if err := UpdateUser(user); err != nil {
			t.Errorf("更新用户失败: %v", err)
		}

		// 验证更新
		updatedUser, err := GetUserByID(user.ID)
		if err != nil {
			t.Errorf("查询更新后的用户失败: %v", err)
		}
		if updatedUser.Age != 36 {
			t.Errorf("期望年龄为 36，实际为 %d", updatedUser.Age)
		}
	})

	// 测试用户列表
	t.Run("List Users", func(t *testing.T) {
		users, err := ListUsers(1, 10)
		if err != nil {
			t.Errorf("查询用户列表失败: %v", err)
		}
		if len(users) == 0 {
			t.Log("用户列表为空")
		}
		fmt.Printf("%+v\n", users[0])
	})

	// 测试删除用户
	t.Run("Delete User", func(t *testing.T) {
		// 先创建一个用户
		user := &models.User{
			Name:     "赵六",
			Email:    "zhaoliu@example.com",
			Age:      40,
			IsActive: true,
		}
		if err := CreateUser(user); err != nil {
			t.Fatalf("创建测试用户失败: %v", err)
		}

		// 删除用户
		if err := DeleteUser(user.ID); err != nil {
			t.Errorf("删除用户失败: %v", err)
		}

		// 验证删除
		_, err := GetUserByID(user.ID)
		if err == nil {
			t.Error("用户应该已被删除")
		}
	})
}
