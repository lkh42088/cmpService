package mysqllayer

import "nubes/dbmigrator/cbmodels"

type CBDBLayer interface {
	// Item
	GetAllItems() ([]cbmodels.Item, error)
	// SubItem
	GetAllSubItems() ([]cbmodels.SubItem, error)
	// Devices
	GetAllDevices() ([]cbmodels.CbDevice, error)
}