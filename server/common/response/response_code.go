package response

const (
	ResponseOkData    = 20000 // 正常数据
	ResponseOkMsg     = 20001 // 正常msg
	ResponseOk        = 20002 // 正常数据和msg
	ResponseOkWarning = 20003 // 返回警告
	ResponseOkInfo    = 20004 // 返回普通提示信息
)

const (
	ResponseError    = 50000
	ResponseErrorMsg = 50001 // 返回异常信息
)

const (
	ResponseBadRequestMsg = 40001 //
	ResponseForbiddenMsg  = 40003
)
