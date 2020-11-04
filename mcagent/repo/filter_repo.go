package repo

import (
	"cmpService/common/mcmodel"
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/mciptables"
	"fmt"
)

var GlobalAccessFilterRepo []mcmodel.McFilterRule

func InitCachingAccessFilter() {
	rules, err := GetAccessFilterAll()
	if err != nil {
		fmt.Println("InitCachingAccessFilter: error -", err)
		return
	}
	GlobalAccessFilterRepo = rules
	fmt.Println("InitCachingAccessFilter:")
	cfg := config2.GetMcGlobalConfig()
	for _, rule := range rules {
		rule.Dump()
		mciptables.AddFFilterWrap(rule.IpAddr+"/32", cfg.McagentPort, "virbr0")
	}
}

func GetAccessFilterAll() ([]mcmodel.McFilterRule, error) {
	rule, err := GetAccessFilterAll2Db()
	if err != nil {
		fmt.Println("GetAccessFilterAll error1: ", err)
	}
	return rule, err
}

func GetAccessFilter(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	rule, err := GetAccessFilter2Db(obj)
	if err != nil {
		fmt.Println("GetAccessFilter error1: ", err)
	}
	return rule, err
}

func AddAccessFilter(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	rule, err := AddAccessFilter2Db(obj)
	if err != nil {
		fmt.Println("AddAccessFilter error: ", err)
	} else {
		cfg := config2.GetMcGlobalConfig()
		mciptables.AddFFilterWrap(rule.IpAddr+"/32", cfg.McagentPort, "virbr0")
	}
	return rule, err
}

func DeleteAccessFilter(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	rule, err := GetAccessFilter2Db(obj)
	if err != nil {
		fmt.Println("DeleteAccessFilter error1: ", err)
		return rule, err
	}
	_, err = DeleteAccessFilter2Db(rule)
	if err != nil {
		fmt.Println("DeleteAccessFilter error2: ", err)
	} else {
		cfg := config2.GetMcGlobalConfig()
		mciptables.DeleteFFilterWrap(rule.IpAddr+"/32", cfg.McagentPort)
	}
	return rule, err
}

// Access Filter
func AddAccessFilter2Db(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	obj.Idx = 0
	return config2.GetMcGlobalConfig().DbOrm.AddMcFilterRule(obj)
}

func DeleteAccessFilter2Db(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	return config2.GetMcGlobalConfig().DbOrm.DeleteMcFilterRule(obj)
}

func GetAccessFilterAll2Db() ([]mcmodel.McFilterRule, error) {
	return config2.GetMcGlobalConfig().DbOrm.
		GetMcFilterRule()
}

func GetAccessFilter2Db(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	return config2.GetMcGlobalConfig().DbOrm.
		GetMcFilterRuleBySerialNumberAndAddr(obj.SerialNumber, obj.IpAddr)
}
