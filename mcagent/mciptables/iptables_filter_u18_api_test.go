package mciptables

import "testing"

func TestGetFilterForwardAllRuleU18(t *testing.T) {
	GetFilterForwardAllRule()
}

var u18virIfName = "br0"
func TestGetFilterForwardRuleByIfNameU18(t *testing.T) {
	GetFilterForwardRuleByIfName(u18virIfName)
}

func TestGetFilterForwardRejectAllRuleU18(t *testing.T) {
	GetFilterForwardRejectAllRule()
}

func TestGetFilterForwardRejectRuleU18(t *testing.T) {
	GetFilterForwardRejectRuleByIfName(virIfName)
}

func TestDeleteFilterForwardRejectRuleU18(t *testing.T) {
	DeleteFilterForwardRejectRuleUbuntu18(virIfName)
}

/**
 * FORWARD -d 211.58.83.151/32
 */
var u18addr = "211.58.83.151/32"
var u18intf= "br0"
func TestAddFilterForwardIpv4AddrRuleU18(t *testing.T) {
	AddFilterForwardIpv4AddrRule(u18addr, u18intf)
}

func TestDeleteFilterForwardIpv4AddrRuleU18(t *testing.T) {
	DeleteFilterForwardIpv4AddrRule(u18addr, u18intf)
}

/**
 * REJECT --reject-with icmp-port-unreachable
 */
func TestAddFilterForwardRejectRuleU18(t *testing.T) {
	AddFilterForwardRejectRule("virbr0")
}

func TestDeleteFilterForwardRejectAllRuleRejectU18(t *testing.T) {
	DeleteFilterForwardRejectAllRule()
}

/**
 * Wrapping Func
 */
var u18addr01 = "211.58.83.151/32"
var u18intf01= "wlo1"
var u18virIfName01= "virbr0"
func TestAddFFilterWrapUbuntu18(t *testing.T) {
	AddFFilterWrapUbuntu18(u18addr01, u18intf01, u18virIfName01)
}

func TestDeleteFFilterWrapUbuntu18(t *testing.T) {
	DeleteFFilterWrapUbuntu18(u18addr01, u18intf01)
}
