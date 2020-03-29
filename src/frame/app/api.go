package app

import (
	"frame/config"
	"frame/datasource"
	"frame/log"
	"frame/oauth"
)

func Init() {
	log.Info("frame.app.Start 开始启动应用")
	config.Init()
	datasource.Init()
	oauth.Init()
}

func Start() {
	listen()
}
