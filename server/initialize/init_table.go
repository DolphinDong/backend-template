package initialize

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"

	"github.com/pkg/errors"
)

// 迁移表
func Migrate() {
	global.Logger.Info("start migrate ......")
	err := global.DB.AutoMigrate()
	tools.CheckErr(errors.Wrap(err, "migrate error"))
	global.Logger.Info("migrate success !!!")
}
