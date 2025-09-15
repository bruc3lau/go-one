package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"go-one/internal/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

// Config 结构体用于映射 config.yaml 的内容
type Config struct {
	Database struct {
		Choice string `yaml:"choice"`
		SQLite struct {
			Path string `yaml:"path"`
		} `yaml:"sqlite"`
		MySQL struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
		} `yaml:"mysql"`
	} `yaml:"database"`
}

// loadConfig 从根目录的 config.yaml 文件加载配置
func loadConfig() (*Config, error) {
	cfg := &Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}

func main() {
	// 从文件加载配置
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var db *sql.DB
	var userRepo database.UserRepo

	switch cfg.Database.Choice {
	case "sqlite":
		// --- SQLite 配置 ---
		db, err = sql.Open("sqlite3", cfg.Database.SQLite.Path)
		if err != nil {
			log.Fatalf("Failed to open sqlite database: %v", err)
		}
		userRepo = database.NewSQLiteUserRepo(db)
		log.Println("Using SQLite database from config.")

	case "mysql":
		// --- MySQL 配置 ---
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.Database.MySQL.User,
			cfg.Database.MySQL.Password,
			cfg.Database.MySQL.Host,
			cfg.Database.MySQL.Port,
			cfg.Database.MySQL.DBName,
		)
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to open mysql database: %v", err)
		}
		userRepo = database.NewMySQLUserRepo(db)
		log.Println("Using MySQL database from config.")

	default:
		log.Fatalf("Invalid dbChoice in config: %s", cfg.Database.Choice)
	}

	defer db.Close()

	// 测试连接
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Database connection successful.")

	// ... (CRUD 操作部分保持不变) ...

	// 初始化数据库
	if err := userRepo.Init(); err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	log.Println("Repository initialized successfully.")

	// --- CRUD 操作演示 ---
	log.Println("--- Creating new user ---")
	newUser := &database.User{Name: "Alice", Email: "alice.from.config@example.com"}
	if err := userRepo.Create(newUser); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("User created with ID: %d", newUser.ID)

	log.Println("--- Reading user ---")
	retrievedUser, err := userRepo.GetByID(newUser.ID)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	log.Printf("Retrieved user: %+v", retrievedUser)

	log.Println("--- Updating user ---")
	retrievedUser.Name = "Alice Smith"
	if err := userRepo.Update(retrievedUser); err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	log.Println("User updated.")

	log.Println("--- Deleting user ---")
	if err := userRepo.Delete(retrievedUser.ID); err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	log.Println("User deleted.")
}
