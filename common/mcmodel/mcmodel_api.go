package mcmodel

func LookupImage(list *[]McImages, target McImages) *McImages {
	if list == nil {
		return nil
	}
	for _, obj := range *list {
		if obj.Name == target.Name {
			return &obj
		}
	}
	return nil
}

func LookupNetwork(list *[]McNetworks, target McNetworks) *McNetworks {
	if list == nil {
		return nil
	}
	for _, obj := range *list {
		if obj.Name == target.Name {
			return &obj
		}
	}
	return nil
}

func LookupVm(list *[]McVm, target McVm) *McVm {
	if list == nil {
		return nil
	}
	for _, obj := range *list {
		if obj.Name == target.Name {
			return &obj
		}
	}
	return nil
}

func LookupMgoImage(list *[]MgoImage, target MgoImage) *MgoImage {
	if list == nil {
		return nil
	}
	for _, obj := range *list {
		if obj.Name == target.Name {
			return &obj
		}
	}
	return nil
}

func LookupMgoNetwork(list *[]MgoNetwork, target MgoNetwork) *MgoNetwork {
	if list == nil {
		return nil
	}
	for _, obj := range *list {
		if obj.Name == target.Name {
			return &obj
		}
	}
	return nil
}

func LookupMgoVm(list *[]MgoVm, target MgoVm) *MgoVm {
	if list == nil {
		return nil
	}
	for _, obj := range *list {
		if obj.Name == target.Name {
			return &obj
		}
	}
	return nil
}

func (o *MgoVm) Compare(n *MgoVm) bool {
	if o.Name != n.Name {
		return true
	}
	if o.Cpu != n.Cpu {
		return true
	}
	if o.Ram != n.Ram {
		return true
	}
	if o.Hdd != n.Hdd {
		return true
	}
	if o.Desc != n.Desc {
		return true
	}
	if o.OS != n.OS {
		return true
	}
	if o.Image != n.Image {
		return true
	}
	if o.Filename != n.Filename {
		return true
	}
	if o.VmIndex != n.VmIndex {
		return true
	}
	if o.FullPath != n.FullPath {
		return true
	}
	if o.IpAddr != n.IpAddr {
		n.IsChangeIpAddr = true
		return true
	}
	if o.Mac != n.Mac {
		return true
	}
	if o.CurrentStatus != n.CurrentStatus {
		return true
	}
	if o.RemoteAddr != n.RemoteAddr {
		return true
	}
	return false
}

func (v MgoNetwork) Compare(n MgoNetwork) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Bridge != n.Bridge {
		return true
	}
	if v.Mode != n.Mode {
		return true
	}
	if v.Mac != n.Mac {
		return true
	}
	if v.DhcpStart != n.DhcpStart {
		return true
	}
	if v.DhcpEnd != n.DhcpEnd {
		return true
	}
	if v.Ip != n.Ip {
		return true
	}
	if v.Netmask != n.Netmask {
		return true
	}
	if v.Prefix != n.Prefix {
		return true
	}
	return false
}

func (v MgoImage) Compare(n MgoImage) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Variant != n.Variant {
		return true
	}
	if v.Hdd != n.Hdd {
		return true
	}
	if v.Desc != n.Desc {
		return true
	}
	if v.FullName != n.FullName {
		return true
	}
	return false
}

