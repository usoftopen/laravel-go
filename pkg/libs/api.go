package libs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct{}

// 实例化
func NewApi() *Api {
	return &Api{}
}

// 接口响应成功
// 200 业务返回成功调用
func (a *Api) Ok(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

// 接口响应失败
// 200 业务返回失败调用，比如参数验证失败
func (a *Api) Fail(c echo.Context, msg string) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 10000,
		"msg":  msg,
		"data": "",
	})
}

// 接口响应异常
// 500 服务器未知错误，比如未知的 error 错误，未知的 panic 错误，这种就需要记录日志报警了
func (a *Api) Error(c echo.Context, err error, code int, message interface{}) error {

	if err != nil {
		// 发送报警记录日志
		// e.Logger.Error(err)
	}

	return c.JSON(code, map[string]interface{}{
		"code": 50000,
		"msg":  message.(echo.Map)["message"].(string),
		"data": message.(echo.Map)["error"].(string),
	})
}
