package orm

import (
	"gorm.io/gorm"
)

const (
	DEFAULT_PAGE = 1
	DEFAULT_SIZE = 20
	MAX_SIZE     = 100
)

type PageResult struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

// 实例化分页
func NewPaginate() *Paginate {
	return &Paginate{}
}

type Paginate struct {
}

// 分页处理
func (p *Paginate) Paginate(query *gorm.DB, results interface{}, page int, size int) PageResult {

	var total int64
	query.Count(&total)

	if page <= 0 {
		page = DEFAULT_PAGE
	}

	if size <= 0 {
		size = DEFAULT_SIZE
	} else if size > MAX_SIZE {
		size = MAX_SIZE
	}

	if err := query.Offset((page - 1) * size).Limit(size).Find(results).Error; err != nil {
		panic(err.Error())
	}

	return PageResult{
		Page:  page,
		Size:  size,
		Total: total,
		Items: results,
	}
}
