package middleware

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

var (
	noLoginUrls     = []string{"/system/login"}
	noPermCheckUrls = []string{"/system/login", "/"}
)

// 权限校验
func PermissionCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		obj := ctx.Request.URL.Path
		obj = strings.TrimSuffix(obj, "/")
		act := ctx.Request.Method
		// 不需要校验权限的url
		if tools.ElementInSlice(obj, noPermCheckUrls) {
			ctx.Next()
			return
		}
		ok, err := tools.HasPermission("efa07b65-ff48-4409-8ae1-6d8aec0f9475", obj, act, true)
		if err != nil {
			global.Logger.Errorf("%+v", errors.WithMessage(err, "鉴权异常"))
			response.ResponseHttpError(ctx, "鉴权异常")
			ctx.Abort()
		}
		if ok {
			ctx.Next()
		} else {
			response.ResponseHttpForbiddenWithMsg(ctx, "Permission denied")
			ctx.Abort()
		}
	}
}
