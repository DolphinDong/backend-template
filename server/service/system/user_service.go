package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/structs"
	"github.com/DolphinDong/backend-template/model/dao/system"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
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

func (us *UserService) GetUsers(query *structs.TableQuery, gender, isAdmin, status string) (tableResponse *structs.TableResponse, err error) {
	tableResponse = new(structs.TableResponse)
	users, total, err := us.UserDao.QueryUser(query.Page, query.PageSize, query.Search, gender, isAdmin, status)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tableResponse.Data = users
	tableResponse.Total = total
	return
}
