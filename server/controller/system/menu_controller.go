package system

import (
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/service/system"
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
	menu, err := mc.MenuService.GetUserMenu("efa07b65-ff48-4409-8ae1-6d8aec0f9475")
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "query menu failed:"))
		response.ResponseHttpError(ctx, "获取用户菜单失败")
		return
	}
	response.ResponseOkWithData(ctx, menu)
}
