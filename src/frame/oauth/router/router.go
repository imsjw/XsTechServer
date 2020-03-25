package router

import (
	"frame/router"
	"oauth/controller"
)

func Init() {
	router.Add("/oauth/token", controller.Token, "POST")
}
