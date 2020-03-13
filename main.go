package main

import (
	"go-web/conf"
	"go-web/http"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := http.NewRouter()
	r.Run(":3000")
}
