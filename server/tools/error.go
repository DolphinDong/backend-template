package tools

import "github.com/DolphinDong/backend-template/global"

func CheckErr(err error) {
	if err != nil {
		global.Logger.Fatalf("%+v", err)
	}
}
