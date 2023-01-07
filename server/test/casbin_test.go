package test

import (
	"fmt"
	"github.com/DolphinDong/backend-template/global"
	"github.com/DolphinDong/backend-template/initialize"
	"github.com/DolphinDong/backend-template/tools"
	"testing"
)

func TestCasbin(t *testing.T) {
	initialize.Init()
	//rules := [][]string{
	//	[]string{"liudong", "data4_admin"},
	//	[]string{"liudong", "data3_admin"},
	//}
	//_, _ = global.Enforcer.AddGroupingPolicies(rules)
	//
	//rules = [][]string{
	//	[]string{"liudong", "/test1", "get"},
	//	[]string{"data4_admin", "/test4", "post"},
	//	[]string{"data3_admin", "/test3", "put"},
	//	[]string{"data3_admin", "/test3", "delete"},
	//}
	//
	//global.Enforcer.AddPolicies(rules)

	//enforce, err := global.Enforcer.Enforce("liudong", "/test/data3", "write")
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Printf("校验结果：%v",enforce)
	res, err := global.Enforcer.Enforce("c5c471fd-07ee-426a-bb7a-b3ede8804d71", "role", "show_menu")
	fmt.Println(res, err)
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
			if !tools.ElementInSlice(act, acts) {
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
			if !tools.ElementInSlice(act, acts) {
				res[obj] = append(acts, act)
			}
		} else {
			res[obj] = []string{act}
		}
	}
	return res
}
