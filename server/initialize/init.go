package initialize

import (
	"flag"
	. "github.com/DolphinDong/backend-template/global"
	"github.com/common-nighthawk/go-figure"
)

const (
	logoInfo = "backend-template"
	info     = `Version: %v		Author: liudong`
	version  = "v1.0.0"
)

func Init() {
	// 初始化log
	initBaseLog()
	// 初始化配置文件
	var configPath string
	flag.StringVar(&configPath, "c", "config.json", "The configuration file path for server")
	flag.Parse()
	initConfig(configPath)

	initProLog()

	Logger.Info("starting server....")
	myFigure := figure.NewFigure(logoInfo, "big", true)
	Logger.Info(myFigure.String())

	Logger.Infof(info, version)
	// 初始化数据库连接
	initMysql()
	// 初始化redis连接
	initRedis()
	// 数据迁移
	Migrate()
	// 初始化casbin
	// InitCasbin()
	Logger.Info("server init success!!!")
}
