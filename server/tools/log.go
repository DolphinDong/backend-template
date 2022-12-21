package tools

import (
	"github.com/DolphinDong/backend-template/common/constant"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 从context 中获取logger
func GetLogger(ctx *gin.Context) *logrus.Entry {
	if logger, exists := ctx.Get(constant.LoggerKey); exists {
		if l, ok := logger.(*logrus.Entry); ok {
			return l
		} else {
			return &logrus.Entry{}
		}
	} else {
		return &logrus.Entry{}
	}
}
