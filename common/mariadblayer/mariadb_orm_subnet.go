package mariadblayer

import "cmpService/common/models"

func (db *DBORM) AddSubnet(subnet models.SubnetMgmt) error {
	return db.Create(&subnet).Error
}

