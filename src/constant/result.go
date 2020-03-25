package constant

import "entity"

/**
框架使用常量范围[1,100]
*/

const (
	ResultCodeOk     = 0
	ResultParamError = 1
)

const (
	ResultMsgOk         = "ok"
	ResultMsgParamError = "invalid parameter"
)

var ResultOk = entity.BaseResult{Code: ResultCodeOk, Msg: "ok"}
var ResultDataBaseConnectError = entity.BaseResult{Code: 1, Msg: "数据库访问异常"}
