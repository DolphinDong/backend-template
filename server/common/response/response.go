package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseOkCodeWithMsgAndData(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResponseWithHttpCodeMsgAndData(ctx *gin.Context, httpCode, code int, msg string, data interface{}) {
	ctx.JSON(httpCode, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResponseHttpError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, &Response{
		Code: ResponseError,
		Msg:  msg,
		Data: nil,
	})
}
func ResponseHttpErrorWithInfo(ctx *gin.Context, info string) {
	ctx.JSON(http.StatusInternalServerError, &Response{
		Code: ResponseErrorMsg,
		Msg:  info,
		Data: nil,
	})
}

func ResponseHttpForbiddenWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusForbidden, &Response{
		Code: ResponseForbiddenMsg,
		Msg:  msg,
		Data: nil,
	})
}

// http 400 返回bad request msg
func ResponseBadRequestWithMsg(ctx *gin.Context, msg string) {
	ResponseWithHttpCodeMsgAndData(ctx, http.StatusBadRequest, ResponseBadRequestMsg, msg, nil)
}

// http 200  返回警告的msg
func ResponseOKCodeWithWarningMessage(ctx *gin.Context, msg string) {
	ResponseOkCodeWithMsgAndData(ctx, ResponseOkWarning, msg, nil)
}

// 返回msg 和 data
func ResponseOkWithMsgAndData(ctx *gin.Context, msg string, data interface{}) {
	ResponseOkCodeWithMsgAndData(ctx, ResponseOk, msg, data)
}

// 只返回message
func ResponseOkWithMessage(ctx *gin.Context, msg string) {
	ResponseOkCodeWithMsgAndData(ctx, ResponseOkMsg, msg, nil)
}

// 只返回data
func ResponseOkWithData(ctx *gin.Context, data interface{}) {
	ResponseOkCodeWithMsgAndData(ctx, ResponseOkData, "", data)
}
