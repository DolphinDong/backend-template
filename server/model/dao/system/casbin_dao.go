package system

import (
	"github.com/DolphinDong/backend-template/global"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CasbinDao struct {
	*gorm.DB
}

func NewCasbinDao() *CasbinDao {
	return &CasbinDao{
		global.DB,
	}
}

func (cd *CasbinDao) DeleteCasbinRuleByReq(tx *gorm.DB, req string) error {
	if tx == nil {
		tx = cd.DB
	}
	tx.Table("casbin_rule").Where("ptype='p' and v0=?", req).Delete(&model2.CasbinRule{})
	return nil
}

func (cd *CasbinDao) AddCasbinPolicys(tx *gorm.DB, rule []*model2.CasbinRule) error {
	if tx == nil {
		tx = cd.DB
	}
	if err := tx.Table("casbin_rule").CreateInBatches(&rule, len(rule)).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
