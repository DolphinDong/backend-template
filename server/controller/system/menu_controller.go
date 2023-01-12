package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	system2 "github.com/DolphinDong/backend-template/model/dao/system"
	"github.com/DolphinDong/backend-template/service/system"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type MenuController struct {
	MenuService *system.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{
		MenuService: system.NewMenuService(),
	}
}
func (mc *MenuController) GetUserMenu(ctx *gin.Context) {
	//menu, err := mc.MenuService.GetUserMenu("user02")
	userId, _ := ctx.Get(constant.UserContextKey)
	menu, err := mc.MenuService.GetUserMenu(userId.(string))
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "query menu failed:"))
		response.ResponseHttpError(ctx, "获取用户菜单失败")
		return
	}
	response.ResponseOkWithData(ctx, menu)
}

func (mc *MenuController) GetMenus(ctx *gin.Context) {
	menuTree, menuIds, err := mc.MenuService.GetMenus()
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "get menu failed:"))
		response.ResponseHttpError(ctx, "获取菜单失败")
		return
	}
	response.ResponseOkWithData(ctx, gin.H{
		"menu_tree": menuTree,
		"menu_ids":  menuIds,
	})
	return
}

func (mc *MenuController) AddMenu(ctx *gin.Context) {
	menu := &system2.MenuAndPermission{}
	err := ctx.ShouldBind(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "add menu failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}
	err = mc.MenuService.AddMenu(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "add menu failed"))
		response.ResponseHttpError(ctx, "新增失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "新增成功")

}

func (mc *MenuController) UpdateMenu(ctx *gin.Context) {
	menu := &system2.MenuAndPermission{}
	err := ctx.ShouldBind(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update menu failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	err = tools.Validate(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "validate parameter failed"))
		response.ResponseHttpErrorWithInfo(ctx, err.Error())
		return
	}

	if menu.Type == system.MenuType && menu.ID == menu.ParentId {
		global.Logger.Errorf("%+v", errors.New("validate parameter failed, Invalid parameter: parentId"))
		response.ResponseHttpErrorWithInfo(ctx, "Invalid parameter: parentId")
		return
	}

	err = mc.MenuService.UpdateMenu(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "update menu failed"))
		response.ResponseHttpError(ctx, "编辑失败"+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "编辑成功")
}

func (mc *MenuController) DeleteMenu(ctx *gin.Context) {
	menu := &system2.MenuAndPermission{}
	err := ctx.ShouldBind(menu)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "delete menu failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	if menu.ID == 0 || menu.Type == 0 {
		global.Logger.Errorf("%+v", errors.New("validate parameter failed, invalid parameter: id or type"))
		response.ResponseHttpErrorWithInfo(ctx, "Invalid parameter: id or type")
		return
	}
	err = mc.MenuService.DeleteMenu(menu.ID, menu.Type)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "delete user failed"))
		response.ResponseHttpError(ctx, "删除失败: "+err.Error())
		return
	}
	response.ResponseOkWithMessage(ctx, "删除成功")
}
