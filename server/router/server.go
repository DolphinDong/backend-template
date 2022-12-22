package router

import (
	"fmt"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/initialize"
	"github.com/DolphinDong/backend-template/middleware"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	ginlogrus "github.com/toorop/gin-logrus"
	"net/http"
	"time"
)

func init() {
	initialize.Init()
}

func Run() {
	engine := gin.New()
	// 注册中间件
	registerMiddleware(engine)
	// 注册路由
	registerRouter(engine)
	err := engine.Run(fmt.Sprintf("%v:%v", global.Config.ServerHost, global.Config.ServerPort))
	tools.CheckErr(errors.Wrap(err, "start server error!!"))
}

// 注册中间件
func registerMiddleware(engine *gin.Engine) {

	engine.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodPatch, http.MethodPut, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:  []string{"x-requested-with", "X-Custom-Header", "accept", "Content-Type", "Access-Token", "Authorization", "responsetype"},
		ExposeHeaders: []string{"Content-Length", "Content-Disposition"},
		MaxAge:        time.Second * time.Duration(3600),
	}))

	engine.Use(ginlogrus.Logger(global.Logger))
	// 注册日志中间件
	//engine.Use(middleware.LogMiddleware())
	// recover
	engine.Use(middleware.Recover())
	// 登录检测
	//engine.Use(middleware.LoginCheck())

	// 打印请求信息
	engine.Use(middleware.PrintRequestInfo())
	// 权限检测
	engine.Use(middleware.PermissionCheck())
}
