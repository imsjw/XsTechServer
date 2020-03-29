package router

import (
	"frame/router"
	"xs/controller"
)

func Init() {
	router.Add("/user/login", controller.Login, "POST")
}
