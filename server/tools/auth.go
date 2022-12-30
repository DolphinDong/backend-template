package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/DolphinDong/backend-template/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strings"
	"time"
)

// 鉴权, ctx 中包含了当前用户的信息
func HasPermission(sub, obj, act string, needLoadPolicy bool) (bool, error) {
	if needLoadPolicy {
		err := global.Enforcer.LoadPolicy()
		if err != nil {
			return false, errors.WithStack(err)
		}
	}

	ok, err := global.Enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ok, nil
}

func QueryPermissionByUserID(userid string) (map[string][]string, error) {
	err := global.Enforcer.LoadPolicy()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	policy := global.Enforcer.GetFilteredPolicy(0, userid)
	filteredGroupingPolicy := global.Enforcer.GetFilteredGroupingPolicy(0, userid)
	for _, g := range filteredGroupingPolicy {
		p := global.Enforcer.GetFilteredPolicy(0, g[1])
		policy = append(policy, p...)
	}
	res := make(map[string][]string)
	for _, p := range policy {
		obj := p[1]
		act := p[2]
		if acts, ok := res[obj]; ok {
			// 不存在则添加
			if !ElementInSlice(act, acts) {
				res[obj] = append(acts, act)
			}
		} else {
			res[obj] = []string{act}
		}
	}
	return res, nil
}
func QueryPermissionByRoleID(roleid string) map[string][]string {
	policy := global.Enforcer.GetFilteredPolicy(0, roleid)
	res := make(map[string][]string)
	for _, p := range policy {
		obj := p[1]
		act := p[2]
		if acts, ok := res[obj]; ok {
			// 不存在则添加
			if !ElementInSlice(act, acts) {
				res[obj] = append(acts, act)
			}
		} else {
			res[obj] = []string{act}
		}
	}
	return res
}

const (
	SecretKey = "b7eabdec-683f-0eb2-ad3c-43e3c2220251"
	Salt      = "75737eb1-3a1d-629c-6617-2c3f85769599"
)

// 生成Token：
// SecretKey 是一个 const 常量
func CreateToken(SecretKey []byte, issuer string, periodMinutes int) (tokenString string, err error) {
	m := time.Duration(periodMinutes)
	claims := jwt.StandardClaims{
		ExpiresAt: int64(time.Now().Add(time.Minute * m).Unix()),
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

// 解析Token
func ParseToken(tokenSrt string, SecretKey []byte) (claims *jwt.StandardClaims, err error) {
	var token *jwt.Token
	token, err = jwt.ParseWithClaims(tokenSrt, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims.(*jwt.StandardClaims)
	return
}

// 生成密码
func GetEncryptedPassword(str string) string {
	str = strings.ToUpper(Salt) + str + Salt
	return MD5Str(str)
}
func MD5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
