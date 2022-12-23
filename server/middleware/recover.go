package middleware

import (
	"fmt"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 捕获错误
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.ResponseHttpError(ctx, "Internal Server Error: "+fmt.Sprintf("%v", err))
				if err2, ok := err.(error); ok {
					global.Logger.Errorf("%+v", errors.WithStack(err2))
				} else {
					global.Logger.Errorf("%+v", errors.New(fmt.Sprintf("%+v", err)))
				}
			}
		}()
		ctx.Next()
	}
}
