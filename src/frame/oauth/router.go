package oauth

import (
	"frame/router"
)

func initRouter() {
	router.Add("/oauth/token", ControllerPostToken, "POST")
}
