package entity

func NewSuccessResult() *BaseResult {
	res := new(BaseResult)
	res.Code = ResultCodeOk
	res.Msg = ResultMsgOk
	return res
}

func NewParamErrorResult(msg string) *BaseResult {
	res := new(BaseResult)
	res.Code = ResultParamError
	res.Msg = msg
	return res
}
