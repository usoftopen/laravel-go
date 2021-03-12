package orm

import (
	"laravel-go/pkg/orm/config"

	"gorm.io/gorm"
)

// 默认数据库连接
var DB *gorm.DB

// 其他数据库连接
// var OtherDB *gorm.DB

// 创建数据库连接
func CreateDBConn(connParam config.ConnParam) *gorm.DB {

	// TODO::这里可以根据数据库的不同驱动扩展对应的链接
	// 使用：if connParam.Driver == "..."

	return NewMysqlConn(connParam)
}
