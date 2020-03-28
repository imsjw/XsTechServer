package oauth

import "frame/entity"

/**
框架使用常量范围[1,100]
*/
var ResultUserOrPasswordError = entity.BaseResult{Code: 100, Msg: "账号不存在或密码错误"}

const (
	ConfigKeyOauthEnable                = "oauth.enable"
	ConfigKeyOauthAccessTokenValidTime  = "oauth.accessToken.validTime"
	ConfigKeyOauthRefreshTokenValidTime = "oauth.refreshToken.validTime"
	ConfigKeyOauthAccessTokenSalt       = "oauth.accessToken.salt"
	ConfigKeyOauthRefreshTokenSalt      = "oauth.refreshToken.salt"
	ConfigKeyOAuthPasswordSalt          = "oauth.password.salt"
)
