package entity

/**
框架使用常量范围[1,100]
*/

const (
	ResultCodeOk     = 0
	ResultParamError = 1
)

const (
	ResultMsgOk         = "成功"
	ResultMsgParamError = "无效参数"
)

var ResultOk = BaseResult{Code: ResultCodeOk, Msg: ResultMsgParamError}
var ResultDataBaseConnectError = BaseResult{Code: 1, Msg: "数据库访问异常"}
