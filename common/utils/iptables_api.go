package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

var gipt *IPTables

func GetIpt() (*IPTables, error) {
	var err error
	if gipt == nil {
		gipt, err = New()
		if err != nil {
			fmt.Println("GetFilterList: error", err)
			return gipt, err
		}
	}
	return gipt, nil
}

func GetFilterList() {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterList: error", err)
	}
	ipt.Dump()
	listChain, err := ipt.ListChains("filter")
	for _, chain := range listChain {
		fmt.Printf("%s\n", chain)
	}
	list, err := ipt.List("filter", "FORWARD")
	for _, rule := range list{
		if strings.Contains(rule, "REJECT") {
			fmt.Printf("%s\n", rule)
			arr := strings.Fields(rule)
			arr = arr[2:]
			fmt.Printf("new %s\n", arr)
			err := ipt.Delete("filter", "FORWARD", arr[0], arr[1], arr[2], arr[3])
			if err != nil {
				fmt.Println("error:", err)
			}
		}
	}
}

func DeleteFilterReject() {
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

type DnatRule struct {
	ToAddr   string
	ToPort   string
	WantAddr string
	WantPort string
}

func (r *DnatRule) Dump() string {
	pretty, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (o *DnatRule) Compare(n *DnatRule) bool {
	if o.ToAddr != n.ToAddr { return false }
	if o.ToPort != n.ToPort { return false }
	if o.WantAddr != n.WantAddr { return false }
	if o.WantPort != n.WantPort { return false }
	return true
}

func AddDNATRule(rule *DnatRule) {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterList: error", err)
		return
	}
	fmt.Println("AddDNATRule:")
	rule.Dump()

	table := "nat"
	chain := "PREROUTING"
	//sudo iptables -t nat -A PREROUTING -d 192.168.0.57 -p tcp --dport 13389 -j DNAT --to 192.168.122.130:3389
	//sudo iptables -t nat -A PREROUTING -d 192.168.0.57 -p udp --dport 13389 -j DNAT --to 192.168.122.130:3389
	//-A PREROUTING -d 192.168.0.89/32 -p tcp -m tcp --dport 15000 -j DNAT --to-destination 192.168.122.99:3389
	//-A PREROUTING -d 192.168.0.89/32 -p udp -m udp --dport 15000 -j DNAT --to-destination 192.168.122.99:3389
	ipt.Append(table, chain, "-d", rule.WantAddr, "-p", "tcp",
		"--dport", rule.WantPort, "-j", "DNAT", "--to", rule.ToAddr+":"+rule.ToPort)
	ipt.Append(table, chain, "-d", rule.WantAddr, "-p", "udp",
		"--dport", rule.WantPort, "-j", "DNAT", "--to", rule.ToAddr+":"+rule.ToPort)
}

func DeleteDNATRule(rule *DnatRule) {
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterList: error", err)
	}
	table := "nat"
	chain := "PREROUTING"
	ipt.Delete(table, chain, "-d", rule.WantAddr, "-p", "tcp",
		"--dport", rule.WantPort, "-j", "DNAT", "--to", rule.ToAddr+":"+rule.ToPort)
	ipt.Delete(table, chain, "-d", rule.WantAddr, "-p", "udp",
		"--dport", rule.WantPort, "-j", "DNAT", "--to", rule.ToAddr+":"+rule.ToPort)
}

func GetNATRule() []string{
	ipt, err := GetIpt()
	if err != nil {
		fmt.Println("GetFilterList: error", err)
	}
	table := "nat"
	chain := "PREROUTING"
	list, err := ipt.List(table, chain)
	for _, rule := range list{
		fmt.Printf("%s\n", rule)
	}
	return list
}