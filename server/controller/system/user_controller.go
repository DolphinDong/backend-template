package system

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/service/system"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type UserController struct {
	UserService *system.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: system.NewUserService(),
	}
}

func (uc *UserController) GetUserInfo(ctx *gin.Context) {
	userId := "user02"
	userId = "efa07b65-ff48-4409-8ae1-6d8aec0f9475"
	userInfo, err := uc.UserService.GetUserInfo(userId)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user info failed"))
		response.ResponseHttpError(ctx, "获取用户信息失败")
		return
	}
	response.ResponseOkWithData(ctx, userInfo)
}
