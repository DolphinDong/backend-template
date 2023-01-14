package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/DolphinDong/backend-template/common/response"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/model/dao/redis"
	"github.com/DolphinDong/backend-template/service/system"
	"github.com/DolphinDong/backend-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type LoginController struct {
	LoginService *system.LoginService
}

func NewLoginController() *LoginController {
	return &LoginController{
		LoginService: system.NewLoginService(),
	}
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lc *LoginController) Login(ctx *gin.Context) {
	params := &Login{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "read user info failed"))
		response.ResponseHttpError(ctx, err.Error())
		return
	}
	user, err := lc.LoginService.Login(params.Username, params.Password, ctx.ClientIP())
	if err != nil {
		global.Logger.Errorf("%+v", errors.WithMessage(err, "login failed"))
		response.ResponseHttpError(ctx, "系统异常 登录失败")
		return
	}
	// 登录成功
	if user != nil {
		token, err := tools.CreateToken([]byte(tools.SecretKey), user.ID, constant.JwTPeriod)
		if err != nil {
			global.Logger.Errorf("%+v", errors.WithMessage(err, "create token failed"))
			response.ResponseHttpError(ctx, "系统异常 登录失败")
			return
		}
		err = redis.NewRedisDao().SetKeyWithExpiration(tools.GetRedisTokenKey(constant.TokenRedisPrefix, user.ID, token), user.ID, constant.TokenPeriod*60)
		if err != nil {
			global.Logger.Errorf("%+v", errors.WithMessage(err, "set token to redis failed"))
			response.ResponseHttpError(ctx, "系统异常 登录失败")
			return
		}
		response.ResponseOkWithData(ctx, map[string]string{
			"token": token,
		})
		return
	} else {
		global.Logger.Error("登录失败")
		response.ResponseHttpErrorWithInfo(ctx, "登录失败")
	}
}

func (lc *LoginController) Logout(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constant.TokenHeader)
	if token != "" {
		err := lc.LoginService.Logout(token)
		if err != nil {
			global.Logger.Errorf("%+v", errors.WithMessage(err, "delete token failed"))
			//response.ResponseHttpError(ctx, "logout failed")
			return
		}
	}
	response.ResponseOkWithMessage(ctx, "logout success")
}
