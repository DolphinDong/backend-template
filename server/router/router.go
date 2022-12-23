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
		systemApi := api.Group("/system")
		{
			menuController := system.NewMenuController()
			systemApi.GET("/menu", menuController.GetUserMenu)

			userController := system.NewUserController()
			systemApi.GET("/user", userController.GetUserInfo)
		}

	}

	api.GET("/test1", func(ctx *gin.Context) {
		response.ResponseHttpErrorWithInfo(ctx, "错误提示")
	})
	api.GET("/test2", func(ctx *gin.Context) {
		response.ResponseHttpError(ctx, "错误")
	})
	api.GET("/test3", func(ctx *gin.Context) {
		response.ResponseOKCodeWithWarningMessage(ctx, "警告")
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
