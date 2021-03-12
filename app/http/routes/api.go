package routes

import (
	"laravel-go/app/http/controllers/demo/module"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 设置 API 路由
func SetApiRoute(e *echo.Echo) {

	// 首页
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo Framework")
	})

	// 示例路由
	demo := e.Group("/demo")

	// 示例-模块路由
	moduleDemo := demo.Group("/module")
	moduleDemo.GET("", module.NewModuleController().PageList)
	moduleDemo.GET("/:id", module.NewModuleController().Find)
	moduleDemo.POST("", module.NewModuleController().Create)
	moduleDemo.PUT("/:id", module.NewModuleController().Update)
	moduleDemo.DELETE("/:id", module.NewModuleController().Delete)
}
