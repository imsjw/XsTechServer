package oauth

import (
	"frame/router"
)

func initRouter() {
	router.Add(UrlToken, ControllerPostToken, "POST")
}
