package main

import (
	"frame/app"
	_ "github.com/go-sql-driver/mysql"
	"xs/router"
)

func main() {
	app.Start()
	router.Init()
}