func (o *MgoServer) Compare(n *MgoServer) bool {
	isChanged := false

	if o.Vms != nil {
		if n.Vms == nil {
			isChanged = true
		} else if len(*(o.Vms)) != len(*(n.Vms)) {
			isChanged = true
		} else {
			for _, obj1 := range *o.Vms {
				obj2 := LookupMgoVm(n.Vms, obj1)
				if obj2 == nil {
					isChanged = true
				} else {
					res := obj1.Compare(obj2)
					if res == true {
						isChanged = true
					}
				}
			}
		}
	} else {
		if n.Vms != nil {
			isChanged = true
		}
	}

	if o.Networks != nil {
		if n.Networks == nil {
			isChanged = true
		} else if len(*(o.Networks)) != len(*(n.Networks)) {
			isChanged = true
		} else {
			for _, obj1 := range *o.Networks {
				obj2 := LookupMgoNetwork(n.Networks, obj1)
				if obj2 == nil {
					isChanged = true
				} else {
					res := obj1.Compare(*obj2)
					if res == true {
						isChanged = true
					}
				}
			}
		}
	} else {
		if n.Networks != nil {
			isChanged = true
		}
	}

	if o.Images != nil {
		if n.Images == nil {
			isChanged = true
		} else if len(*(o.Images)) != len(*(n.Images)) {
			isChanged = true
		} else {
			for _, obj1 := range *o.Images {
				obj2 := LookupMgoImage(n.Images, obj1)
				if obj2 == nil {
					isChanged = true
				} else {
					res := obj1.Compare(*obj2)
					if res == true {
						isChanged = true
					}
				}
			}
		}
	} else {
		if n.Images != nil {
			isChanged = true
		}
	}

	return isChanged
}

//----------------------------------------------------------------------

func (v McVm) Compare(n McVm) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Cpu != n.Cpu {
		return true
	}
	if v.Ram != n.Ram {
		return true
	}
	if v.Hdd != n.Hdd {
		return true
	}
	if v.Desc != n.Desc {
		return true
	}
	if v.OS != n.OS {
		return true
	}
	if v.Image != n.Image {
		return true
	}
	if v.Filename != n.Filename {
		return true
	}
	if v.VmIndex != n.VmIndex {
		return true
	}
	if v.FullPath != n.FullPath {
		return true
	}
	if v.IpAddr != n.IpAddr {
		return true
	}
	if v.Mac != n.Mac {
		return true
	}
	if v.CurrentStatus != n.CurrentStatus {
		return true
	}
	if v.RemoteAddr != n.RemoteAddr {
		return true
	}
	return false
}

func (v McNetworks) Compare(n McNetworks) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Bridge != n.Bridge {
		return true
	}
	if v.Mode != n.Mode {
		return true
	}
	if v.Mac != n.Mac {
		return true
	}
	if v.DhcpStart != n.DhcpStart {
		return true
	}
	if v.DhcpEnd != n.DhcpEnd {
		return true
	}
	if v.Ip != n.Ip {
		return true
	}
	if v.Netmask != n.Netmask {
		return true
	}
	if v.Prefix != n.Prefix {
		return true
	}
	return false
}

func (v McImages) Compare(n McImages) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Variant != n.Variant {
		return true
	}
	if v.Hdd != n.Hdd {
		return true
	}
	if v.Desc != n.Desc {
		return true
	}
	if v.FullName != n.FullName {
		return true
	}
	return false
}

func (s *McServerMsg) Compare(n *McServerMsg) bool {
	isChanged := false

	if s.Vms != nil {
		if n.Vms == nil {
			return true
		}
		if len(*(s.Vms)) != len(*(n.Vms)) {
			return true
		}
		for _, obj1 := range *s.Vms {
			obj2 := LookupVm(n.Vms, obj1)
			if obj2 == nil {
				return true
			}
			res := obj1.Compare(*obj2)
			if res == true {
				return res
			}
		}
	} else {
		if n.Vms != nil {
			return true
		}
	}

	if s.Networks != nil {
		if n.Networks == nil {
			return true
		}
		if len(*(s.Networks)) != len(*(n.Networks)) {
			return true
		}
		for _, obj1 := range *s.Networks {
			obj2 := LookupNetwork(n.Networks, obj1)
			if obj2 == nil {
				return true
			}
			res := obj1.Compare(*obj2)
			if res == true {
				return res
			}
		}
	} else {
		if n.Networks != nil {
			return true
		}
	}

	if s.Images != nil {
		if n.Images == nil {
			return true
		}
		if len(*(s.Images)) != len(*(n.Images)) {
			return true
		}
		for _, obj1 := range *s.Images {
			obj2 := LookupImage(n.Images, obj1)
			if obj2 == nil {
				return true
			}
			res := obj1.Compare(*obj2)
			if res == true {
				return res
			}
		}
	} else {
		if n.Images != nil {
			return true
		}
	}

	return isChanged
}
