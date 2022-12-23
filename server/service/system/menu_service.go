package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/model/dao/system"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"sort"
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
	allMenuAndPermissions, err := ms.MenuDao.QueryAllMenuAndPermission()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	menuAndPermissionMap := ms.groupMenuById(allMenuAndPermissions)
	for _, menuAndPermissions := range menuAndPermissionMap {
		// 获取第0个
		m1 := menuAndPermissions[0]
		hasPermission, err := tools.HasPermission(userId, m1.Name, constant.MenuAct)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if !hasPermission {
			continue
		}
		systemMenu := &model2.SystemMenu{
			ID:        m1.ID,
			ParentId:  m1.ParentId,
			Name:      m1.Name,
			Path:      m1.Path,
			Component: m1.Component,
			Redirect:  m1.Redirect,
			Sort:      m1.Sort,
			Meta: &model2.MenuMeta{
				Title:               m1.Title,
				Icon:                m1.Icon,
				Target:              m1.Target,
				Show:                m1.Show,
				HideChildren:        m1.HideChildren,
				HiddenHeaderContent: m1.HiddenHeaderContent,
			},
		}
		permissions := make([]string, 0)
		permissions = append(permissions, m1.Name)
		for _, menu := range menuAndPermissions {
			if menu.Action != "" && menu.Identify != "" {
				permissions = append(permissions, menu.Identify)
			}
		}
		// 权限去重
		systemMenu.Permission = tools.RemoveDuplicateElement(permissions)
		menus = append(menus, systemMenu)
	}
	// 菜单排序 先按照sort字段从大到小排序，如果sort字段相同则按照id从小到大排序
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort > menus[j].Sort || (menus[i].Sort == menus[j].Sort && menus[i].ID < menus[j].ID)
	})

	return
}
func (ms *MenuService) groupMenuById(menuAndPermissions []*system.MenuAndPermission) (menuAndPermissionMap map[int][]*system.MenuAndPermission) {
	menuAndPermissionMap = make(map[int][]*system.MenuAndPermission)
	for _, menuAndPermission := range menuAndPermissions {
		if m, exist := menuAndPermissionMap[menuAndPermission.ID]; exist {
			menuAndPermissionMap[menuAndPermission.ID] = append(m, menuAndPermission)
		} else {
			menuAndPermissionMap[menuAndPermission.ID] = []*system.MenuAndPermission{menuAndPermission}
		}
	}
	return
}

func (ms *MenuService) GetUserMenu2(userId string) (menus []*model2.SystemMenu, err error) {
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
