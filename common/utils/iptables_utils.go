package utils

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
)


func GetFilterList() {
	ipt, err := iptables.New()
	if err != nil {
		fmt.Println("GetFilterList: error", err);
	}
	listChain, err := ipt.ListChains("filter")
	for _, chain := range listChain {
		fmt.Printf("%s\n", chain)
	}
}

