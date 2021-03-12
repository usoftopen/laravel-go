package config

import (
	"time"
)

// 数据库连接参数
type ConnParam struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// 模型基础
type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
