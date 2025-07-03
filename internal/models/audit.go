package models

import (
	"gorm.io/gorm"
	"time"
)

// AuditModel 提供审计功能的基础模型
type AuditModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedBy string         `gorm:"size:255"`
	UpdatedBy string         `gorm:"size:255"`
}

// BeforeCreate GORM 钩子，在创建记录前自动设置审计信息
func (m *AuditModel) BeforeCreate(tx *gorm.DB) error {
	if user, ok := tx.Statement.Context.Value("current_user").(string); ok {
		m.CreatedBy = user
		m.UpdatedBy = user
	}
	return nil
}

// BeforeUpdate GORM 钩子，在更新记录前自动设置审计信息
func (m *AuditModel) BeforeUpdate(tx *gorm.DB) error {
	if user, ok := tx.Statement.Context.Value("current_user").(string); ok {
		m.UpdatedBy = user
	}
	return nil
}
