package mciptables

import "testing"

func dnatRule () *DnatRule {
	return &DnatRule{
		"192.168.122.99",
		"3389",
		"192.168.0.89",
		"13389",
	}
}

func TestAddDNATRule(t *testing.T) {
	rule := dnatRule()
	AddDNATRule(rule)
}

func TestGetNATRule(t *testing.T) {
	GetNATRule()
}

func TestDeleteDNATRule(t *testing.T) {
	rule := dnatRule()
	DeleteDNATRule(rule)
}
