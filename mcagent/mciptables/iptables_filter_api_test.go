package mciptables

import "testing"

func TestGetFilterForwardAllRule(t *testing.T) {
	GetFilterForwardAllRule()
}

var virIfName = "br0"
func TestGetFilterForwardRuleByIfName(t *testing.T) {
	GetFilterForwardRuleByIfName(virIfName)
}

func TestGetFilterForwardRejectAllRule(t *testing.T) {
	GetFilterForwardRejectAllRule()
}

func TestGetFilterForwardRejectRule(t *testing.T) {
	GetFilterForwardRejectRuleByIfName(virIfName)
}

func TestDeleteFilterForwardRejectRule(t *testing.T) {
	DeleteFilterForwardRejectRule(virIfName)
}

/**
 * FORWARD -d 211.58.83.151/32
 */
var addr = "211.58.83.151/32"
var intf= "br0"
func TestAddFilterForwardIpv4AddrRule(t *testing.T) {
	AddFilterForwardIpv4AddrRule(addr, intf)
}

func TestDeleteFilterForwardIpv4AddrRule(t *testing.T) {
	DeleteFilterForwardIpv4AddrRule(addr, intf)
}

/**
 * REJECT --reject-with icmp-port-unreachable
 */
func TestAddFilterForwardRejectRule(t *testing.T) {
	AddFilterForwardRejectRule("virbr0")
}

func TestDeleteFilterForwardRejectAllRuleReject(t *testing.T) {
	DeleteFilterForwardRejectAllRule()
}

/**
 * Wrapping Func
 */
var addr01 = "211.58.83.151/32"
var intf01= "br0"
var virIfName01= "virbr0"
func TestAddFFilterWrap(t *testing.T) {
	AddFFilterWrap(addr01, intf01, virIfName01)
}

func TestDeleteFFilterWrap(t *testing.T) {
	DeleteFFilterWrap(addr01, intf01)
}
