package kvm

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type CronScheduler struct {
	Interval int
	Cr       *cron.Cron
	SnapVms  []SnapVm
	BackupVms []BackupVm
}

type SnapVm struct {
	VmName string
	CronId cron.EntryID
	Timer string
}

type BackupVm struct {
	VmName string
	CronId cron.EntryID
	Timer string
}

var vmSliceMutex sync.Mutex

func (c *CronScheduler) LookupBackupVm(vmName string) bool {
	for _, vm := range c.BackupVms {
		if vm.VmName == vmName {
			return true
		}
	}
	return false
}

func (c *CronScheduler) DeleteBackupVm(vmName string) bool {
	findIt := -1
	for index, vm := range c.BackupVms {
		if vm.VmName == vmName {
			DeleteCronEntry(vm.CronId)
			findIt = index
		}
	}
	if findIt > 0 {
		vmSliceMutex.Lock()
		c.BackupVms = append(c.BackupVms[:findIt], c.BackupVms[findIt+1:]...)
		vmSliceMutex.Unlock()
		return true
	}
	return false
}

func (c *CronScheduler) LookupSnapVm(vmName string) bool {
	for _, vm := range c.SnapVms {
		if vm.VmName == vmName {
			return true
		}
	}
	return false
}

func (c *CronScheduler) DeleteSnapVm(vmName string) bool {
	findIt := -1
	for index, vm := range c.SnapVms {
		if vm.VmName == vmName {
			DeleteCronEntry(vm.CronId)
			findIt = index
		}
	}
	if findIt > 0 {
		vmSliceMutex.Lock()
		c.SnapVms = append(c.SnapVms[:findIt], c.SnapVms[findIt+1:]...)
		vmSliceMutex.Unlock()
		return true
	}
	return false
}

func DeleteCronEntry(id cron.EntryID) {
	CronSch.Cr.Remove(id)
}
