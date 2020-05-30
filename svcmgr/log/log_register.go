package log

import (
	"cmpService/common/lib"
	"cmpService/common/mariadblayer"
	"cmpService/common/models"
	"fmt"
	"time"
)

type Handler struct {
	db mariadblayer.MariaDBLayer
}

func NewHandler(db *mariadblayer.DBORM) (*Handler, error) {
	h := new(Handler)
	h.db = db
	return h, nil
}

func (h *Handler) AutoAddLog(log models.DeviceLog) {
	if h.db == nil {
		return
	}
	if log.DeviceCode == "" || log.WorkCode == 0 {
		return
	}
	if err := h.db.AddLog(log); err != nil {
		lib.LogWarn("[AutoAddLog] %s\n", err)
		return
	}
	lib.LogInfo("[AutoAddLog] Log message stored(workCode=%d).\n", log.WorkCode)
}

func RegisterDeviceLog(code string) error {
	log := models.DeviceLog{
		DeviceCode:   code,
		WorkCode:     lib.RegisterDevice,
		LogLevel:     lib.LevelInfo,
		RegisterId: "",   //todo
		RegisterName: "",    //todo
		RegisterDate: time.Now(),
	}
	fmt.Println(log)
	//err := AutoAddLog(log)
	//if err != nil {
	//	return errors.New("[RegisterDeviceLog] Failed to insert log message in DB")
	//}

	return nil
}