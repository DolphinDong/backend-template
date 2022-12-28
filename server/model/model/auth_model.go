package model

import "gorm.io/gorm"

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:50;uniqueIndex:unique_index"`
	V0    string `gorm:"size:100;uniqueIndex:unique_index"`
	V1    string `gorm:"size:100;uniqueIndex:unique_index"`
	V2    string `gorm:"size:100;uniqueIndex:unique_index"`
	V3    string `gorm:"size:100;uniqueIndex:unique_index"`
	V4    string `gorm:"size:100;uniqueIndex:unique_index"`
	V5    string `gorm:"size:100;uniqueIndex:unique_index"`
	V6    string `gorm:"size:100;uniqueIndex:unique_index"`
	V7    string `gorm:"size:100;uniqueIndex:unique_index"`
}

type User struct {
	ID            string `gorm:"primaryKey;size:50" json:"id" `
	LoginName     string `gorm:"size:50;not null;comment:用户ID;unique" json:"login_name" validate:"min=5,max=30,required"`
	Username      string `gorm:"size:255;not null;comment:用户名" json:"username" validate:"required"`
	Gender        int    `gorm:"size:5;comment:性别" json:"gender" validate:"required"`
	Avatar        string `gorm:"size:255;comment:头像地址" json:"avatar"`
	PhoneNumber   string `gorm:"size:11;not null;comment:手机号码;unique" json:"phone_number" validate:"len=11,required"`
	Password      string `gorm:"size:255;not null;comment:密码" json:"-"`
	Email         string `gorm:"size:50;not null;comment:邮箱" json:"email" validate:"email,required"`
	IsAdmin       bool   `gorm:"comment:是否是超级管理员;default:false" json:"is_admin"`
	Status        bool   `gorm:"comment:状态;default:true" json:"status"`
	LastLoginTime int64  `gorm:"comment:上次登录时间;" json:"last_login_time"`
	LastLoginIp   string `gorm:"comment:上次登录Ip;size:50" json:"last_login_ip"`
	Role          map[string]interface {
	} `gorm:"-" json:"role"`
	gorm.Model `json:"-"`
}

// 角色表
type Role struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id,string"`
	RoleName     string `gorm:"size:100;comment:角色名称;not null;uniqueIndex" json:"role_name"`
	RoleIdentify string `gorm:"size:100;comment:角色标识;not null;uniqueIndex" json:"role_identify"`
	gorm.Model
}
