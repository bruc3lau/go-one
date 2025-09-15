package database

import (
	"database/sql"
	"fmt"
)

// User 是我们的数据模型
type User struct {
	ID    int64
	Name  string
	Email string
}

// UserRepo 定义了用户数据的操作接口
type UserRepo interface {
	Init() error // 用于初始化，例如创建表
	Create(user *User) error
	GetByID(id int64) (*User, error)
	Update(user *User) error
	Delete(id int64) error
}

// --- SQLite 实现 ---

type sqliteUserRepo struct {
	db *sql.DB
}

// NewSQLiteUserRepo 创建一个新的 SQLite 用户仓库
func NewSQLiteUserRepo(db *sql.DB) UserRepo {
	return &sqliteUserRepo{db: db}
}

// Init 创建 users 表
func (r *sqliteUserRepo) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);
	`
	_, err := r.db.Exec(query)
	return err
}

// Create 插入一个新用户
func (r *sqliteUserRepo) Create(user *User) error {
	res, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}
	user.ID = id
	return nil
}

// GetByID 根据 ID 查询用户
func (r *sqliteUserRepo) GetByID(id int64) (*User, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// Update 更新用户信息
func (r *sqliteUserRepo) Update(user *User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete 删除用户
func (r *sqliteUserRepo) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

/*
--- 如何切换到 MySQL? ---

当您准备好使用 MySQL 时，可以按如下步骤操作：

1. 创建一个新的 MySQL 实现文件，例如 `mysql_user_repo.go`。
2. 在新文件中定义一个 `mysqlUserRepo` 结构体。
3. 为 `mysqlUserRepo` 实现 `UserRepo` 接口的所有方法。SQL 语句可能需要微调（例如，占位符 `?` 在某些驱动中可能不同）。

示例:
```go
package database

import "database/sql"

type mysqlUserRepo struct {
    db *sql.DB
}

func NewMySQLUserRepo(db *sql.DB) UserRepo {
    return &mysqlUserRepo{db: db}
}

// ... 实现接口的所有方法 ...
// func (r *mysqlUserRepo) Create(user *User) error { ... }
// ...
```
4. 在您的主程序中，根据配置选择初始化 `NewSQLiteUserRepo` 或 `NewMySQLUserRepo`。
*/
