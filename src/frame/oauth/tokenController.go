package oauth

import (
	"frame/entity"
	"frame/interfaces"
)

func ControllerPostToken(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
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
			loginRes := PassworMethodAuthorize(p.UserName, p.Password, p.Client)
			resp.SetObjResult(loginRes)
		}
	default:
		resp.SetObjResult(entity.NewParamErrorResult("[GrantType]参数错误"))
		return
	}

}
