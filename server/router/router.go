package router

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/controller/system"
	"github.com/gin-gonic/gin"
)

func registerRouter(api *gin.RouterGroup) {

	{
		systemApi := api.Group("/system")
		{
			menuController := system.NewMenuController()
			systemApi.GET("/menus", menuController.GetUserMenu)
			systemApi.GET("/menu", menuController.GetMenus)
			systemApi.POST("/menu", menuController.AddMenu)
			systemApi.PUT("/menu", menuController.UpdateMenu)
			systemApi.DELETE("/menu", menuController.DeleteMenu)

			userController := system.NewUserController()
			systemApi.GET("/userInfo", userController.GetUserInfo)
			systemApi.GET("/user", userController.GetUsers)
			systemApi.POST("/user", userController.AddUser)
			systemApi.PUT("/user", userController.UpdateUser)
			systemApi.DELETE("/user", userController.DeleteUser)
			systemApi.PUT("/user/resetPwd", userController.ResetUserPwd)
			systemApi.GET("/user/permission", userController.GetUserPermissions)
			systemApi.PUT("/user/permission", userController.UpdateUserPermission)
			systemApi.GET("/user/role", userController.GetUserRoles)
			systemApi.PUT("/user/role", userController.UpdateUserRole)
			systemApi.POST("/user/avatar", userController.UploadUserAvatar)

			roleController := system.NewRoleController()
			systemApi.GET("/role", roleController.GetRoles)
			systemApi.POST("/role", roleController.AddRole)
			systemApi.PUT("/role", roleController.UpdateRole)
			systemApi.DELETE("/role", roleController.DeleteRole)
			systemApi.GET("/role/permission", roleController.GetUserPermissions)
			systemApi.PUT("/role/permission", roleController.UpdateRolePermission)

			loginController := system.NewLoginController()
			systemApi.POST("/login", loginController.Login)
			systemApi.POST("/logout", loginController.Logout)
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

}
