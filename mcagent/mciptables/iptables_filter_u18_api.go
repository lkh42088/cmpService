package mciptables

import (
	"fmt"
	"strings"
)

/**
 * Ubuntu 18.04
 *
 * Reference filter rule config. (e.g. iptables-save > ref.conf)
 *
-A FORWARD -d 192.168.122.0/24 -o virbr0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A FORWARD -s 192.168.122.0/24 -i virbr0 -j ACCEPT
-A FORWARD -i virbr0 -o virbr0 -j ACCEPT
-A FORWARD -o virbr0 -j REJECT --reject-with icmp-port-unreachable
-A FORWARD -i virbr0 -j REJECT --reject-with icmp-port-unreachable
 *
*/

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

func GetFilterForwardAllRule() bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardAllRule: error", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		fmt.Printf("%s\n", rule)
		arr := strings.Fields(rule)
		arr = arr[2:]
	}
	return true
}

func GetFilterForwardRuleByIfName(ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRuleByIfName: error", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
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

func GetFilterForwardRejectAllRule() bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRejectAllRule: error", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
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

func GetFilterForwardRejectRuleByIfName(ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterForwardRejectRuleByIfName: error", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
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

func DeleteFilterForwardRejectAllRuleUbuntu18() {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterList: error", err)
		return
	}
	table := "filter"
	chain := "FORWARD"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		if strings.Contains(rule, "REJECT") {
			fmt.Printf("%s\n", rule)
			arr := strings.Fields(rule)
			arr = arr[2:]
			//fmt.Printf("new %s\n", arr)
			err := ipt.Delete(table, chain, arr[0], arr[1], arr[2], arr[3])
			if err != nil {
				fmt.Println("error:", err)
			}
		}
	}
}

func DeleteFilterForwardRejectRuleUbuntu18(ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("DeleteFilterForwardRejectRule: error", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
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

func AddFilterForwardRejectRule(ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("AddFilterForwardRejectRule error:", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
	err = ipt.Append(table, chain, "-o", ifName, "-j", "REJECT",
		"--reject-with", "icmp-port-unreachable")
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error 1:", err)
		return false
	}
	err = ipt.Append(table, chain, "-i", ifName, "-j", "REJECT",
		"--reject-with", "icmp-port-unreachable")
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error 2:", err)
		return false
	}
	return true
}

func AddFilterForwardIpv4AddrRule(addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
	err = ipt.Append(table, chain, "-d", addr, "-o", ifName, "-m", "conntrack",
		"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error 1:", err)
		return false
	}
	err = ipt.Append(table, chain, "-s", addr, "-i", ifName, "-j", "ACCEPT")
	if err != nil {
		fmt.Println("AddFilterForwardIpv4AddrRule error 2:", err)
		return false
	}
	return true
}

func DeleteFilterForwardIpv4AddrRule(addr, ifName string) bool {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error:", err)
		return false
	}
	table := "filter"
	chain := "FORWARD"
	err = ipt.Delete(table, chain, "-d", addr, "-o", ifName, "-m", "conntrack",
		"--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 1:", err)
		return false
	}
	err = ipt.Delete(table, chain, "-s", addr, "-i", ifName, "-j", "ACCEPT")
	if err != nil {
		fmt.Println("DeleteFilterForwardIpv4AddrRule error 2:", err)
		return false
	}
	return true
}

func AddFFilterWrapUbuntu18(addr, ifName, virIfName string){
	// Delete REJECT Rule
	DeleteFilterForwardRejectRuleUbuntu18(virIfName)
	// Add Filter Rule
	AddFilterForwardIpv4AddrRule(addr, ifName)
	// Add REJECT Rule
	AddFilterForwardRejectRule(virIfName)
}

func DeleteFFilterWrapUbuntu18(addr, ifName string){
	// Delete Filter Rule
	DeleteFilterForwardIpv4AddrRule(addr, ifName)
}

