package device

import (
	"fmt"
)

type InMemoryAccessor struct {
	devices map[ID]Device
	nextID int64
}

func NewMemoryDataAccess() Accessor {
	return &InMemoryAccessor{
		devices: map[ID]Device{},
		nextID:  1,
	}
}

func (m *InMemoryAccessor) Get(id ID) (Device, error) {
	d, exists := m.devices[id]
	if !exists {
		return Device{}, ErrDeviceNotExist
	}
	return d, nil
}

func (m *InMemoryAccessor) Put(id ID, d Device) error {
	if _, exists := m.devices[id]; !exists {
		return ErrDeviceNotExist
	}
	m.devices[id] = d
	return nil
}

func (m *InMemoryAccessor) Post(d Device) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.devices[id] = d
	return id, nil
}

func (m *InMemoryAccessor) Delete(id ID) error {
	if _, exists := m.devices[id]; !exists {
		return ErrDeviceNotExist
	}
	delete(m.devices, id)
	return nil
}

