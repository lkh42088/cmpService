package mciptables

import (
	"cmpService/mcagent/config"
	"fmt"
	"strings"
)

/**
 * Ubuntu 20.04
 *
 * Reference filter rule config. (e.g. iptables-save > ref.conf)
 *
-A LIBVIRT_FWI -d 192.168.122.0/24 -o virbr0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A LIBVIRT_FWI -o virbr0 -j REJECT --reject-with icmp-port-unreachable
-A LIBVIRT_FWO -s 192.168.122.0/24 -i virbr0 -j ACCEPT
-A LIBVIRT_FWO -i virbr0 -j REJECT --reject-with icmp-port-unreachable
-A LIBVIRT_FWX -i virbr0 -o virbr0 -j ACCEPT
 *
*/

func GetFilterForwardAllRuleFromChain(chain string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardAllRule: error", err)
		return false
	}
	table := "filter"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		fmt.Printf("%s\n", rule)
		arr := strings.Fields(rule)
		arr = arr[2:]
	}
	return true
}

func GetFilterForwardRuleByIfNameFromChain(chain, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRuleByIfName: error", err)
		return false
	}
	table := "filter"
	//chain := "FORWARD"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		if strings.Contains(rule, ifName) {
			fmt.Printf("%s\n", rule)
			arr := strings.Fields(rule)
			arr = arr[2:]
		}
	}
	return true
}

func GetFilterForwardRejectAllRuleFromChain(chain string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRejectAllRule: error", err)
		return false
	}
	table := "filter"
	//chain := "FORWARD"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		if strings.Contains(rule, "REJECT") {
			arr := strings.Fields(rule)
			arr = arr[2:]
			fmt.Printf("%s\n", rule)
		}
	}
	return true
}

func CheckFilterAddrRule(chain, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRejectRuleByIfName: error", err)
		return false
	}
	table := "filter"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		if strings.Contains(rule, addr) {
			arr := strings.Fields(rule)
			fmt.Printf("Rule: %s\n", rule)
			//fmt.Printf(">> Array: %d, %s\n", len(arr), arr)
			ruleAddr:= arr[3]
			//fmt.Printf(">> IP: %s\n", ruleAddr)
			ruleIfName:= arr[5]
			//fmt.Printf(">> IFNAME: %s\n", ruleIfName)
			if addr == ruleAddr && ifName == ruleIfName {
				fmt.Printf(">> Find it\n")
				return true
			}
		}
	}
	return false
}

func GetFilterForwardRejectRuleByIfNameFromChain(chain, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRejectRuleByIfName: error", err)
		return false
	}
	table := "filter"
	//chain := "FORWARD"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		if strings.Contains(rule, "REJECT") {
			arr := strings.Fields(rule)
			arr = arr[2:]
			if arr[1] != ifName {
				continue
			}
			fmt.Printf("%s\n", rule)
		}
	}
	return true
}

func DeleteFilterForwardRejectRuleFromChain(chain, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("DeleteFilterForwardRejectRule: error", err)
		return false
	}
	table := "filter"
	//chain := "FORWARD"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		if strings.Contains(rule, "REJECT") {
			fmt.Printf("%s\n", rule)
			arr := strings.Fields(rule)
			arr = arr[2:]
			if arr[1] != ifName {
				continue
			}
			fmt.Printf("new %s\n", arr)
			fmt.Printf("ifName %s\n", arr[1])
			err := ipt.Delete(table, chain, arr[0], arr[1], arr[2], arr[3])
			if err != nil {
				fmt.Println("error:", err)
			}
		}
	}
	return true
}

