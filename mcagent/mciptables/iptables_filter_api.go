package mciptables

import (
	"fmt"
	"strings"
)

func ConvertPrefix(ip string) string {
	arr := strings.Split(ip, ".")
	prefix := fmt.Sprintf("%s.%s.%s.0/24",
		arr[0], arr[1], arr[2])
	return prefix
}

func DeleteFilterForwardRejectAllRule() {
	/**
	 * Ubuntu 18.04
	 */
	//DeleteFilterForwardRejectAllRuleUbuntu18()

	/**
	 * Ubuntu 20.04 : BEGIN Rule
	 */
	DeleteFilterForwardRejectAllRuleUbuntu20()

	// Add Default Rule
	AddDefaultFilterForwardRule()

	/**
	 * Ubuntu 20.04 : END Rule
	 */
}

func AddFFilterWrap(addr, ifName, virIfName string){
	/**
	 * Ubuntu 18.04
	 */
	//AddFFilterWrapUbuntu18(addr, ifName, virIfName)

	/**
	 * Ubuntu 20.04
	 */
	AddFFilterWrapUbuntu20(addr, ifName, virIfName)
}

func DeleteFFilterWrap(addr, ifName string){
	/**
	 * Ubuntu 18.04
	 */
	//DeleteFFilterWrapUbuntu18(addr, ifName)

	/**
	 * Ubuntu 20.04
	 */
	DeleteFFilterWrapUbuntu20(addr, ifName)
}
