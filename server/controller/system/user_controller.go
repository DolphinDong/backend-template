package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/common/structs"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/service/system"
	"github.com/DolphinDong/backend-template/tools"
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
	// userId := "user02"
	//userId = "efa07b65-ff48-4409-8ae1-6d8aec0f9475"
	userId, _ := ctx.Get(constant.UserContextKey)
	userInfo, err := uc.UserService.GetUserInfo(userId.(string))
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user info failed"))
		response.ResponseHttpError(ctx, "获取用户信息失败")
		return
	}
	response.ResponseOkWithData(ctx, userInfo)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	query := &structs.TableQuery{}
	err := ctx.ShouldBind(query)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user failed"))
		response.ResponseHttpError(ctx, "获取用户列表失败")
		return
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}
	if query.Page == 0 {
		query.PageSize = 1
	}
	gender := ctx.Query("gender")
	//isAdmin := ctx.Query("is_admin")
	status := ctx.Query("status")

	users, err := uc.UserService.GetUsers(query, gender, status)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user failed"))
		response.ResponseHttpError(ctx, "获取用户列表失败")
		return
	}
	response.ResponseOkWithData(ctx, users)
}

func (uc *UserController) AddUser(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "add user failed"))
		response.ResponseHttpError(ctx, "获取用户信息失败")
		return
	}
	err = tools.Validate(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err = uc.UserService.AddUser(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "add user failed"))
		response.ResponseHttpError(ctx, "添加失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "添加成功")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user failed"))
		response.ResponseHttpError(ctx, "获取用户信息失败")
		return
	}
	err = tools.Validate(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err = uc.UserService.UpdateUser(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user failed"))
		response.ResponseHttpError(ctx, "编辑失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "编辑成功")
}

func (uc *UserController) ResetUserPwd(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "reset user password failed"))
		response.ResponseHttpError(ctx, "获取用户信息失败")
		return
	}
	err = uc.UserService.ResetUserPassword(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "reset user password failed"))
		response.ResponseHttpError(ctx, "重置失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "重置成功")
}
