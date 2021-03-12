package module

import (
	"laravel-go/app/service/demo"
	"laravel-go/pkg/orm"
)

// 数据字段
type moduleItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// 实例化
func NewModuleQueryService() *ModuleQueryService {
	return &ModuleQueryService{}
}

// 模块查询服务
type ModuleQueryService struct{}

// 数据详情
func (q *ModuleQueryService) Find(id int) (interface{}, error) {

	var result *moduleItem

	if err := orm.DB.Model(&demo.Demo{}).First(&result, id).Error; err != nil {
		panic(err)
	}

	return result, nil
}

// 分页列表
func (q *ModuleQueryService) PageList(title string, page int, size int) (interface{}, error) {

	var items []*moduleItem

	query := orm.DB.Model(&demo.Demo{})

	if title != "" {
		query = query.Where("title like ?", "%"+title+"%")
	}

	results := orm.NewPaginate().Paginate(query, &items, page, size)

	return results, nil
}
