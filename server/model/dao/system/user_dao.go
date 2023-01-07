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

func (ud *UserDao) QueryUser(page int, size int, search, gender, status string) (users []*model.User, total int64, err error) {
	result := ud.Model(&model.User{})
	if search != "" {
		result.Where("login_name like ? or  username like ? or phone_number like ? or email like ?", fmt.Sprintf("%%%v%%", search),
			fmt.Sprintf("%%%v%%", search),
			fmt.Sprintf("%%%v%%", search), fmt.Sprintf("%%%v%%", search))
	}
	if gender == "1" || gender == "2" {
		result.Where("gender=?", gender)
	}
	//if isAdmin == "1" || isAdmin == "0" {
	//	result.Where("is_admin=?", isAdmin)
	//}
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

func (ud *UserDao) QueryUserByLoginNameAndPwd(loginName, password string) (user *model.User, err error) {
	err = ud.Where("status=?", 1).Where("login_name=?", loginName).Where("password=?", password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return
}

func (ud *UserDao) UpdateUserLoginInfo(user *model.User) error {
	if err := ud.Select("last_login_ip", "last_login_time").Updates(user).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (ud *UserDao) AddUser(db *gorm.DB, user *model.User) error {
	err := db.Select("*").Create(user).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (ud *UserDao) UpdateUserInfo(user *model.User) error {
	if err := ud.Select("username", "gender", "phone_number", "email", "status").
		Updates(user).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (ud *UserDao) UpdateUserPassword(user *model.User) error {
	if err := ud.Select("password").
		Updates(user).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (ud *UserDao) DeleteUserById(db *gorm.DB, id string) (err error) {
	if db == nil {
		db = ud.DB
	}
	if err = db.Unscoped().Where("id=?", id).Delete(&model.User{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return
}

func (ud *UserDao) DeleteUserPermissionById(db *gorm.DB, id string) (err error) {
	if db == nil {
		db = ud.DB
	}
	if err = db.Table("casbin_rule").Where("v0=?", id).Delete(&model.CasbinRule{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return
}

func NewUserDao() *UserDao {
	return &UserDao{
		global.DB,
	}
}
