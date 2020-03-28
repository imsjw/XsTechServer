package app

import (
	"frame/config"
	"frame/datasource"
	"frame/log"
)

func Start() {
	log.Info("frame.app.Start 开始启动应用")
	config.Init()
	datasource.Init()
	listen()
}
