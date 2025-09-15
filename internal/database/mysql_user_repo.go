package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

// --- MySQL 实现 ---

type mysqlUserRepo struct {
	db *sql.DB
}

// NewMySQLUserRepo 创建一个新的 MySQL 用户仓库
func NewMySQLUserRepo(db *sql.DB) UserRepo {
	return &mysqlUserRepo{db: db}
}

// Init 创建 users 表 (MySQL 版本)
func (r *mysqlUserRepo) Init() error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id BIGINT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE
    ) ENGINE=InnoDB;
    `
	_, err := r.db.Exec(query)
	return err
}

// Create 插入一个新用户 (与 SQLite 实现相同)
func (r *mysqlUserRepo) Create(user *User) error {
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

// GetByID 根据 ID 查询用户 (与 SQLite 实现相同)
func (r *mysqlUserRepo) GetByID(id int64) (*User, error) {
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

// Update 更新用户信息 (与 SQLite 实现相同)
func (r *mysqlUserRepo) Update(user *User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete 删除用户 (与 SQLite 实现相同)
func (r *mysqlUserRepo) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
