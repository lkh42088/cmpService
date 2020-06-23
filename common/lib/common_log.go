package lib

const (
	_ = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

const (
	_ = iota
	RegisterDevice
	ChangeInformation
	ImportDevice
	ExportDevice
	MovedDevice
)

const (
	RestFailGetData      = "Failed to get data.\n"
	RestFailAddData      = "Failed to add data.\n"
	RestFailDeleteData   = "Failed to delete data.\n"
	RestFailDbConnection = "Failed to connect DB.\n"
	RestAbnormalParam    = "Message parameter abnormal.\n"
	RestFailedUserId     = "Can't find user-id.\n"
	RestFailFindData     = "Don't find wanted data.\n"
	RestDoNotCreateUser  = "Can to action only create user.\n"
	RestFailConvertData  = "Can't convert from raw data.\n"

	LogFailAddRegister = "Can't add a log.\n"
)
