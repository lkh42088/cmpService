package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
)

func (db *DBORM) AddSubnet(subnet models.SubnetMgmt) error {
	return db.Create(&subnet).Error
}

func (db *DBORM) GetSubnets(param models.PageRequestForSearch) (
	subnet models.SubnetMgmtResponse, err error) {

	var searchParam string
	if param.SearchParam != "" {
		searchParam = "sub_ip_start like '%" + param.SearchParam + "%' or " +
			"sub_ip_end like '%" + param.SearchParam + "%' or " +
			"sub_tag like '%" + param.SearchParam + "%' or " +
			"sub_mask like '%" + param.SearchParam + "%' or " +
			"sub_gateway like '%" + param.SearchParam + "%'"
	} else {
		searchParam = ""
	}

	db.Model(&subnet.Subnet).Where(searchParam).Count(&subnet.Page.TotalCount)

	err = db.
		//Debug().
		Order(subnet.GetOrderBy(param.OrderBy, param.Order)).
		Limit(param.RowsPerPage).
		Offset(param.Offset).
		Where(searchParam).
		Find(&subnet.Subnet).Error

	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	subnet.Page.RowsPerPage = param.RowsPerPage
	subnet.Page.Offset = param.Offset
	subnet.Page.Order = param.Order
	subnet.Page.OrderBy = param.OrderBy

	return subnet, err
}

func (db *DBORM) UpdateSubnet(subnet models.SubnetMgmt) error {
	fmt.Printf("%+v", subnet)
	return db.
		Model(models.SubnetMgmt{}).
		Where(models.SubnetMgmt{Idx: subnet.Idx}).
		Update(&subnet).
		Error
}

func (db *DBORM) DeleteSubnets(idx []string) error {
	return db.Debug().Where("sub_idx in (?)", idx).Delete(&models.SubnetMgmt{}).Error
}
