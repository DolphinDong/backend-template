package system

import (
	"github.com/DolphinDong/backend-template/global"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MenuDao struct {
	*gorm.DB
}

func NewMenuDao() *MenuDao {
	return &MenuDao{
		global.DB,
	}
}
func (md *MenuDao) QueryAllMenu() (menus []*model2.SystemMenu, err error) {
	if md.Order("sort desc").Find(&menus).Error != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (md *MenuDao) QueryMenuMetaByMenuID(menuId int) (menuMeta *model2.MenuMeta, err error) {
	result := md.Where("menu_id=?", menuId).Find(&menuMeta)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (md *MenuDao) QueryMenuPermissionsByMenuID(menuId int) (permissions []*model2.Permission, err error) {
	result := md.Where("menu_id=?", menuId).Find(&permissions)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
