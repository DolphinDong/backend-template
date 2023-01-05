package tools

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strings"
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
