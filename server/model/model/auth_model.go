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
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId        string `gorm:"size:50;not null;comment:用户ID;unique" json:"user_id"`
	Username      string `gorm:"size:255;comment:用户名" json:"username"`
	Gender        string `gorm:"size:5;comment:性别" json:"gender"`
	PhoneNumber   string `gorm:"size:11;not null;comment:手机号码;unique" json:"phone_number"`
	Password      string `gorm:"size:255;not null;comment:密码" json:"password"`
	Email         string `gorm:"size:255;comment:邮箱" json:"email"`
	IsAdmin       bool   `gorm:"comment:是否是超级管理员;default:false" json:"is_admin"`
	LastLoginTime int    `gorm:"comment:上次登录时间;" json:"last_login_time"`
	LastLoginIp   int    `gorm:"comment:上次登录Ip;" json:"last_login_ip"`
	gorm.Model
}

// 角色表
type Role struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id,string"`
	RoleName     string `gorm:"size:100;comment:角色名称;not null;uniqueIndex" json:"role_name"`
	RoleIdentify string `gorm:"size:100;comment:角色标识;not null;uniqueIndex" json:"role_identify"`
	gorm.Model
}
