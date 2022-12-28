package system

import (
	"fmt"
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

func (ud *UserDao) QueryUser(page int, size int, search, gender, isAdmin, status string) (users []*model.User, total int64, err error) {
	result := ud.Model(&model.User{})
	if search != "" {
		result.Where("login_name like ? or  username like ? or phone_number like ? or email like ?", fmt.Sprintf("%%%v%%", search),
			fmt.Sprintf("%%%v%%", search),
			fmt.Sprintf("%%%v%%", search), fmt.Sprintf("%%%v%%", search))
	}
	if gender == "1" || gender == "2" {
		result.Where("gender=?", gender)
	}
	if isAdmin == "1" || isAdmin == "0" {
		result.Where("is_admin=?", isAdmin)
	}
	if status == "1" || status == "0" {
		result.Where("status=?", status)
	}
	if result.Count(&total).Error != nil {
		return nil, 0, errors.WithStack(err)
	}
	result.Offset((page - 1) * size).Limit(size)
	if result.Find(&users).Error != nil {
		return nil, 0, errors.WithStack(err)
	}
	return
}

func NewUserDao() *UserDao {
	return &UserDao{
		global.DB,
	}
}
