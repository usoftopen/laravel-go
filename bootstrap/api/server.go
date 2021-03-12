package main

import (
	cfg "laravel-go/config"
	"laravel-go/pkg/config"
	"laravel-go/pkg/http"
	"laravel-go/pkg/orm"
)

// API 引导入口
func main() {
	initConfig()
	initDB()
	initHTTP()
}

// 加载配置文件
func initConfig() {
	config.LoadConfig(".env", "env")
}

// 配置数据库
func initDB() {
	dbConf := cfg.NewDBConfig()

	// 设置默认数据库连接
	defaultDB := dbConf.Connections[dbConf.Default]
	orm.DB = orm.CreateDBConn(defaultDB)

	// 设置其他数据库连接
	// otherDB := dbConf.Connections["mysql-other"]
	// orm.OtherDB = orm.CreateDBConn(otherDB)
}

// 启动 HTTP 服务
func initHTTP() {
	http.New(":1323")
}
