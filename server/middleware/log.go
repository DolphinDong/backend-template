package middleware

import (
	"bytes"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// 打印请求信息
func PrintRequestInfo() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ip := ctx.RemoteIP()
		url := ctx.Request.URL
		// 请求方法为post需要将请求体内的东西打印出来
		if ctx.Request.Method == http.MethodPost {
			all, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				response.ResponseHttpErrorWithMsg(ctx, "read request body error")
				global.Logger.Error("read request body error")
				ctx.Abort()
				return
			}
			defer ctx.Request.Body.Close()
			data := string(all)
			// 将内容重新赋值给 ctx.Request.Body
			closer := ioutil.NopCloser(bytes.NewBuffer(all))
			ctx.Request.Body = closer
			global.Logger.Infof("[%v] IP:%v URL:%v RequestBody= %v", ctx.Request.Method, ip, url, data)
		} else {
			global.Logger.Infof("[%v] IP:%v URL:%v", ctx.Request.Method, ip, url)
		}
		// 放行
		ctx.Next()

	}
}
