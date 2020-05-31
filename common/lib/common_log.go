package lib

const (
	_				= iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

const (
	_					= iota
	RegisterDevice
	ChangeInformation
	ExportDevice
	MovedDevice
)

const (
	RestFailGetData				= "Failed to get data."
	RestFailAddData				= "Failed to add data."
	RestFailDeleteData			= "Failed to delete data"
	RestAbnormalParam			= "Message parameter abnormal."
	RestFailedUserId			= "Can't find user-id."
	RestFailFindData			= "Don't find wanted data."
	RestDoNotCreateUser			= "Can to action only create user."
	RestFailConvertData			= "Can't convert from raw data."
)
