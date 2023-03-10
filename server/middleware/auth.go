package middleware

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/dao/redis"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

var (
	noLoginUrls     = []string{"/api/system/login", "/api/system/logout"}
	noPermCheckUrls = []string{"/api/system/login", "/api/system/logout", "/"}
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
			token := ctx.Request.Header.Get(constant.TokenHeader)
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
			redisDao := redis.NewRedisDao()
			// 查询redis中是否有这个用户的token
			value, err := redisDao.GetKey(tools.GetRedisTokenKey(constant.TokenRedisPrefix, claim.Issuer, token))
			if err != nil {
				global.Logger.Errorf("%+v", errors.WithMessage(err, "get token failed"))
				response.ResponseHttpError(ctx, "get token failed")
				ctx.Abort()
				return
			}
			if value == nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": "授权过期请重新登录",
				})
				ctx.Abort()
				return
			}

			// 设置token在redis中的时间
			err = redisDao.SetKeyExpiration(tools.GetRedisTokenKey(constant.TokenRedisPrefix, claim.Issuer, token), constant.TokenPeriod*60)
			if err != nil {
				global.Logger.Errorf("%+v", errors.WithMessage(err, "set token expiration failed"))
				response.ResponseHttpError(ctx, "set token expiration failed")
				ctx.Abort()
				return
			}
			//// 如果快要过期了就重新生成一个token
			//if claim.ExpiresAt < time.Now().Unix()+5*60 {
			//	newToken, err := tools.CreateToken([]byte(tools.SecretKey), claim.Issuer, constant.TokenPeriod)
			//	if err != nil {
			//		global.Logger.Errorf("%+v", errors.WithMessage(err, "create token failed"))
			//	} else {
			//		ctx.Writer.Header().Set("New-Token", newToken)
			//	}
			//}
		}
		ctx.Next()
	}
}
