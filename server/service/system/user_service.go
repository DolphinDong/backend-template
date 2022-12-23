package system

import (
	"github.com/DolphinDong/backend-template/model/dao/system"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
)

type UserService struct {
	UserDao *system.UserDao
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
	userPermissionsMap := tools.QueryPermissionByUserID(userId)
	permissionDao := system.NewPermissionDao()
	for obj, acts := range userPermissionsMap {
		p := make(map[string]interface{})
		p["permissionId"] = obj
		permissions, err := permissionDao.QueryPermissionsByIdentifyAndActions(obj, acts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
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

func NewUserService() *UserService {
	return &UserService{
		UserDao: system.NewUserDao(),
	}
}
