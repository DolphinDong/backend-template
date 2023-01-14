package system

import (
	"github.com/DolphinDong/backend-template/common/constant"
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
	Name      string `json:"name" validate:"required"`
	Path      string `json:"path"`
	Component string `json:"component" validate:"required_if=Type 1"`
	Redirect  string `json:"redirect"`
	Sort      int    `json:"sort"`

	Title               string `json:"title" validate:"required"`
	Icon                string `json:"icon"`
	Target              string `json:"target"`
	Show                bool   `json:"show"`
	HideChildren        bool   `json:"hideChildren"`
	HiddenHeaderContent bool   `json:"hiddenHeaderContent"`

	Describe string               `json:"describe"`
	Identify string               `json:"identify"`
	Action   string               `json:"action" validate:"required_if=Type 2"`
	Type     int                  `json:"type,omitempty" validate:"required,oneof=1 2"`
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
	order by t1.sort desc, t1.id asc
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

func (md *MenuDao) AddMenu(tx *gorm.DB, menu *model2.SystemMenu) (*model2.SystemMenu, error) {
	if tx == nil {
		tx = md.DB
	}
	if err := md.Create(&menu).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return menu, nil
}

func (md *MenuDao) AddMenuMeta(tx *gorm.DB, meta *model2.MenuMeta) error {
	if tx == nil {
		tx = md.DB
	}
	if err := md.Create(&meta).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) AddPermission(tx *gorm.DB, permission *model2.Permission) error {
	if tx == nil {
		tx = md.DB
	}
	if err := md.Create(&permission).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) UpdateMenu(tx *gorm.DB, m *model2.SystemMenu) error {
	if tx == nil {
		tx = md.DB
	}
	if err := md.Select("*").Omit("id").Where("id=?", m.ID).Updates(&m).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) UpdateMenuMeta(tx *gorm.DB, meta *model2.MenuMeta) error {
	if tx == nil {
		tx = md.DB
	}
	if err := md.Select("*").Omit("id", "hide_children", "hidden_header_content").Where("menu_id=?", meta.MenuID).Updates(&meta).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) QuerySystemMenuById(id int) (menu *model2.SystemMenu, err error) {
	if err = md.Where("id=?", id).First(&menu).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (md *MenuDao) UpdateMenuCasbin(db *gorm.DB, oldName, newName string) error {
	if db == nil {
		db = md.DB
	}
	sql := "update casbin_rule set v1=? where v1=? and v2=?"
	// 删除角色的权限
	if err := db.Exec(sql, newName, oldName, constant.MenuAct).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) UpdatePermission(tx *gorm.DB, p *model2.Permission) error {
	if tx == nil {
		tx = md.DB
	}
	if err := tx.Where("id=?", p.ID).Updates(&p).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) QueryMenuPermissionByID(id int) (permission *model2.Permission, err error) {
	if err = md.Where("id=?", id).First(&permission).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (md *MenuDao) UpdatePermissionCasbin(tx *gorm.DB, oldObj, newObj, oldAct, newAct string) error {
	if tx == nil {
		tx = md.DB
	}
	sql := "update casbin_rule set v1=?,v2=? where v1=? and v2=?"
	// 删除角色的权限
	if err := tx.Exec(sql, newObj, newAct, oldObj, oldAct).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) DeleteSystemMenuById(tx *gorm.DB, id int) (err error) {
	if tx == nil {
		tx = md.DB
	}
	if err = tx.Unscoped().Where("id=?", id).Delete(&model2.SystemMenu{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return
}

func (md *MenuDao) DeleteMenuMetaByMenuId(tx *gorm.DB, menuId int) error {
	if tx == nil {
		tx = md.DB
	}
	if err := tx.Unscoped().Where("menu_id=?", menuId).Delete(&model2.MenuMeta{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) DeleteCasbinByObjAct(tx *gorm.DB, obj, act string) error {
	if tx == nil {
		tx = md.DB
	}
	// 删除角色的权限
	if err := tx.Table("casbin_rule").Where("ptype='p' and v1=? and v2=?", obj, act).Delete(&model2.CasbinRule{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (md *MenuDao) DeletePermissionById(tx *gorm.DB, id int) (err error) {
	if tx == nil {
		tx = md.DB
	}
	if err := tx.Unscoped().Where("id=?", id).Delete(&model2.Permission{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return
}

func (md *MenuDao) QuerySystemMenuCountByParentId(id int) (count int64, err error) {
	if err = md.Model(model2.SystemMenu{}).Where("parent_id=?", id).Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}

	return
}

func (md *MenuDao) QuerySystemMenuByIds(ids []int64) (systemMenu []*model2.SystemMenu, err error) {
	if err = md.Where("id in ?", ids).Find(&systemMenu).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (md *MenuDao) QueryPermissionByIds(ids []int64) (permissions []*model2.Permission, err error) {
	if err = md.Where("id in ?", ids).Find(&permissions).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
