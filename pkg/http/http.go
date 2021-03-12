package http

import (
	"encoding/json"
	"fmt"
	"laravel-go/app/http/routes"
	systemLog "laravel-go/app/service/system/log"
	systemPolice "laravel-go/app/service/system/police"
	"laravel-go/pkg/libs"
	"net/http"
	"runtime"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 实例化 HTTP 服务
func New(port string) {

	e := echo.New()

	e.Debug = false

	// 配置 API 路由
	routes.SetApiRoute(e)

	// 请求日志 Debug
	// e.Use(middleware.Logger())

	// Panic 恢复处理
	// e.Use(middleware.Recover())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))

	// 自定义异常处理
	// 捕获到 panic、error 错误接口响应 500
	e.HTTPErrorHandler = HTTPErrorHandler

	e.Start(port)
}

// HTTP 错误处理
func HTTPErrorHandler(err error, c echo.Context) {

	he, ok := err.(*echo.HTTPError)

	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	code := he.Code
	message := he.Message
	if m, ok := he.Message.(string); ok {
		message = echo.Map{"message": m, "error": err.Error()}
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			libs.NewApi().Error(c, err, code, message)
		}

		// 记录日志
		if err != nil {
			createHTTPErrorLog(c, err)
		}
	}
}

// 记录 HTTP 错误日志
func createHTTPErrorLog(c echo.Context, err error) {

	// 请求参数
	req := c.Request()
	formParams, _ := c.FormParams()

	params, _ := json.Marshal(map[string]interface{}{
		"remote_ip":    c.RealIP(),
		"host":         req.Host,
		"uri":          req.RequestURI,
		"method":       req.Method,
		"path":         req.URL.Path,
		"protocol":     req.Proto,
		"user_agent":   req.UserAgent(),
		"query_params": c.QueryParams(),
		"form_params":  formParams,
	})

	// 错误堆栈
	stack := make([]byte, 2<<10) // 2KB 堆栈
	length := runtime.Stack(stack, true)
	message := fmt.Sprintf("[%v] %s\n", err, stack[:length])

	// 写入日志
	insertId, _ := systemLog.NewLogUpdateService().Create(string(params), message)
	newInsertId := strconv.Itoa(int(insertId.(uint)))

	// 发送报警信息
	errStr := err.Error()
	msg := errStr + "-ID:" + string(newInsertId)
	systemPolice.NewPoliceUpdateService().Send(msg)
}
