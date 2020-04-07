package oauth

import (
	"frame/protocol/http"
	"frame/router"
)

func initRouter() {
	router.Add(urlToken, controllerUserLogin, http.MethodPOST)
	router.Add(urlToken, controllerGetToken, http.MethodGET)
	router.Add(urlTokenRefresh, controllerRefreshToken, http.MethodPOST)
}
