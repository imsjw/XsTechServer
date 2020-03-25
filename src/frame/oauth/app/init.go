package app

import (
	"frame/config"
	"oauth/router"
)

func init() {
	if config.GetOAuthEnable() {
		router.Init()
	}
}
