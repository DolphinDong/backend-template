package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
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
	userId, _ := ctx.Get(constant.UserContextKey)
	menu, err := mc.MenuService.GetUserMenu(userId.(string))
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "query menu failed:"))
		response.ResponseHttpError(ctx, "获取用户菜单失败")
		return
	}
	response.ResponseOkWithData(ctx, menu)
}