func AddFilterForwardRejectRuleFromChain(chainIn, chainOut, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("AddFilterForwardRejectRuleFromChain error:", err)
		return false
	}
	table := "filter"
	//chain := "FORWARD"
	err = ipt.Append(table, chainIn, "-o", ifName, "-j", "REJECT",
		"--reject-with", "icmp-port-unreachable")
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRuleFromChain error 1:", err)
		return false
	}
	err = ipt.Append(table, chainOut, "-i", ifName, "-j", "REJECT",
		"--reject-with", "icmp-port-unreachable")
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRuleFromChain error 2:", err)
		return false
	}
	return true
}

func AddFilterForwardIpv4AddrRuleFromChainNew(chainIn, chainOut, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	/**
	 * Check Chain In
	 */
	//var haveFilterRule = false
	var haveFilterRule = CheckFilterAddrRule(chainIn, addr, ifName)
	if haveFilterRule == false {
		//err = ipt.Append(table, chainIn, "-d", addr, "-o", ifName, "-m", "conntrack",
		//	"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
		err = ipt.Append(table, chainIn, "-d", addr, "-o", ifName, "-j", "ACCEPT")
		if err != nil {
			fmt.Println("AddFilterForwardIpv4AddrRule error 1:", err)
		}
	}
	/**
	 * Check Chain Out
	 */
	haveFilterRule = CheckFilterAddrRule(chainOut, addr, ifName)
	if haveFilterRule == false {
		err = ipt.Append(table, chainOut, "-s", addr, "-i", ifName, "-j", "ACCEPT")
		if err != nil {
			fmt.Println("AddFilterForwardIpv4AddrRule error 2:", err)
		}
	}
	return true
}

func AddFilterForwardIpv4AddrRuleFromChain(chainIn, chainOut, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	/**
	 * Check Chain In
	 */
	//var haveFilterRule = false
	var haveFilterRule = CheckFilterAddrRule(chainIn, addr, ifName)
	if haveFilterRule == false {
		err = ipt.Append(table, chainIn, "-d", addr, "-o", ifName, "-m", "conntrack",
			"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
		if err != nil {
			fmt.Println("AddFilterForwardIpv4AddrRule error 1:", err)
		}
	}
	/**
	 * Check Chain Out
	 */
	haveFilterRule = CheckFilterAddrRule(chainOut, addr, ifName)
	if haveFilterRule == false {
		err = ipt.Append(table, chainOut, "-s", addr, "-i", ifName, "-j", "ACCEPT")
		if err != nil {
			fmt.Println("AddFilterForwardIpv4AddrRule error 2:", err)
		}
	}
	return true
}

func AddFilterForwardIpv4AddrRuleFromChainReverse(chainIn, chainOut, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	/**
	 * Check Chain In
	 */
	//var haveFilterRule = false
	var haveFilterRule = CheckFilterAddrRule(chainIn, addr, ifName)
	if haveFilterRule == false {
		err = ipt.Append(table, chainOut, "-d", addr, "-o", ifName, "-m", "conntrack",
			"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
		if err != nil {
			fmt.Println("AddFilterForwardIpv4AddrRule error 1:", err)
		}
	}
	/**
	 * Check Chain Out
	 */
	haveFilterRule = CheckFilterAddrRule(chainOut, addr, ifName)
	if haveFilterRule == false {
		err = ipt.Append(table, chainIn, "-s", addr, "-i", ifName, "-j", "ACCEPT")
		if err != nil {
			fmt.Println("AddFilterForwardIpv4AddrRule error 2:", err)
		}
	}
	return true
}

func DeleteFilterForwardIpv4AddrRuleFromChainNew(chainIn, chainOut, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	/**
	 * Check Chain In
	 */
	//err = ipt.Delete(table, chainIn, "-d", addr, "-o", ifName, "-m", "conntrack",
	//	"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
	err = ipt.Delete(table, chainIn, "-d", addr, "-o", ifName, "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 1:", err)
		//return false
	}
	/**
	 * Check Chain Out
	 */
	err = ipt.Delete(table, chainOut, "-s", addr, "-i", ifName, "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 2:", err)
		//return false
	}
	return true
}

