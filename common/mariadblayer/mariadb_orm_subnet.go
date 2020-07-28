package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
)

func (db *DBORM) AddSubnet(subnet models.SubnetMgmt) error {
	return db.Create(&subnet).Error
}

func (db *DBORM) GetSubnets(page models.Pagination) (subnet models.SubnetMgmtResponse, err error) {
	db.Model(&subnet.Subnet).Count(&page.TotalCount)
	err = db.
		Order(subnet.GetOrderBy(page.OrderBy, page.Order)).
		Limit(page.RowsPerPage).
		Offset(page.Offset).
		Find(&subnet.Subnet).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	subnet.Page = page
	return subnet, err
}

func (db *DBORM) DeleteSubnets(idx []string) error {
	return db.Debug().Where("sub_idx in (?)", idx).Delete(&models.SubnetMgmt{}).Error
}
