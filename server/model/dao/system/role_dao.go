package system

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RoleDao struct {
	*gorm.DB
}

func (rd *RoleDao) QueryRole(page int, size int, search string) (roles []*model.Role, total int64, err error) {
	result := rd.Model(&model.Role{})
	if search != "" {
		result.Where("role_name like ? or role_identify like ?", tools.GetQueryString(search),
			tools.GetQueryString(search))
	}
	if err = result.Count(&total).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}
	if err = result.Offset((page - 1) * size).Limit(size).Find(&roles).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return
}

func (rd *RoleDao) AddRole(role *model.Role) (err error) {
	if err = rd.Create(&role).Error; err != nil {
		return errors.WithStack(err)
	}
	return
}

func (rd *RoleDao) UpdateRole(role *model.Role) (err error) {
	if err = rd.Select("role_name").Updates(role).Error; err != nil {
		return errors.WithStack(err)
	}
	return
}

func (rd *RoleDao) QueryRoleById(id uint) (role *model.Role, err error) {
	if err = rd.Where("id=?", id).First(&role).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (rd *RoleDao) DeleteRoleById(db *gorm.DB, id uint) error {
	if db == nil {
		db = rd.DB
	}
	if err := db.Unscoped().Where("id=?", id).Delete(&model.Role{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (rd *RoleDao) DeleteRolePermission(db *gorm.DB, identify string) error {
	if db == nil {
		db = rd.DB
	}
	// 删除角色的权限
	if err := db.Table("casbin_rule").Where("ptype='p' and v0=?", identify).Delete(&model.CasbinRule{}).Error; err != nil {
		return errors.WithStack(err)
	}
	// 删除角色对应的用户记录
	if err := db.Table("casbin_rule").Where("ptype='g' and v1=?", identify).Delete(&model.CasbinRule{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (rd *RoleDao) QueryAllRoles() (roles []*model.Role, err error) {
	if err = rd.Find(&roles).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (rd *RoleDao) QueryRoleByIds(ids []string) (roles []*model.Role, err error) {
	if err = rd.Where("id in ?", ids).Find(&roles).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func NewRoleDao() *RoleDao {
	return &RoleDao{
		global.DB,
	}
}