func DeleteFilterForwardIpv4AddrRuleFromChain(chainIn, chainOut, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	/**
	 * Check Chain In
	 */
	err = ipt.Delete(table, chainIn, "-d", addr, "-o", ifName, "-m", "conntrack",
		"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 1:", err)
		//return false
	}
	/**
	 * Check Chain Out
	 */
	err = ipt.Delete(table, chainOut, "-s", addr, "-i", ifName, "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 2:", err)
		//return false
	}
	return true
}

func DeleteFilterForwardIpv4AddrRuleFromChainReverse(chainIn, chainOut, addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	/**
	 * Check Chain In
	 */
	err = ipt.Delete(table, chainOut, "-d", addr, "-o", ifName, "-m", "conntrack",
		"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 1:", err)
		//return false
	}
	/**
	 * Check Chain Out
	 */
	err = ipt.Delete(table, chainIn, "-s", addr, "-i", ifName, "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 2:", err)
		//return false
	}
	return true
}

func DeleteFilterForwardRejectAllRuleUbuntu20() {
	var ifName = "virbr0"
	var chainIn = "LIBVIRT_FWI"
	var chainOut = "LIBVIRT_FWO"

	// Delete REJECT Rule
	DeleteFilterForwardRejectRuleFromChain(chainIn, ifName)
	DeleteFilterForwardRejectRuleFromChain(chainOut, ifName)
}

func AddDefaultFilterForwardRule() {
	cfg := config.GetMcGlobalConfig()
	prefix := ConvertPrefix(cfg.SystemInfo.IP)
	ifName := cfg.SystemInfo.IfName
	virtIfName := "virbr0"
	AddFFilterWrapUbuntu20(prefix, ifName, virtIfName)
}

func AddFFilterWrapUbuntu20(addr, ifName, virIfName string){
	var chainIn = "LIBVIRT_FWI"
	var chainOut = "LIBVIRT_FWO"

	// Delete REJECT Rule
	DeleteFilterForwardRejectRuleFromChain(chainIn, virIfName)
	DeleteFilterForwardRejectRuleFromChain(chainOut, virIfName)

	// Add Filter Rule
	//AddFilterForwardIpv4AddrRuleFromChain(chainIn, chainOut, addr, ifName)
	AddFilterForwardIpv4AddrRuleFromChainNew(chainIn, chainOut, addr, ifName)

	// Add REJECT Rule
	AddFilterForwardRejectRuleFromChain(chainIn, chainOut, virIfName)
}

func AddFFilterWrapUbuntu20Reverse(addr, ifName, virIfName string){
	var chainIn = "LIBVIRT_FWI"
	var chainOut = "LIBVIRT_FWO"

	// Delete REJECT Rule
	DeleteFilterForwardRejectRuleFromChain(chainIn, virIfName)
	DeleteFilterForwardRejectRuleFromChain(chainOut, virIfName)

	// Add Filter Rule
	AddFilterForwardIpv4AddrRuleFromChainReverse(chainIn, chainOut, addr, ifName)

	// Add REJECT Rule
	AddFilterForwardRejectRuleFromChain(chainIn, chainOut, virIfName)
}

func DeleteFFilterWrapUbuntu20(addr, ifName string){
	var chainIn = "LIBVIRT_FWI"
	var chainOut = "LIBVIRT_FWO"

	// Delete Filter Rule
	//DeleteFilterForwardIpv4AddrRuleFromChain(chainIn, chainOut, addr, ifName)
	DeleteFilterForwardIpv4AddrRuleFromChainNew(chainIn, chainOut, addr, ifName)
}

func DeleteFFilterWrapUbuntu20Reverse(addr, ifName string){
	var chainIn = "LIBVIRT_FWI"
	var chainOut = "LIBVIRT_FWO"

	// Delete Filter Rule
	DeleteFilterForwardIpv4AddrRuleFromChainReverse(chainIn, chainOut, addr, ifName)
}

