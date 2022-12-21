package tools

import (
	"github.com/DolphinDong/backend-template/global"
	"github.com/pkg/errors"
)

// 鉴权, ctx 中包含了当前用户的信息
func HasPermission(sub, obj, act string) (bool, error) {

	err := global.Enforcer.LoadPolicy()
	if err != nil {
		return false, errors.WithStack(err)
	}

	ok, err := global.Enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ok, nil
}

func QueryPermissionByUserID(userid string) map[string][]string {
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
	return res
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
