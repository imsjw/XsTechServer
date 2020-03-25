package entity

import "constant"

type Result struct {
	Code int
	Data interface{}
}

func NewSuccessResult(data interface{}) *Result {
	res := new(Result)
	res.Code = constant.ResultCodeOk
	res.Data = data
	return res
}

func NewParamErrorResult(data string) *Result {
	res := new(Result)
	res.Code = constant.ResultParamError
	res.Data = data
	return res
}
