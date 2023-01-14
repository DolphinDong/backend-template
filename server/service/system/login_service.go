package system

import (
	"github.com/DolphinDong/backend-template/model/dao/redis"
	"github.com/DolphinDong/backend-template/model/dao/system"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"time"
)

type LoginService struct {
	UserDao *system.UserDao
}

func NewLoginService() *LoginService {
	return &LoginService{
		UserDao: system.NewUserDao(),
	}
}
func (ls *LoginService) Login(username, password, ip string) (user *model.User, err error) {
	password = tools.GetEncryptedPassword(password)
	user, err = ls.UserDao.QueryUserByLoginNameAndPwd(username, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if user != nil {
		user.LastLoginIp = ip
		user.LastLoginTime = time.Now().Unix()
		err = ls.UserDao.UpdateUserLoginInfo(user)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return
}

func (ls *LoginService) Logout(token string) error {
	if err := redis.NewRedisDao().DeleteKey(token); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
