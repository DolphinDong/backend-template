package model

import "gorm.io/gorm"

// 菜单
type SystemMenu struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ParentId   int       `json:"parentId" gorm:"not null;comment:父级菜单ID"`
	Name       string    `json:"name" gorm:"not null;size:255;uniqueIndex:name"`
	Path       string    `json:"path" gorm:"not null;comment:菜单路径;size:255"`
	Component  string    `json:"component" gorm:"not null;size:255;comment:菜单对应的组件"`
	Redirect   string    `json:"redirect" gorm:"size:255;comment:重定向地址"`
	Permission []string  `json:"permission" gorm:"-"` // 改菜单下对应的API权限
	Sort       int       `json:"sort" gorm:"size:11;comment:菜单排序"`
	Meta       *MenuMeta `json:"meta" gorm:"-"`
	gorm.Model `json:"-"`
}
type MenuMeta struct {
	ID                  int    `json:"id" gorm:"primaryKey;autoIncrement"`
	MenuID              int    `json:"menu_id" gorm:"not null;index;unique;size:20"`
	Title               string `json:"title" gorm:"not null;comment:菜单名称;size:50"`
	Icon                string `json:"icon" gorm:"comment:菜单名称;size:50"`
	Target              string `json:"target" gorm:"size:50"`
	Show                bool   `json:"show" gorm:"default:true"`
	HideChildren        bool   `json:"hideChildren" gorm:""`
	HiddenHeaderContent bool   `json:"hiddenHeaderContent" gorm:""`
	gorm.Model          `json:"-"`
}
type Permission struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	MenuID       int    `json:"menu_id" gorm:"not null;index"`
	Describe     string `json:"describe" gorm:"not null;size:255;comment:权限描述"`
	Identify     string `json:"identify" gorm:"not null;size:255;uniqueIndex:identify_action;comment:权限标识"`
	Action       string `json:"action" gorm:"not null;size:50;uniqueIndex:identify_action;comment:动作"`
	DefaultCheck bool   `json:"defaultCheck" gorm:"-"`
	gorm.Model   `json:"-"`
}
