package system

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PermissionDao struct {
	*gorm.DB
}

func NewPermissionDao() *PermissionDao {
	return &PermissionDao{global.DB}
}

func (pd *PermissionDao) QueryPermissionsByIdentifyAndActions(identify string, actions []string) (permissions []*model.Permission, err error) {
	result := pd.Where("identify=?", identify).Where("action in ?", actions).Find(&permissions)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return
}
