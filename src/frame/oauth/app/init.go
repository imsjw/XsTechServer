package app

import (
	"frame/config"
	"frame/oauth/router"
)

func init() {
	if config.GetOAuthEnable() {
		router.Init()
	}
}
