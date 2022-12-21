package initialize

import (
	"encoding/json"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"io/ioutil"
)

// 初始化配置文件
func initConfig(configPath string) {
	global.Config = new(global.Server)
	global.Logger.Infof("start init config \"%v\"...", configPath)
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		tools.CheckErr(errors.Wrap(err, "read config error"))
	}
	err = json.Unmarshal(content, global.Config)
	if err != nil {
		tools.CheckErr(errors.Wrap(err, "Unmarshal config error"))
	}
	global.Logger.Infof("init config success :%v ", string(content))
}
