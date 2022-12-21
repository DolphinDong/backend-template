package initialize

import (
	"fmt"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	_ "github.com/zput/zxcTool/ztLog/zt_formatter"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	logFileName     = "server.log"
	taskLogFileName = "dolphin_task.log"
	DEBUG           = "DEBUG"
	INFO            = "INFO"
	ERROR           = "ERROR"
)

// 初始化基础的日志
func initBaseLog() {
	global.Logger = logrus.New()
	/*
		GlobalLoger.SetReportCaller(true)
		GlobalLoger.SetFormatter(&zt_formatter.ZtFormatter{
			Formatter: nested.Formatter{
				TimestampFormat: "2006-01-02 15:04:05",
				NoColors:        true,
				NoFieldsColors:  true,
				ShowFullLevel:   true,
			},
		})
	*/
	global.Logger.Formatter = &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		NoFieldsColors:  true,
		ShowFullLevel:   true,
	}
}

// 初始化pro版的日志
func initProLog() {
	if global.Config.Log.LogPath == "" {
		global.Config.Log.LogPath = "logs"
	}
	if global.Config.Log.LogLevel == "" {
		global.Config.Log.LogLevel = INFO
	}
	// 创建日志存放的目录
	err := os.MkdirAll(global.Config.Log.LogPath, 0775)
	tools.CheckErr(errors.Wrap(err, "make log dir err"))
	filePath := filepath.Join(global.Config.Log.LogPath, logFileName)
	// 设置日志打印位置
	logfile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	tools.CheckErr(errors.Wrap(err, "open "+filePath+" err"))
	global.Logger.SetOutput(io.MultiWriter(logfile, os.Stdout))
	// 设置日志级别
	logLevel := strings.ToUpper(global.Config.Log.LogLevel)
	switch logLevel {
	case DEBUG:
		global.Logger.SetLevel(logrus.DebugLevel)
	case INFO:
		global.Logger.SetLevel(logrus.InfoLevel)
	case ERROR:
		global.Logger.SetLevel(logrus.ErrorLevel)
	default:
		tools.CheckErr(errors.New(fmt.Sprintf("log.log_level=\"%v\" is illegal", global.Config.Log.LogLevel)))
	}
	global.Logger.Info("init logger success!!")
}
