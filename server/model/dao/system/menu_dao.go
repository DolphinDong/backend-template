package system

import (
	"github.com/DolphinDong/backend-template/global"
	model2 "github.com/DolphinDong/backend-template/model/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MenuDao struct {
	*gorm.DB
}

func NewMenuDao() *MenuDao {
	return &MenuDao{
		global.DB,
	}
}

type MenuAndPermission struct {
	ID        int    `json:"id"`
	ParentId  int    `json:"parentId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Sort      int    `json:"sort"`

	Title               string `json:"title"`
	Icon                string `json:"icon"`
	Target              string `json:"target"`
	Show                bool   `json:"show"`
	HideChildren        bool   `json:"hideChildren"`
	HiddenHeaderContent bool   `json:"hiddenHeaderContent"`

	Describe string               `json:"describe"`
	Identify string               `json:"identify"`
	Action   string               `json:"action"`
	Type     int                  `json:"type,omitempty"`
	Children []*MenuAndPermission `json:"children,omitempty" gorm:"-"`
}

func (md *MenuDao) QueryAllMenuAndPermission() (menuAndPermission []*MenuAndPermission, err error) {
	sql := `SELECT
	t1.*,
	t2.title,
	t2.icon,
	t2.target,
	t2.show,
	t2.hide_children,
	t2.hidden_header_content ,
	t3.describe,
	t3.identify,
	t3.action
FROM
	system_menus AS t1
	LEFT JOIN menu_meta AS t2 ON t1.id = t2.menu_id 
	LEFT JOIN permissions AS t3 ON t3.menu_id = t1.id`
	result := md.Raw(sql).Scan(&menuAndPermission)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return
}

func (md *MenuDao) QueryAllMenu() (menus []*model2.SystemMenu, err error) {
	result := md.Order("sort desc").Find(&menus)
	if result.Error != nil {
		err = errors.WithStack(result.Error)
		return
	}
	return
}

func (md *MenuDao) QueryMenuMetaByMenuID(menuId int) (menuMeta *model2.MenuMeta, err error) {
	result := md.Where("menu_id=?", menuId).Find(&menuMeta)
	if result.Error != nil {
		err = errors.WithStack(result.Error)
		return
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return
}

func (md *MenuDao) QueryMenuPermissionsByMenuID(menuId int) (permissions []*model2.Permission, err error) {
	result := md.Where("menu_id=?", menuId).Find(&permissions)
	if result.Error != nil {
		err = errors.WithStack(result.Error)
		return
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return
}

func (md *MenuDao) QueryAllMenus() (allMenus []*MenuAndPermission, err error) {
	sql := `SELECT
	t1.*,
	t2.title,
	t2.icon,
	t2.target,
	t2.SHOW as 'show',
	t2.hide_children,
	t2.hidden_header_content,
	1 as type
FROM
	system_menus AS t1
	LEFT JOIN menu_meta AS t2 ON t1.id = t2.menu_id
`
	result := md.Raw(sql).Scan(&allMenus)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return
}

func (md *MenuDao) QueryAllPermission() (permissions []*model2.Permission, err error) {
	if err = md.Find(&permissions).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
