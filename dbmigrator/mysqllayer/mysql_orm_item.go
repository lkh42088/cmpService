package mysqllayer

import "cmpService/dbmigrator/cbmodels"

func (db *CBORM) GetAllItems() (items []cbmodels.Item, err error) {
	return items, db.Find(&items).Error
}

func (db *CBORM) GetAllSubItems() (subitems []cbmodels.SubItem, err error) {
	return subitems, db.Find(&subitems).Error
}

func (db *CBORM) GetAllDevices() (devices []cbmodels.CbDevice, err error) {
	return devices, db.Find(&devices).Error
}

func (db *CBORM) GetAllDevicesServerFromOldDB() (devices []cbmodels.ServerDevice, err error) {
	return devices, db.Find(&devices).Error
}

func (db *CBORM) GetAllDevicesNetworkFromOldDB() (devices []cbmodels.NetworkDevice, err error) {
	return devices, db.Find(&devices).Error
}

func (db *CBORM) GetAllDevicesPartFromOldDB() (devices []cbmodels.PartDevice, err error) {
	return devices, db.Find(&devices).Error
}
