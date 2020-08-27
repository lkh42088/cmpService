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

