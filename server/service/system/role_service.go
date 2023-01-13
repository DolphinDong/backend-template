package system

import (
	"github.com/DolphinDong/backend-template/common/structs"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/dao/system"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RoleService struct {
	RoleDao *system.RoleDao
}

func (rs *RoleService) GetRoles(query *structs.TableQuery) (tableResponse *structs.TableResponse, err error) {
	tableResponse = new(structs.TableResponse)
	roles, total, err := rs.RoleDao.QueryRole(query.Page, query.PageSize, query.Search)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tableResponse.Data = roles
	tableResponse.Total = total
	return
}

func (rs *RoleService) AddRole(role *model.Role) error {
	err := rs.RoleDao.AddRole(role)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (rs *RoleService) UpdateRole(role *model.Role) error {
	if err := rs.RoleDao.UpdateRole(role); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (rs *RoleService) DeleteRole(role *model.Role) error {
	r, err := rs.RoleDao.QueryRoleById(role.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = rs.RoleDao.DeleteRoleById(tx, r.ID)
		if err != nil {
			return errors.WithStack(err)
		}
		err = rs.RoleDao.DeleteRolePermission(tx, r.RoleIdentify)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (rs *RoleService) GetRolePermissions(roleId int64) (permissions []interface{}, err error) {
	role, err := rs.RoleDao.QueryRoleById(uint(roleId))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	userService := NewUserService()
	permissions, err = userService.GetReqPermission(role.RoleIdentify)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (rs *RoleService) UpdateRolePermission(id int, permissions []interface{}) error {
	role, err := rs.RoleDao.QueryRoleById(uint(id))
	if err != nil {
		return errors.WithStack(err)
	}
	userService := NewUserService()
	err = userService.UpdateReqPermission(role.RoleIdentify, permissions)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func NewRoleService() *RoleService {
	return &RoleService{
		RoleDao: system.NewRoleDao(),
	}
}
