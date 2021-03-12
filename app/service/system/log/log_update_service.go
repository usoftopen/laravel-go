package system

import (
	"laravel-go/app/service/system"
	"laravel-go/pkg/orm"
)

// 实例化
func NewLogUpdateService() *LogUpdateService {
	return &LogUpdateService{}
}

// 日志更新服务
type LogUpdateService struct {
}

// 创建
func (s *LogUpdateService) Create(request string, message string) (interface{}, error) {

	result := system.SystemLog{Request: request, Message: message}
	res := orm.DB.Create(&result)

	return result.ID, res.Error
}
