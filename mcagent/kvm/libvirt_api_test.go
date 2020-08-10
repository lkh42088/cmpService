package kvm

import (
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"testing"
)

type test struct {
	XMLName xml.Name `xml:"test"`
	Abc     abc      `xml:"abc"`
	Eee     string   `xml:"eee"`
}

type abc struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func TestXml(t *testing.T) {
	a := &abc{Key: "tester", Value: "aaaaa"}
	v := &test{Abc: *a, Eee: "eeee"}

	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

	var te test
	err = xml.Unmarshal(output, &te)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(te)
}

func TestLibvirtXml(t *testing.T) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	dom, err := conn.LookupDomainByName("win10-bhjung")
	if err != nil {
		fmt.Println("error2")
	}
	xmldoc, err := dom.GetXMLDesc(0)
	if err != nil {
		fmt.Println("error3")
	}

	domcfg := &libvirtxml.Domain{}
	err = domcfg.Unmarshal(xmldoc)
	if err != nil {
		fmt.Println("error4")
	}

	output, _ := xml.Marshal(domcfg)
	fmt.Printf("Virt type %s\n", domcfg.Type)
	fmt.Printf("Virt %s\n", string(output))
	x := xmlfmt.FormatXML(string(output), "\t", "  ")
	print(x)
}

func TestGetXmlDomain(t *testing.T) {
	GetXmlDomain("win10-bhjung")
	//domcfg := GetXmlDomain("win10-bhjung")
	//output, _ := xml.Marshal(domcfg)
	//DumpXml(string(output))
}

func TestGetAllNetwork(t *testing.T) {
	GetAllNetwork()
}

func TestGetDomain(t *testing.T) {
	GetDomain()
}

func TestGetXmlNetwork(t *testing.T) {
	GetXmlNetwork()
}
