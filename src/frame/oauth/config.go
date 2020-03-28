package oauth

import (
	"fmt"
	"frame/config"
	"frame/log"
)

var configAccessTokenValidTime int64
var configRefreshTokenValidTime int64
var configAccessTokenSalt string
var configRefreshTokenSalt string
var configPasswordSalt string

func initConfig() {
	if config.Exist(ConfigKeyOauthEnable) {
		return
	}
	if config.GetBool(ConfigKeyOauthEnable) {
		return
	}
	if !config.Exist(ConfigKeyOauthAccessTokenValidTime) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", ConfigKeyOauthAccessTokenValidTime, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(ConfigKeyOauthRefreshTokenValidTime) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", ConfigKeyOauthRefreshTokenValidTime, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(ConfigKeyOauthAccessTokenSalt) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", ConfigKeyOauthAccessTokenSalt, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(ConfigKeyOauthRefreshTokenSalt) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", ConfigKeyOauthRefreshTokenSalt, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	if !config.Exist(ConfigKeyOAuthPasswordSalt) {
		errMsg := fmt.Sprint("frame.oauth.Init 缺少配置:[", ConfigKeyOAuthPasswordSalt, "]")
		log.Error(errMsg)
		panic(errMsg)
	}

	configAccessTokenValidTime = config.GetInt64(ConfigKeyOauthAccessTokenValidTime)
	configRefreshTokenValidTime = config.GetInt64(ConfigKeyOauthRefreshTokenValidTime)
	configAccessTokenSalt = config.GetString(ConfigKeyOauthAccessTokenSalt)
	configRefreshTokenSalt = config.GetString(ConfigKeyOauthRefreshTokenSalt)
	configPasswordSalt = config.GetString(ConfigKeyOAuthPasswordSalt)
}
