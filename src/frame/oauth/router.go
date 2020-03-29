package oauth

import (
	"frame/protocol/http"
	"frame/router"
)

func initRouter() {
	router.Add(UrlToken, ControllerPostToken, http.MethodPOST)
	router.Add(UrlToken, ControllerPostToken, http.MethodGET)
}
