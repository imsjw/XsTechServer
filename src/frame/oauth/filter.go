package oauth

import (
	"frame/handler"
	"frame/interfaces"
	"frame/protocol/http"
	"time"
)

var filterWhiteList map[string][]string

func init() {
	filterWhiteList = make(map[string][]string)
}

func filterInit() {
	handler.AddFilterWhiteList(http.MethodPOST, urlToken)
	addFilterWhiteList(http.MethodGET, urlToken)
	addFilterWhiteList(http.MethodPOST, urlTokenRefresh)
}

func addFilterWhiteList(method string, url string) {
	urls, exist := filterWhiteList[method]
	if !exist {
		urls = []string{}
	}
	urls = append(urls, url)
	filterWhiteList[method] = urls
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

	if isWhiteList(req) {
		return true
	}

	return serviceExistResourceByUserId(auth.UserId, url, method)
}

func isWhiteList(req interfaces.Request) bool {
	//判断是否是白名单中的url
	urls, exist := filterWhiteList[req.GetMethod()]
	if exist {
		for _, url := range urls {
			if url == req.GetURL() {
				return true
			}
		}
	}

	return false
}
