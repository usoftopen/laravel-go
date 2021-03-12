package demo

import "laravel-go/pkg/orm/config"

type Demo struct {
	config.Model
	Title string `json:"title"`
}
