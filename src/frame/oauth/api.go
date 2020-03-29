package oauth

import (
	"frame/handler"
	"frame/interfaces"
	"frame/protocol/http"
	"time"
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
	token, exist := req.GetHeader(HeaderKeyToken)
	if !exist {
		return false
	}
	auth := DaoSelectOauthByAccessToken(token)
	if auth == nil {
		return false
	}

	if auth.AccessTokenExpiresTime < time.Now().Unix() {
		return false
	}

	return false
}
