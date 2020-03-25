package router

import (
	"frame/oauth/controller"
	"frame/router"
)

func Init() {
	router.Add("/oauth/token", controller.Token, "POST")
}
