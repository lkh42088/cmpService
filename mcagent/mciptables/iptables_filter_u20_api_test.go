package mciptables

import (
	"fmt"
	"testing"
)

/**
 * Wrapping Func
 */
//var u20addr01 = "211.58.83.151/32" // home public
var u20addr01 = "221.148.133.88/32" // company public
//var u20addr01 = "192.168.254.0/24" // home local
//var u20intf01= "br0"
var u20intf01= "wlo1"
//var u20intf01= "virbr0"
var u20virIfName01= "virbr0"
func TestAddFFilterWrapUbuntu20(t *testing.T) {
	AddFFilterWrapUbuntu20(u20addr01, u20intf01, u20virIfName01)
}

func TestDeleteFFilterWrapUbuntu20(t *testing.T) {
	DeleteFFilterWrapUbuntu20(u20addr01, u20intf01)
}

/**
 * Reverse
 */
func TestAddFFilterWrapUbuntu20Reverse(t *testing.T) {
	AddFFilterWrapUbuntu20Reverse(u20addr01, u20intf01, u20virIfName01)
}

func TestDeleteFFilterWrapUbuntu20Reverse(t *testing.T) {
	DeleteFFilterWrapUbuntu20Reverse(u20addr01, u20intf01)
}

/**
 * Show Rules
 */
func TestGetFilterForwardAllRuleFromChain(t *testing.T) {
	var chainIn = "LIBVIRT_FWI"
	var chainOut = "LIBVIRT_FWO"
	GetFilterForwardAllRuleFromChain(chainIn)
	GetFilterForwardAllRuleFromChain(chainOut)
}

/**
 * Delete  REJECT
 */
func TestDeleteFilterForwardRejectAllRuleUbuntu20(t *testing.T) {
	DeleteFilterForwardRejectAllRuleUbuntu20()
}

func TestCheckFilterRule(t *testing.T) {
	var chain = "LIBVIRT_FWI"
	var addr = "192.168.254.0/24"
	var ifName = "wlo1"
	res := CheckFilterAddrRule(chain, addr, ifName)
	fmt.Println("result:", res)
}
