package controller

import (
	"frame/interfaces"
	"xs/entity"
)

func Login(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) {
	p := struct {
		Account   string
		Password  string
		Client    string
		GrantType string
	}{}
	err := req.GetObjParam(p)
	if err != nil {
		resp.SetObjResult(entity.NewParamErrorResult("请传入合法参数"))
	}

	if p.Account == "" {
		resp.SetObjResult(entity.NewParamErrorResult("用户名不能为空"))
	}

	if p.Password == "" {
		resp.SetObjResult(entity.NewParamErrorResult("密码不能为空"))
	}

	if p.Client == "" {
		resp.SetObjResult(entity.NewParamErrorResult("客户端类型不能为空"))
	}

}
