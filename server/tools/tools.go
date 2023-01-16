package tools

import (
	"fmt"
	"github.com/DolphinDong/backend-template/global"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"math/rand"
	"strings"
	"time"
)

// 校验数据
func Validate(struc interface{}) error {

	err := global.GlobalValidator.Struct(struc)
	if err != nil {
		msg := "Invalid parameter: "
		es := err.(validator.ValidationErrors)
		errFields := []string{}
		for _, e := range es {
			errFields = append(errFields, e.Field())
		}
		// 拼接不符合要求的字段
		msg += strings.Join(errFields, ",")
		return errors.New(msg)
	}
	return nil
}

func GetUUID() (uuidStr string) {
	u := uuid.New()
	key := u.String()
	return key
}

func GetQueryString(query string) string {
	return fmt.Sprintf("%%%v%%", query)
}

// 获取用户的token在redis中的key
func GetRedisTokenKey(prefix, userId, token string) string {
	return fmt.Sprintf("%v_%v_%v", prefix, userId, token)
}

func GenerateUUID(l int) string {
	s := "0123456789ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvxwyz"
	rand.Seed(time.Now().UnixNano())
	buffer := make([]byte, l)
	for i := 0; i < l; i++ {
		buffer[i] = s[rand.Intn(len(s))]
	}
	return string(buffer)
}
