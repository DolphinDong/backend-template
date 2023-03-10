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
	"strings"
)

type UpdateUserPermission struct {
	ID          string        `json:"id" validate:"required"`
	Permissions []interface{} `json:"permissions" validate:"required"`
}
type UpdateUserRole struct {
	ID      string   `json:"id" validate:"required"`
	RoleIds []string `json:"roleIds" validate:"required"`
}
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
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}
	if query.Page == 0 {
		query.PageSize = 1
	}
	query.Search = strings.TrimSpace(query.Search)
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
		response.ResponseHttpError(ctx, err.Error())
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
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	if user.ID == "" {
		global.Logger.Errorf("%+v", errors.New("validate parameter failed: id"))
		response.ResponseHttpErrorWithInfo(ctx, "Invalid parameter: id")
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
		response.ResponseHttpError(ctx, err.Error())
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

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(user)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "reset user password failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	if user.ID == "" {
		global.Logger.Errorf("%+v", "missing parameter: id")
		response.ResponseHttpErrorWithInfo(ctx, "missing parameter: id")
		return
	}
	currentUserId, _ := ctx.Get(constant.UserContextKey)
	if user.ID == currentUserId {
		global.Logger.Errorf("%+v", "无法删除当前登录用户")
		response.ResponseHttpErrorWithInfo(ctx, "无法删除当前登录用户")
		return
	}
	err = uc.UserService.DeleteUser(user.ID)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "delete user failed"))
		response.ResponseHttpError(ctx, "删除失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "删除成功")
}

func (uc *UserController) GetUserPermissions(ctx *gin.Context) {
	userId, ok := ctx.GetQuery("id")
	if !ok {
		global.Logger.Errorf("%+v", errors.New("Missing parameter: id"))
		response.ResponseHttpErrorWithInfo(ctx, "Missing parameter: id")
		return
	}
	permissions, err := uc.UserService.GetReqPermission(userId)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user permission failed"))
		response.ResponseHttpError(ctx, "获取用户权限失败"+err.Error())
		return
	}
	response.ResponseOkWithData(ctx, permissions)
}

func (uc *UserController) UpdateUserPermission(ctx *gin.Context) {
	params := &UpdateUserPermission{}
	err := ctx.ShouldBind(params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user permission failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err = uc.UserService.UpdateReqPermission(params.ID, params.Permissions)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user permission failed"))
		response.ResponseHttpError(ctx, "修改失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "修改成功")
}

func (uc *UserController) GetUserRoles(ctx *gin.Context) {
	userId, ok := ctx.GetQuery("user_id")
	if !ok {
		global.Logger.Errorf("%+v", errors.New("Missing parameter: user_id"))
		response.ResponseHttpErrorWithInfo(ctx, "Missing parameter: user_id")
		return
	}
	userRoles, err := uc.UserService.GetUserRoles(userId)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user role failed"))
		response.ResponseHttpError(ctx, "获取用户角色失败 "+err.Error())
		return
	}
	response.ResponseOkWithData(ctx, userRoles)
}

func (uc *UserController) UpdateUserRole(ctx *gin.Context) {
	params := &UpdateUserRole{}
	err := ctx.ShouldBind(params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user role failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err = uc.UserService.UpdateUserRole(params.ID, params.RoleIds)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user role failed"))
		response.ResponseHttpError(ctx, "修改失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "修改成功")
}

func (uc *UserController) UploadUserAvatar(ctx *gin.Context) {
	id, exists := ctx.Get(constant.UserContextKey)
	userId := id.(string)
	if !exists || id == "" {
		global.Logger.Errorf("%+v", errors.New("validate user"))
		response.ResponseHttpError(ctx, "validate user")
		return
	}
	url, err := uc.UserService.UploadUserAvatar(ctx, userId)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update user avatar failed"))
		response.ResponseHttpError(ctx, "修改失败"+err.Error())
		return
	}
	response.ResponseOkWithData(ctx, gin.H{
		"url": url,
	})
}
