package mysqllayer

import "cmpService/dbmigrator/cbmodels"

func (db *CBORM) GetAllMemberFromOldDB() (members []cbmodels.CbMember, err error) {
	return members, db.Find(&members).Error
}


