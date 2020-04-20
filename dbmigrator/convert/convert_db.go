package convert

import (
	"fmt"
	"nubes/common/db"
	"nubes/common/mariadblayer"
	"nubes/common/models"
	"nubes/dbmigrator/cbmodels"
	"nubes/dbmigrator/config"
	"nubes/dbmigrator/mysqllayer"
)

func RunConvertDb() {
	// Old Database: Mysql
	oldConfig := config.GetOldDatabaseConfig()
	oldOptions := db.GetDataSourceName(oldConfig)
	oldDb, err := mysqllayer.NewCBORM(oldConfig.DBDriver, oldOptions)
	if err != nil {
		fmt.Println("oldConfig Error:", err)
		return
	}
	defer oldDb.Close()

	// New Database: Mariadb
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()

	ConvertItem(oldDb, newDb)
	ConvertItemSub(oldDb, newDb)
}

func ConvertItem(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllItems()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for num, old := range olds {
		new := GetCodeByItem(old)
		fmt.Println(num, ":", old, "-->", new)
		ndb.AddCode(new)
	}
}

func ConvertItemSub(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllSubItems()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for num, old := range olds {
		new := GetSubCodeByItemSub(old)
		fmt.Println(num, ":", old, "-->", new)
		ndb.AddSubCode(new)
	}
}

func GetCodeByItem(item cbmodels.Item) (code models.Code) {
	code.CodeID = item.ItemID
	code.Type = item.Table
	code.SubType = item.Column
	code.Name = item.Item
	code.Order = item.Desc
	return code
}

func GetSubCodeByItemSub(subitem cbmodels.SubItem) (subcode models.SubCode) {
	subcode.ID = subitem.SubItemID
	subcode.Name = subitem.SubItem
	subcode.CodeID = subitem.ItemID
	subcode.Order = subitem.Desc
	return subcode
}

