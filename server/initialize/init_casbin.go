package initialize

import (
	"github.com/DolphinDong/backend-template/global"
	model2 "github.com/DolphinDong/backend-template/model"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 初始化casbin
func InitCasbin() {
	global.Logger.Info("start init casbin .....")
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.DB, &model2.CasbinRule{})
	if err != nil {
		global.Logger.Errorf("create casbin adapter error: %+v", err)
	}
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && method_match(r.act , p.act)
m = g(r.sub, p.sub) && keyMatch2(r.obj , p.obj) && method_match(r.act , p.act)
m = g(r.sub, p.sub) && keyMatch(r.obj , p.obj) && method_match(r.act , p.act)
`)
	if err != nil {
		global.Logger.Fatalf("error: model: %+v", err)
	}
	global.Enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		global.Logger.Fatalf("error: enforcer: %+v", err)
	}
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		global.Logger.Fatalf("load policy error: %+v", err)
	}
	global.Enforcer.AddFunction("method_match", requestMethodMatchFunc)
	global.Logger.Info("init casbin success!!")
}

func requestMethodMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(requestMethodMatch(name1, name2)), nil
}

// POST > GET 的权限
// key1 请求进来的权限
// key2 数据库中的权限
func requestMethodMatch(key1 string, key2 string) bool {
	// 不区分大小写
	key1 = strings.ToUpper(key1)
	key2 = strings.ToUpper(key2)

	var MethodMap = map[string]int{
		"GET":  1,
		"POST": 2,
	}

	key1Num := MethodMap[key1]
	key2Num := MethodMap[key2]
	if key1Num != 0 { // 说明请求方法为 GET 或 POST
		if key2Num >= key1Num { // 如果数据库中的权限 >= 请求的权限则直接返回true
			return true
		} else {
			return false
		}
	} else { // 如果验证方式不是通过请求方式则就按照正常的比较即可
		return key1 == key2
	}
}
