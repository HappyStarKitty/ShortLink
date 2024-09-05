// link model
package model

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID          uint   `gorm:"primaryKey"`
	OriginalURL string `gorm:"not null"`
	ShortCode   string `gorm:"unique;not null"`
	StartTime   time.Time
	EndTime     time.Time
	IsActive    bool `gorm:"default:true"`
	UserID      uint
	Comment     string
	CreatedAt   time.Time // 等价于StartTime
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
