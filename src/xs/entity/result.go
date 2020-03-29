package entity

const (
	ResultCodeOk         int = 0
	ResultCodeParamError int = 1
)

type Result struct {
	Code int
	Data interface{}
}

func NewSuccessResult(data interface{}) *Result {
	res := new(Result)
	res.Code = ResultCodeOk
	res.Data = data
	return res
}

func NewParamErrorResult(data string) *Result {
	res := new(Result)
	res.Code = ResultCodeParamError
	res.Data = data
	return res
}
