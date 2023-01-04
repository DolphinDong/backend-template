package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config          *Server
	Logger          *logrus.Logger
	DB              *gorm.DB
	RedisPool       *redis.Pool
	Enforcer        *casbin.Enforcer
	GlobalValidator = validator.New()
)
