package collectdevice

import (
	"fmt"
)

type ColletDevice struct {
	Id ID `json:"id,omitempty"`
	Ip string `json:"ip,omitempty"`
	Port int `json:"port,omitempty"`
	SnmpCommunity string `json:"community,omitempty"`
}

func (d ColletDevice) String() string {
	return fmt.Sprintf("{ id: %s, ip: %s, port: %d, community: %s }",
		d.GetIdString(), d.Ip, d.Port, d.SnmpCommunity)
}

func (d *ColletDevice) GetIdString() ID {
	//return CodeID(fmt.Sprintf("%x", d.Id))
	return ID(d.Id)
}
