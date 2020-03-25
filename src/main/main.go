package main

import (
	"frame/application"
	_ "github.com/go-sql-driver/mysql"
	"xs/router"
)

func main() {
	application.Start()
	router.Init()
}
