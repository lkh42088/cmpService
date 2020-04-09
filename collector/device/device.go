package device

import (
	"fmt"
)

type Device struct {
	Id ID `json:"id,omitempty"`
	Ip string `json:"ip,omitempty"`
	Port int `json:"port,omitempty"`
	SnmpCommunity string `json:"community,omitempty"`
}

func (d Device) String() string {
	return fmt.Sprintf("{ id: %s, ip: %s, port: %d, community: %s }",
		d.GetIdString(), d.Ip, d.Port, d.SnmpCommunity)
}

func (d *Device) GetIdString() ID {
	//return CodeID(fmt.Sprintf("%x", d.Id))
	return ID(d.Id)
}
