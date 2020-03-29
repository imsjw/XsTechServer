package entity

func NewSuccessResult(data interface{}) *BaseResult {
	res := new(BaseResult)
	res.Code = ResultCodeOk
	res.Msg = ResultMsgOk
	res.Data = data
	return res
}

func NewParamErrorResult(msg string) *BaseResult {
	res := new(BaseResult)
	res.Code = ResultParamError
	res.Msg = msg
	return res
}
