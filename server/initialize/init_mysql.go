package initialize

import (
	"fmt"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	dsn = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
)

func initMysql() {
	global.Logger.Info("start connecting to the database....")
	mysqlDsn := fmt.Sprintf(dsn,
		global.Config.Mysql.MysqlUser,
		global.Config.Mysql.MysqlPassword,
		global.Config.Mysql.MysqlHost,
		global.Config.Mysql.MysqlPort,
		global.Config.Mysql.DBName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别 logger.Error
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	DB, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: mysqlDsn,
		}),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		tools.CheckErr(errors.Wrap(err, "connect to mysql error"))
	}

	global.DB = DB
	sqlDB, err := global.DB.DB()
	if err != nil {
		tools.CheckErr(errors.Wrap(err, "connect to mysql error"))
	}
	if global.Config.Mysql.MaxIdleConns != 0 {
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	}
	if global.Config.Mysql.MaxOpenConns != 0 {
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	}
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.Logger.Info("start connecting to the database success!!!")
}
