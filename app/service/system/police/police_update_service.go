package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"laravel-go/config"
	"net/http"
)

// 实例化
func NewPoliceUpdateService() *PoliceUpdateService {
	return &PoliceUpdateService{}
}

// 报警服务
type PoliceUpdateService struct {
}

// 发送报警信息
func (s *PoliceUpdateService) Send(message string) {

	appCfg := config.NewAppConfig()

	// 非生产环境不发送
	if (appCfg.AppEnv != appCfg.AppEnvProd) || appCfg.AppPoliceUrl == "" {
		return
	}

	value := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
	}

	jsonData, _ := json.Marshal(value)

	resp, err := http.Post(appCfg.AppPoliceUrl, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
}
