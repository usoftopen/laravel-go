package module

import (
	"laravel-go/app/service/demo"
	"laravel-go/pkg/orm"
)

// 实例化
func NewModuleUpdateService() *ModuleUpdateService {
	return &ModuleUpdateService{}
}

// 模块更新服务
type ModuleUpdateService struct {
}

// 创建
func (s *ModuleUpdateService) Create(title string) (interface{}, error) {

	result := demo.Demo{Title: title}
	res := orm.DB.Create(&result)

	if res.Error != nil {
		panic(res.Error)
	}

	return result.ID, nil
}

// 修改
func (s *ModuleUpdateService) Update(id int, title string) (interface{}, error) {

	newData := map[string]interface{}{"title": title}

	if err := orm.DB.Model(&demo.Demo{}).Where("id = ?", id).Updates(newData).Error; err != nil {
		panic(err)
	}

	return id, nil
}

// 删除
func (s *ModuleUpdateService) Delete(id int) (interface{}, error) {

	if err := orm.DB.Delete(&demo.Demo{}, id).Error; err != nil {
		panic(err)
	}

	return "操作成功", nil
}
