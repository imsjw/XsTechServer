package app

import (
	"frame/config"
	"frame/datasource"
	"frame/log"
)

func Init() {
	log.Info("frame.app.Start 开始启动应用")
	config.Init()
	datasource.Init()
}

func Start() {
	listen()
}
