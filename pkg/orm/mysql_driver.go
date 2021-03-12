package orm

import (
	"laravel-go/pkg/orm/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConn(conn config.ConnParam) *gorm.DB {

	dsn := conn.Username + ":" + conn.Password + "@tcp(" + conn.Host + ":" + conn.Port + ")/" + conn.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
