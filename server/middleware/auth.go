package middleware

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

var (
	noLoginUrls     = []string{"/api/system/login","/api/system/logout"}
	noPermCheckUrls = []string{"/api/system/login", "/api/system/logout","/"}
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
		userId, exists := ctx.Get(constant.UserContextKey)
		if !exists {
			global.Logger.Error("user id 不存在 鉴权失败")
			response.ResponseHttpForbiddenWithMsg(ctx, "Permission denied")
			ctx.Abort()
		}
		ok, err := tools.HasPermission(userId.(string), obj, act, true)
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

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		path = strings.TrimSuffix(path, "/")
		if !tools.ElementInSlice(path, noLoginUrls) {
			token := ctx.Request.Header.Get("Access-Token")
			if token == "" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": "授权过期请重新登录",
				})
				ctx.Abort()
				return
			}
			claim, err := tools.ParseToken(token, []byte(tools.SecretKey))
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": "授权过期请重新登录",
				})
				ctx.Abort()
				return
			}
			ctx.Set(constant.UserContextKey, claim.Issuer)
			// 如果快要过期了就重新生成一个token
			if claim.ExpiresAt < time.Now().Unix()+5*60 {
				newToken, err := tools.CreateToken([]byte(tools.SecretKey), claim.Issuer, constant.TokenPeriod)
				if err != nil {
					global.Logger.Errorf("%+v", errors.WithMessage(err, "create token failed"))
				} else {
					ctx.Writer.Header().Set("New-Token", newToken)
				}
			}
		}
		ctx.Next()
	}
}
