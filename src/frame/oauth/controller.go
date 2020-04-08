package oauth

import (
	"frame/entity"
	"frame/interfaces"
	"time"
)

/**
登陆接口
*/
func controllerUserLogin(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
	p := new(struct {
		UserName  string `json:"userName"`
		Password  string `json:"password"`
		Client    string `json:"client"`
		GrantType string `json:"grantType"`
	})
	err := req.GetObjParam(p)
	if err != nil {
		resp.SetObjResult(entity.NewParamErrorResult("请传入合法参数"))
		return
	}

	if p.UserName == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[userName]不能为空"))
		return
	}

	if p.Password == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[password]不能为空"))
		return
	}

	if p.Client == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[client]不能为空"))
		return
	}

	switch p.GrantType {
	case "password":
		{
			loginRes := servicePassworMethodAuthorize(p.UserName, p.Password, p.Client)
			resp.SetObjResult(loginRes)
		}
	default:
		resp.SetObjResult(entity.NewParamErrorResult("[grantType]参数错误"))
		return
	}
}

/**
查询token信息接口
*/
func controllerGetToken(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
	token, _ := req.GetHeader(headerKeyToken)
	oauth := serviceGetAuthByAccessToken(token)
	if oauth == nil {
		resp.SetObjResult(entity.BaseResult{1000, "Token不存在", nil})
		return
	}

	res := new(struct {
		UserId                  int    `json:"userId"`
		Client                  string `json:"client"`
		AccessToken             string `json:"accessToken"`
		AccessTokenExpiresTime  int64  `json:"accessTokenExpiresTime"`
		RefreshToken            string `json:"refreshToken"`
		RefreshTokenExpiresTime int64  `json:"refreshTokenExpiresTime"`
	})

	res.UserId = oauth.UserId
	res.Client = oauth.Client
	res.AccessToken = oauth.AccessToken
	res.AccessTokenExpiresTime = oauth.AccessTokenExpiresTime
	res.RefreshToken = oauth.RefreshToken
	res.RefreshTokenExpiresTime = oauth.RefreshTokenExpiresTime
	resp.SetObjResult(entity.NewSuccessResult(res))
}

/**
刷新token的接口
*/
func controllerRefreshToken(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
	p := new(struct {
		RefreshToken string `json:"refreshToken"`
	})
	err := req.GetObjParam(p)
	if err != nil {
		resp.SetObjResult(entity.NewParamErrorResult("请传入合法参数"))
		return
	}
	if p.RefreshToken == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[refreshToken]不能为空"))
		return
	}

	token, _ := req.GetHeader(headerKeyToken)
	auth := serviceGetAuthByAccessToken(token)
	if auth == nil {
		resp.SetObjResult(entity.BaseResult{1000, "Token不存在", nil})
		return
	}
	if auth.RefreshToken != p.RefreshToken {
		resp.SetObjResult(entity.NewParamErrorResult("refreshToken不正确"))
		return
	}
	if auth.RefreshTokenExpiresTime < time.Now().Unix() {
		resp.SetObjResult(entity.NewParamErrorResult("refreshToken已过期"))
		return
	}
	auth = serviceRefreshTokenById(auth)
	res := new(struct {
		UserId                  int    `json:"userId"`
		Client                  string `json:"client"`
		AccessToken             string `json:"accessToken"`
		AccessTokenExpiresTime  int64  `json:"accessTokenExpiresTime"`
		RefreshToken            string `json:"refreshToken"`
		RefreshTokenExpiresTime int64  `json:"refreshTokenExpiresTime"`
	})
	res.UserId = auth.UserId
	res.Client = auth.Client
	res.AccessToken = auth.AccessToken
	res.AccessTokenExpiresTime = auth.AccessTokenExpiresTime
	res.RefreshToken = auth.RefreshToken
	res.RefreshTokenExpiresTime = auth.RefreshTokenExpiresTime

	resp.SetObjResult(entity.NewSuccessResult(res))
}
