package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) AddMcImage(obj mcmodel.McImages) (mcmodel.McImages, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) GetMcImagesByServerIdx(serverIdx int) (obj []mcmodel.McImages, err error) {
	return obj, db.Table("mc_image_tb").
		Where(mcmodel.McImages{McServerIdx: serverIdx}).
		Find(&obj).Error
}

func (db *DBORM) DeleteMcImage(obj mcmodel.McImages) (mcmodel.McImages, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcImagesPage(paging models.Pagination) (images mcmodel.McImagePage, err error) {
	err = db.
		Table("mc_image_tb").
		Select("mc_image_tb.*, c.cp_name, m.mc_serial_number").
		Joins("LEFT JOIN mc_server_tb m ON m.mc_idx = mc_image_tb.img_server_idx").
		Joins("LEFT JOIN company_tb c ON c.cp_idx = m.mc_cp_idx").
		//Order(images.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&images.Images).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(images.Images)
	images.Page = paging
	return images, err
}

