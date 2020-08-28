package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) AddMcNetwork(obj mcmodel.McNetworks) (mcmodel.McNetworks, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) UpdateMcNetwork(obj mcmodel.McNetworks) (mcmodel.McNetworks, error) {
	return obj, db.Model(&obj).
		Updates(map[string]interface{}{
			"net_Mode":obj.Mode,
			"net_mac":obj.Mac,
			"net_dhcp_start":obj.DhcpStart,
			"net_dhcp_end":obj.DhcpEnd,
			"net_ip":obj.Ip,
			"net_netmask":obj.Netmask,
			"net_prefix":obj.Prefix,
		}).Error
}

func (db *DBORM) GetMcNetworksByServerIdx(serverIdx int) (obj []mcmodel.McNetworks, err error) {
	return obj, db.Table("mc_network_tb").
		Where(mcmodel.McNetworks{McServerIdx: serverIdx}).
		Find(&obj).Error
}

func (db *DBORM) DeleteMcNetwork(obj mcmodel.McNetworks) (mcmodel.McNetworks, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcNetworksPage(paging models.Pagination) (networks mcmodel.McNetworkPage, err error) {
	err = db.
		Table("mc_network_tb").
		Select("mc_network_tb.*, c.cp_name, m.mc_serial_number").
		Joins("LEFT JOIN mc_server_tb m ON m.mc_idx = mc_network_tb.net_server_idx").
		Joins("LEFT JOIN company_tb c ON c.cp_idx = m.mc_cp_idx").
		//Order(networks.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&networks.Networks).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(networks.Networks)
	networks.Page = paging
	return networks, err
}

