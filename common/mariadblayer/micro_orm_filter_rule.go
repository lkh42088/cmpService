package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) GetMcFilterRule() (obj []mcmodel.McFilterRule, err error) {
	return obj, db.Table("mc_filter_rule_tb").
		Select("mc_filter_rule_tb.*").
		Find(&obj).Error
}

func (db *DBORM) GetMcFilterRuleBySerialNumberAndAddr(sn, addr string) (obj []mcmodel.McFilterRule, err error) {
	return obj, db.Table("mc_filter_rule_tb").
		Select("mc_filter_rule_tb.*").
		Where(mcmodel.McFilterRule{
			SerialNumber: sn,
			IpAddr: addr,
		}).
		Find(&obj).Error
}

func (db *DBORM) GetMcFilterRuleBySerialNumber(sn string) (obj []mcmodel.McFilterRule, err error) {
	return obj, db.Table("mc_filter_rule_tb").
		Select("mc_filter_rule_tb.*").
		Where(mcmodel.McFilterRule{SerialNumber: sn}).
		Find(&obj).Error
}

func (db *DBORM) GetMcFilterRulePage(paging models.Pagination, cpName string) (
	obj mcmodel.McFilterRulePage, err error) {
	var query string
	if cpName == "all" {
		query = ""
	} else {
		query = "c.cp_name = '" + cpName + "'"
	}
	err = db.Debug().
		Table("mc_filter_rule_tb").
		Select("mc_filter_rule_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_filter_rule_tb.mc_cp_idx").
		Order(obj.GetOrderBy(paging.OrderBy, paging.Order)).
		/*Limit(paging.RowsPerPage).*/
		Offset(paging.Offset).
		//Where("c.cp_name = ?", cpName).
		Where(query).
		Find(&obj.FilterRules).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(obj.FilterRules)
	obj.Page = paging

	return obj, err
}

func (db *DBORM) GetMcFilterRuleByServerIdx(idx uint) (obj mcmodel.McFilterRuleDetail, err error) {
	err = db.
		Table("mc_filter_rule_tb").
		Select("mc_filter_rule_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_filter_rule_tb.mc_cp_idx").
		Where(mcmodel.McFilterRule{Idx: idx}).
		Find(&obj).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	return obj, err
}

func (db *DBORM) AddMcFilterRule(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteMcFilterRule(obj mcmodel.McFilterRule) (mcmodel.McFilterRule, error) {
	return obj, db.Delete(&obj).Error
}
