package system

import "laravel-go/pkg/orm/config"

type SystemLog struct {
	config.Model
	Request string `json:"request"`
	Message string `json:"message"`
}
