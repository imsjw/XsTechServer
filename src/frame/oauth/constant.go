package oauth

import "frame/entity"

/**
框架使用常量范围[1,100]
*/
var ResultUserOrPasswordError = entity.BaseResult{Code: 100, Msg: "账号不存在或密码错误"}

//ConfigKey
const (
	configKeyOauthEnable                = "oauth.enable"
	configKeyOauthAccessTokenValidTime  = "oauth.accessToken.validTime"
	configKeyOauthRefreshTokenValidTime = "oauth.refreshToken.validTime"
	configKeyOauthAccessTokenSalt       = "oauth.accessToken.salt"
	configKeyOauthRefreshTokenSalt      = "oauth.refreshToken.salt"
	configKeyOAuthPasswordSalt          = "oauth.password.salt"
)

//header中的key
const (
	headerKeyToken = "Token"
)

//RouterUrl
const (
	urlToken = "/oauth/token"
)
