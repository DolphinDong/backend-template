package initialize

import (
	"github.com/DolphinDong/backend-template/global"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/DolphinDong/backend-template/tools"

	"github.com/pkg/errors"
)

// 迁移表
func Migrate() {
	global.Logger.Info("start migrate ......")
	err := global.DB.AutoMigrate(
		&model2.SystemMenu{},
		&model2.MenuMeta{},
		&model2.Permission{},
	)
	tools.CheckErr(errors.Wrap(err, "migrate error"))
	global.Logger.Info("migrate success !!!")
}
