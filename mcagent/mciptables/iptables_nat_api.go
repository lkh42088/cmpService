package mciptables

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

const (
	DNAT_NEXT_DST_IP = 1
	DNAT_NEXT_DPORT = 2
	DNAT_NEXT_TO_DEST = 3
)

func GetDnatRuleConfigByRule(rule string) *DnatRule {
	var dnat DnatRule
	if strings.Contains(rule, "DNAT") == false {
		return nil
	}
	arr := strings.Fields(rule)
	var next int
	for _, obj := range arr {
		//fmt.Println("GetDnatRuleConfigByRule:", obj)
		if next > 0 {
			switch next {
			case DNAT_NEXT_DST_IP:
				if strings.Contains(obj, "/") {
					tmp := strings.Split(obj, "/")
					dnat.WantAddr = tmp[0]
				} else {
					dnat.WantAddr = obj
				}
			case DNAT_NEXT_DPORT:
				dnat.WantPort = obj
			case DNAT_NEXT_TO_DEST:
				tmp := strings.Split(obj, ":")
				dnat.ToAddr = tmp[0]
				dnat.ToPort = tmp[1]
			default:
			}
			next = 0
			continue
		}
		if obj == "-d" {
			next = DNAT_NEXT_DST_IP
		} else if obj == "--dport" {
			next = DNAT_NEXT_DPORT
		} else if obj == "--to-destination" {
			next = DNAT_NEXT_TO_DEST
		} else {
			next = 0
		}
	}
	return &dnat
}

func GetDnatList() *[]DnatRule {
	var DNATList []DnatRule
	natList := GetNATRule()
	for _, rule := range natList {
		dnat := GetDnatRuleConfigByRule(rule)
		if dnat != nil {
			DNATList = append(DNATList, *dnat)
		}
	}
	return &DNATList
}

func DeleteAllDnat() {
	list := GetDnatList()
	if list == nil {
		return
	}

	for _, rule := range *list {
		DeleteDNATRule(&rule)
	}
}