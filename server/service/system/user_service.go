package system

import (
	"fmt"
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/structs"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/dao/system"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type UserService struct {
	UserDao *system.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		UserDao: system.NewUserDao(),
	}
}

func (us *UserService) GetUserInfo(userId string) (userInfo *model.User, err error) {
	userInfo, err = us.UserDao.QueryUserInfoByUserID(userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if userInfo == nil {
		return
	}
	userAllPermission := make([]map[string]interface{}, 0)
	userPermissionsMap, err := us.GetPermissionsMap(userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for obj, permissions := range userPermissionsMap {
		p := make(map[string]interface{})
		p["permissionId"] = obj
		if len(permissions) > 0 {
			p["actionEntitySet"] = permissions
			userAllPermission = append(userAllPermission, p)
		}
	}
	userInfo.Role = map[string](interface{}){
		"permissions": userAllPermission,
	}
	return

}
func (us *UserService) GetPermissionsMap(userId string) (permissionsMap map[string][]*model.Permission, err error) {
	userPermissionsMap, err := tools.QueryPermissionByUserID(userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	identifies := []string{}
	// 通过Identifies去查询对应权限
	for identify, actions := range userPermissionsMap {
		// 如果这个权限只有查看菜单的权限则跳过，因为permission表中不会有菜单的查看权限的信息，减少冗余的查询
		if len(actions) == 1 && actions[0] == constant.MenuAct {
			continue
		}
		identifies = append(identifies, identify)
	}
	permissionDao := system.NewPermissionDao()
	permissions, err := permissionDao.QueryPermissionsByIdentifies(identifies)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	permissionsMap = make(map[string][]*model.Permission)
	for _, permission := range permissions {
		// 如果用户的权限表中有查出来的这个action 才选择添加到用户的permissionsMap中
		if tools.ElementInSlice(permission.Action, userPermissionsMap[permission.Identify]) {
			if _, ok := permissionsMap[permission.Identify]; ok {
				permissionsMap[permission.Identify] = append(permissionsMap[permission.Identify], permission)
			} else {
				permissionsMap[permission.Identify] = []*model.Permission{permission}
			}
		}

	}
	return
}

func (us *UserService) GetUsers(query *structs.TableQuery, gender, status string) (tableResponse *structs.TableResponse, err error) {
	tableResponse = new(structs.TableResponse)
	users, total, err := us.UserDao.QueryUser(query.Page, query.PageSize, query.Search, gender, status)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tableResponse.Data = users
	tableResponse.Total = total
	return
}

func (us *UserService) AddUser(user *model.User) (err error) {
	user.ID = tools.GetUUID()
	user.Password = tools.GetEncryptedPassword(tools.MD5Str(constant.UserDefaultPassword))
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = us.UserDao.AddUser(tx, user)
		if err != nil {
			return errors.WithStack(err)
		}
		userCasbin := &model.CasbinRule{
			Ptype: "g",
			V0:    user.ID,
			V1:    constant.UserDefaultRole,
		}
		err = system.NewCasbinDao().AddCasbinRows(tx, []*model.CasbinRule{userCasbin})
		//_, err = global.Enforcer.AddGroupingPolicy(user.ID, constant.UserDefaultRole)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return
}

func (us *UserService) UpdateUser(user *model.User) (err error) {
	err = us.UserDao.UpdateUserInfo(user)
	if err != nil {
		return errors.WithStack(err)
	}
	return
}

func (us *UserService) ResetUserPassword(user *model.User) (err error) {
	user.Password = tools.GetEncryptedPassword(tools.MD5Str(constant.UserDefaultPassword))
	err = us.UserDao.UpdateUserPassword(user)
	if err != nil {
		return errors.WithStack(err)
	}
	return
}

func (us *UserService) DeleteUser(userId string) (err error) {

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err2 := us.UserDao.DeleteUserById(tx, userId)
		if err2 != nil {
			return errors.WithStack(err)
		}
		// 采用sql的方式直接删除用户的角色和权限
		err2 = us.UserDao.DeleteUserPermissionById(tx, userId)
		if err2 != nil {
			return errors.WithStack(err)
		}
		//// 删除用户的权限
		//_, err2 = global.Enforcer.RemoveFilteredPolicy(0, userId)
		//if err2 != nil {
		//	return errors.WithStack(err)
		//}
		//// 删除用户的角色
		//_, err2 = global.Enforcer.RemoveNamedGroupingPolicy("g", userId)
		//if err2 != nil {
		//	return errors.WithStack(err)
		//}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return
}

func (us *UserService) GetReqPermission(req string) (permissions []interface{}, err error) {
	menuService := NewMenuService()
	// 查询客户拥有的菜单
	systemMenus, err := menuService.MenuDao.QueryAllMenu()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, menu := range systemMenus {
		// 如果有这个权限
		if global.Enforcer.HasPolicy(req, menu.Name, constant.MenuAct) {
			permissions = append(permissions, menu.ID)
		}
	}

	// 查询用户拥有的权限
	allPermission, err := menuService.MenuDao.QueryAllPermission()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, permission := range allPermission {
		// 如果有这个权限
		if global.Enforcer.HasPolicy(req, permission.Identify, permission.Action) {
			permissions = append(permissions, fmt.Sprintf("p%v", permission.ID))
		}
	}

	return
}

func (us *UserService) UpdateReqPermission(req string, permissions []interface{}) (err error) {
	menuDao := system.NewMenuDao()
	menuIds := []int64{}
	permissionIds := []int64{}
	// 将ID分成菜单和权限
	for _, pid := range permissions {
		switch pid.(type) {
		case float64:
			menuIds = append(menuIds, int64(pid.(float64)))
		case string:
			if !strings.Contains(pid.(string), "p") {
				continue
			}
			idStr := strings.ReplaceAll(pid.(string), "p", "")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return errors.Errorf("Invalid parameter: ", err.Error())
			}
			permissionIds = append(permissionIds, id)
		}
	}
	allCasbinRule := []*model.CasbinRule{}
	systemMenus, err := menuDao.QuerySystemMenuByIds(menuIds)
	if err != nil {
		return errors.WithStack(err)
	}
	queryPermissions, err := menuDao.QueryPermissionByIds(permissionIds)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, systemMenu := range systemMenus {
		allCasbinRule = append(allCasbinRule, &model.CasbinRule{
			Ptype: "p",
			V0:    req,
			V1:    systemMenu.Name,
			V2:    constant.MenuAct,
		})
	}
	for _, permission := range queryPermissions {
		allCasbinRule = append(allCasbinRule, &model.CasbinRule{
			Ptype: "p",
			V0:    req,
			V1:    permission.Identify,
			V2:    permission.Action,
		})
	}
	casbinDao := system.NewCasbinDao()
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = casbinDao.DeleteCasbinRuleByReq(tx, req)
		if err != nil {
			return errors.WithStack(err)
		}
		err = casbinDao.AddCasbinRows(tx, allCasbinRule)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return
}
