package router

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/controller/system"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRouter(engine *gin.Engine) {
	api := engine.Group("/api")

	{
		menuController := system.NewMenuController()
		systemApi := api.Group("/system")
		systemApi.GET("/menu", menuController.GetUserMenu)

	}

	api.GET("/test", func(ctx *gin.Context) {
		response.ResponseOkWithMessage(ctx, "hello backend-template!!!!")
	})

	api.GET("/get-permission", func(ctx *gin.Context) {
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
