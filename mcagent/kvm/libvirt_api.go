package kvm

import "fmt"

func GetMcVirtInfo() {
	// Get Vms Domains
	doms, err := GetDomainListAll()
	if err == nil {
		fmt.Println("--------------------------------")
		fmt.Println("VMs")
		for index, dom := range doms {
			name, _ := dom.GetName()
			fmt.Println(index, ". ", name)
		}
	}

	// Get Networks
	nets, err := GetAllNetwork()
	if err == nil {
		fmt.Println("--------------------------------")
		fmt.Println("Networks")
		for index, net := range nets {
			name, _ := net.GetName()
			fmt.Println(index, ". ", name)
		}
	}

	// Get Images
	images := GetImages()
	if err == nil {
		fmt.Println("--------------------------------")
		fmt.Println("Images")
		for index, img := range images {
			fmt.Println(index, ". ", img.Name)
		}
	}
}