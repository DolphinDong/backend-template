package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/model/dao/system"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
)

type MenuService struct {
	MenuDao *system.MenuDao
}

func NewMenuService() *MenuService {
	return &MenuService{
		MenuDao: system.NewMenuDao(),
	}
}
func (ms *MenuService) GetUserMenu(userId string) (menus []*model2.SystemMenu, err error) {
	allMenus, err := ms.MenuDao.QueryAllMenu()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	menus = []*model2.SystemMenu{}
	for _, menu := range allMenus {
		hasPermission, err := tools.HasPermission(userId, menu.Name, constant.MenuAct)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if !hasPermission {
			continue
		}
		menuMeta, err := ms.MenuDao.QueryMenuMetaByMenuID(menu.ID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		menu.Meta = menuMeta
		permissions, err := ms.MenuDao.QueryMenuPermissionsByMenuID(menu.ID)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		p := []string{menu.Name}
		for _, permission := range permissions {
			p = append(p, permission.Identify)
		}
		menu.Permission = tools.RemoveDuplicateElement(p)
		menus = append(menus, menu)
	}
	return
}
