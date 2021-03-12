package module

import (
	"laravel-go/app/service/demo/module"
	"laravel-go/pkg/libs"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ModuleController struct{}

func NewModuleController() *ModuleController {
	return &ModuleController{}
}

func (m *ModuleController) Find(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := module.NewModuleQueryService().Find(id)

	if err != nil {
		return libs.NewApi().Fail(c, err.Error())
	}

	return libs.NewApi().Ok(c, data)
}

func (m *ModuleController) PageList(c echo.Context) error {

	title := c.QueryParam("title")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))

	datas, err := module.NewModuleQueryService().PageList(title, page, size)

	if err != nil {
		return libs.NewApi().Fail(c, err.Error())
	}

	return libs.NewApi().Ok(c, datas)
}

func (m *ModuleController) Create(c echo.Context) error {

	title := c.FormValue("title")

	datas, err := module.NewModuleUpdateService().Create(title)

	if err != nil {
		return libs.NewApi().Fail(c, err.Error())
	}

	return libs.NewApi().Ok(c, datas)
}

func (m *ModuleController) Update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")

	datas, err := module.NewModuleUpdateService().Update(id, title)

	if err != nil {
		return libs.NewApi().Fail(c, err.Error())
	}

	return libs.NewApi().Ok(c, datas)
}

func (m *ModuleController) Delete(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	datas, err := module.NewModuleUpdateService().Delete(id)

	if err != nil {
		return libs.NewApi().Fail(c, err.Error())
	}

	return libs.NewApi().Ok(c, datas)
}
