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
	addFilterWhiteList()
	handler.AddFilter(filter)
}

func addFilterWhiteList() {
	handler.AddFilterWhiteList(http.MethodPOST, urlToken)
}

func filter(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) bool {
	token, exist := req.GetHeader(headerKeyToken)
	if !exist {
		return false
	}
	auth := daoSelectOauthByAccessToken(token)
	if auth == nil {
		return false
	}

	if auth.AccessTokenExpiresTime < time.Now().Unix() {
		return false
	}

	url := req.GetURL()
	method := req.GetMethod()

	if url == urlToken && method == http.MethodGET {
		return true
	}

	return serviceExistResourceByUserId(auth.UserId, url, method)
}
