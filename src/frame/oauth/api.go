package oauth

import (
	"frame/handler"
	"frame/interfaces"
)

func Init() {
	initConfig()
	if !configEnable {
		return
	}
	initRouter()
	handler.AddFilter(filter)
}

func filter(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) bool {
	return false
}
