package oauth

import (
	"frame/handler"
)

func Init() {
	initConfig()
	if !configEnable {
		return
	}
	initRouter()
	filterInit()
	handler.AddFilter(filter)
}
