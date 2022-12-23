package system

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (ud *UserDao) QueryUserInfoByUserID(userId string) (userInfo *model.User, err error) {
	result := ud.Where("id=?", userId).Find(&userInfo)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return
}

func NewUserDao() *UserDao {
	return &UserDao{
		global.DB,
	}
}
