package constant

import "frame/entity"

/**
框架使用常量范围[1,100]
*/
var ResultUserOrPasswordError = entity.BaseResult{Code: 100, Msg: "账号不存在或密码错误"}
