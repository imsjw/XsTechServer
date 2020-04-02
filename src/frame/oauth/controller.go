package oauth

import (
	"frame/entity"
	"frame/interfaces"
)

func controllerPostToken(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
	p := new(struct {
		UserName  string
		Password  string
		Client    string
		GrantType string
	})
	err := req.GetObjParam(p)
	if err != nil {
		resp.SetObjResult(entity.NewParamErrorResult("请传入合法参数"))
		return
	}

	if p.UserName == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[UserName]不能为空"))
		return
	}

	if p.Password == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[Password]不能为空"))
		return
	}

	if p.Client == "" {
		resp.SetObjResult(entity.NewParamErrorResult("[Client]不能为空"))
		return
	}

	switch p.GrantType {
	case "password":
		{
			loginRes := ServicePassworMethodAuthorize(p.UserName, p.Password, p.Client)
			resp.SetObjResult(loginRes)
		}
	default:
		resp.SetObjResult(entity.NewParamErrorResult("[GrantType]参数错误"))
		return
	}
}

func ControllerGetToken(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
	token, _ := req.GetHeader(headerKeyToken)
	oauth := ServiceGetAuthByAccessToken(token)
	if oauth == nil {
		resp.SetObjResult(entity.BaseResult{1000, "Token不存在", nil})
		return
	}

	res := new(struct {
		UserId                  int
		Client                  string
		AccessToken             string
		AccessTokenExpiresTime  int64
		RefreshToken            string
		RefreshTokenExpiresTime int64
	})

	res.UserId = oauth.UserId
	res.Client = oauth.Client
	res.AccessToken = oauth.AccessToken
	res.AccessTokenExpiresTime = oauth.AccessTokenExpiresTime
	res.RefreshToken = oauth.RefreshToken
	res.RefreshTokenExpiresTime = oauth.RefreshTokenExpiresTime
	resp.SetObjResult(entity.NewSuccessResult(res))
}
