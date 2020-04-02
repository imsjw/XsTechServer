package oauth

import (
	"frame/protocol/http"
	"frame/router"
)

func initRouter() {
	router.Add(urlToken, controllerPostToken, http.MethodPOST)
	router.Add(urlToken, ControllerGetToken, http.MethodGET)
}
