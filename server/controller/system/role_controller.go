package system

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/common/structs"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/service/system"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type RoleController struct {
	RoleService *system.RoleService
}
type UpdateRolePermission struct {
	ID          int           `json:"id,string" validate:"required"`
	Permissions []interface{} `json:"permissions" validate:"required"`
}

func NewRoleController() *RoleController {
	return &RoleController{
		RoleService: system.NewRoleService(),
	}
}
func (rc *RoleController) GetRoles(ctx *gin.Context) {
	query := &structs.TableQuery{}
	err := ctx.ShouldBind(query)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get role failed"))
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
	roles, err := rc.RoleService.GetRoles(query)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get role failed"))
		response.ResponseHttpError(ctx, "获取角色列表失败")
		return
	}
	response.ResponseOkWithData(ctx, roles)
}

func (rc *RoleController) AddRole(ctx *gin.Context) {
	role := &model.Role{}
	err := ctx.ShouldBind(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "add role failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err = rc.RoleService.AddRole(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "add role failed"))
		response.ResponseHttpError(ctx, "添加失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "添加成功")
}

func (rc *RoleController) UpdateRole(ctx *gin.Context) {
	role := &model.Role{}
	err := ctx.ShouldBind(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update role failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	if role.ID == 0 {
		global.Logger.Errorf("%+v", errors.New("validate parameter failed: id"))
		response.ResponseHttpErrorWithInfo(ctx, "Invalid parameter: id")
		return
	}
	err = rc.RoleService.UpdateRole(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update role failed"))
		response.ResponseHttpError(ctx, "编辑失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "编辑成功")

}

func (rc *RoleController) DeleteRole(ctx *gin.Context) {
	role := &model.Role{}
	err := ctx.ShouldBind(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "reset role password failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	if role.ID == 0 {
		global.Logger.Errorf("%+v", "missing parameter: id")
		response.ResponseHttpErrorWithInfo(ctx, "missing parameter: id")
		return
	}
	err = rc.RoleService.DeleteRole(role)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "delete role failed"))
		response.ResponseHttpError(ctx, "删除失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "删除成功")
}

func (rc *RoleController) GetUserPermissions(ctx *gin.Context) {
	roleStr, ok := ctx.GetQuery("id")
	if !ok {
		global.Logger.Errorf("%+v", errors.New("Missing parameter: id"))
		response.ResponseHttpErrorWithInfo(ctx, "Missing parameter: id")
		return
	}
	roleId, err := strconv.ParseInt(roleStr, 10, 64)
	if err != nil {
		global.Logger.Errorf("%+v", errors.New("Invalid parameter: id"))
		response.ResponseHttpErrorWithInfo(ctx, "Invalid parameter: id")
		return
	}

	permissions, err := rc.RoleService.GetRolePermissions(roleId)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get user permission failed"))
		response.ResponseHttpError(ctx, "获取用户权限失败"+err.Error())
		return
	}
	response.ResponseOkWithData(ctx, permissions)

}

func (rc *RoleController) UpdateRolePermission(ctx *gin.Context) {
	params := &UpdateRolePermission{}
	err := ctx.ShouldBind(params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update role permission failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err=rc.RoleService.UpdateRolePermission(params.ID,params.Permissions)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update role permission failed"))
		response.ResponseHttpError(ctx, "修改失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "修改成功")
}
