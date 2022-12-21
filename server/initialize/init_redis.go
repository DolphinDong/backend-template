package initialize

import (
	"fmt"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
)

const (
	defaultMaxIdle        = 5
	defaultMaxActive      = 100
	defaultConnectTimeout = 10  //秒
	defaultIdleTimeout    = 100 //秒
)

func initRedis() {
	global.Logger.Info("start init redis pool.....")
	maxIdle := global.Config.Redis.PoolMaxIdle
	if maxIdle == 0 {
		maxIdle = defaultMaxIdle
	}

	maxActive := global.Config.Redis.PoolMaxActive
	if maxActive == 0 {
		maxActive = defaultMaxActive
	}
	if global.Config.Redis.ConnectTimeOut == 0 {
		global.Config.Redis.ConnectTimeOut = defaultConnectTimeout
	}

	if global.Config.Redis.IdleTimeout == 0 {
		global.Config.Redis.IdleTimeout = defaultIdleTimeout
	}
	address := fmt.Sprintf("%v:%v", global.Config.Redis.RedisHost, global.Config.Redis.RedisPort)
	global.RedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Second * time.Duration(global.Config.Redis.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				address,
				redis.DialPassword(global.Config.Redis.RedisPassword),
				redis.DialDatabase(global.Config.Redis.RedisDB),
				redis.DialConnectTimeout(time.Second*time.Duration(global.Config.Redis.ConnectTimeOut)))
		},
	}
	_, err := global.RedisPool.Get().Do("ping")
	tools.CheckErr(errors.Wrap(err, "init redis pool error"))
	global.Logger.Info("start init redis pool success!!")
}
