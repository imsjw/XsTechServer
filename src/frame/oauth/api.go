package oauth

import (
	"frame/handler"
	"frame/interfaces"
	"frame/protocol/http"
)

func Init() {
	initConfig()
	if !configEnable {
		return
	}
	initRouter()
	AddFilterWhiteList()
	handler.AddFilter(filter)
}

func AddFilterWhiteList() {
	handler.AddFilterWhiteList(http.MethodPOST, UrlToken)
}

func filter(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) bool {
	return false
}
