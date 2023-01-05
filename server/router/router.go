package router

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/controller/system"
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
			systemApi.GET("/userInfo", userController.GetUserInfo)
			systemApi.GET("/user", userController.GetUsers)
			systemApi.POST("/user", userController.AddUser)
			systemApi.PUT("/user", userController.UpdateUser)
			systemApi.PUT("/user/resetPwd", userController.ResetUserPwd)

			loginController := system.NewLoginController()
			systemApi.POST("/login", loginController.Login)
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

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code: http.StatusNotFound,
			Msg:  "404 page not found",
		})
	})
}
