package errors

import "github.com/pkg/errors"

var (
	HasSubMenuError =errors.New("该菜单存在子菜单无法删除")
)

