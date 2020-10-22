package mariadblayer

import (
	"cmpService/common/mcmodel"
	"fmt"
	"testing"
)

func TestDBORM_AddMcVmBackup(t *testing.T) {
	db, _ := getTestLkhDb()
	backup := mcmodel.McVmBackup{
		CompanyIdx:      191,
		McServerIdx:     19,
		McServerSn:      "SN87",
		VmName:          "SN87-VM-01",
		Desc:            "",
		Year:            0,
		Month:           0,
		Day:             1,
		Hour:            0,
		Minute:          0,
		Second:          0,
	}
	_, err := db.AddMcVmBackup(backup)
	fmt.Println("ERR: ", err)
}

