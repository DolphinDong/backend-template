package router

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRouter(engine *gin.Engine) {
	engine.GET("/test", func(ctx *gin.Context) {
		response.ResponseOkWithMessage(ctx, "hello backend-template!!!!")
	})

	engine.GET("/get-permission", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "liudong")
		permissions := tools.QueryPermissionByUserID(name)
		response.ResponseOkWithData(ctx, permissions)
		return
	})

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code: http.StatusNotFound,
			Msg:  "404 page not found",
		})
	})
}
