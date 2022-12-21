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

func ResponseHttpErrorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, &Response{
		Code: ResponseErrorMsg,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseHttpBadRequestWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, &Response{
		Code: ResponseErrorMsg,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseOkCodeWithMsgAndData(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// http 200  返回错误的msg
func ResponseErrorCodeWithMsg(ctx *gin.Context, msg string) {
	ResponseOkCodeWithMsgAndData(ctx, ResponseErrorMsg, msg, nil)
}

// http 200 返回bad request msg
func ResponseBadRequestCodeWithMsg(ctx *gin.Context, msg string) {
	ResponseOkCodeWithMsgAndData(ctx, ResponseBadRequestMsg, msg, nil)
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
