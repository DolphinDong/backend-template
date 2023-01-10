package system

import (
	"fmt"
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/model/dao/system"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"sort"
)

const (
	PermissionType = 2
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
	// 第一次校验权限需要重新加载权限表中的数据
	isFirst := true
	for _, menuAndPermissions := range menuAndPermissionMap {
		// 获取第0个
		m1 := menuAndPermissions[0]
		hasPermission, err := tools.HasPermission(userId, m1.Name, constant.MenuAct, isFirst)
		isFirst = false
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
		hasPermission, err := tools.HasPermission(userId, menu.Name, constant.MenuAct, false)
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

func (ms *MenuService) GetMenus() (menuTree []*system.MenuAndPermission, menuIds []int, err error) {
	allMenus, err := ms.MenuDao.QueryAllMenus()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	allPermission, err := ms.MenuDao.QueryAllPermission()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	allMenus = ms.AppendMenuPermissions(allMenus, allPermission)
	menuTree = ms.BuildMenuTree(0, allMenus)
	// 获取所有菜单的ID, 方便前端展开树状结构
	for _, m := range allMenus {
		if m.ParentId == 0 && m.Type != PermissionType {
			menuIds = append(menuIds, m.ID)
		}

	}
	return
}

// 给菜单添加上子权限
func (ms *MenuService) AppendMenuPermissions(menus []*system.MenuAndPermission, permissions []*model2.Permission) []*system.MenuAndPermission {
	result := make([]*system.MenuAndPermission, 0, len(menus))
	restPermission := permissions
	var menuPermissions []*model2.Permission
	for _, menu := range menus {
		menuPermissions, restPermission = ms.PickMenuPermissions(menu.ID, restPermission)
		permissions = restPermission
		menuChildren := make([]*system.MenuAndPermission, 0, len(menuPermissions))
		for _, menuPermission := range menuPermissions {
			menuChildren = append(menuChildren, &system.MenuAndPermission{
				ID:    menuPermission.ID,
				Type:  PermissionType, // 权限的类型为2
				Name:  fmt.Sprintf("%v : %v", menuPermission.Identify, menuPermission.Action),
				Title: menuPermission.Describe,
			})
		}
		menu.Children = menuChildren
		result = append(result, menu)
	}
	// 遍历剩余的
	for _, p := range restPermission {
		result = append(result, &system.MenuAndPermission{
			ID:    p.ID,
			Type:  PermissionType, // 权限的类型为2
			Name:  fmt.Sprintf("%v : %v", p.Identify, p.Action),
			Title: p.Describe,
		})
	}
	return result
}

// 查询菜单的子权限并返回剩余的权限
func (ms *MenuService) PickMenuPermissions(menuId int, permissions []*model2.Permission) (menuPermissions, restPermissions []*model2.Permission) {
	for _, p := range permissions {
		// 如果当前权限为menuId的子权限则添加到menuPermissions, 否则添加到剩余的权限中restPermissions
		if p.MenuID == menuId {
			menuPermissions = append(menuPermissions, p)
		} else {
			restPermissions = append(restPermissions, p)
		}
	}
	return
}
func (ms *MenuService) BuildMenuTree(menuId int, allMenus []*system.MenuAndPermission) (menuTree []*system.MenuAndPermission) {
	for _, menu := range allMenus {
		if menu.ParentId == menuId {
			if menu.Type == 1 {
				menu.Children = append(menu.Children, ms.BuildMenuTree(menu.ID, allMenus)...)
			}
			menuTree = append(menuTree, menu)
		}
	}
	return
}
