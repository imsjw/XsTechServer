package oauth

import (
	"fmt"
	"frame/config"
	"frame/log"
)

var configEnable bool
var configAccessTokenValidTime int64
var configRefreshTokenValidTime int64
var configAccessTokenSalt string
var configRefreshTokenSalt string
var configPasswordSalt string

func initConfig() {
	if !config.Exist(configKeyOauthEnable) {
		return
	}
	configEnable = config.GetBool(configKeyOauthEnable)
	if !configEnable {
		return
	}
	if !config.Exist(configKeyOauthAccessTokenValidTime) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", configKeyOauthAccessTokenValidTime, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(configKeyOauthRefreshTokenValidTime) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", configKeyOauthRefreshTokenValidTime, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(configKeyOauthAccessTokenSalt) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", configKeyOauthAccessTokenSalt, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(configKeyOauthRefreshTokenSalt) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", configKeyOauthRefreshTokenSalt, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(configKeyOAuthPasswordSalt) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", configKeyOAuthPasswordSalt, "]")
		log.Error(errMsg)
		panic(errMsg)
	}

	configAccessTokenValidTime = config.GetInt64(configKeyOauthAccessTokenValidTime)
	configRefreshTokenValidTime = config.GetInt64(configKeyOauthRefreshTokenValidTime)
	configAccessTokenSalt = config.GetString(configKeyOauthAccessTokenSalt)
	configRefreshTokenSalt = config.GetString(configKeyOauthRefreshTokenSalt)
	configPasswordSalt = config.GetString(configKeyOAuthPasswordSalt)
}
