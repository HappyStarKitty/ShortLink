// user model
package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           `gorm:"primaryKey"`
	Name         string         `gorm:"not null"`
	Email        string         `gorm:"unique;not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"` // 自动设置创建时间
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"` // 自动更新修改时间
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Password     string         `gorm:"not null"`
	PasswordHash string         `gorm:"not null"`
}
