package main

import (
	"frame/app"
	"frame/oauth"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.Init()
	oauth.Init()
	app.Start()
}
