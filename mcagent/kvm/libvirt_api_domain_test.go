package kvm

import "testing"

var vmName2 = "vm03"

func TestShutdownVm2(t *testing.T) {
	KvmShutdownVm(vmName2)
}

func TestKvmSuspendVm(t *testing.T) {
	KvmSuspendVm(vmName2)
}

func TestKvmResumeVm(t *testing.T) {
	KvmResumeVm(vmName2)
}

func TestKvmDestroyVm(t *testing.T) {
	KvmDestroyVm(vmName2)
}

func TestKvmUndefineVm(t *testing.T) {
	KvmUndefineVm(vmName2)
}

func TestKvmStartVm(t *testing.T) {
	KvmStartVm(vmName2)
}

func TestGetKvmVmState(t *testing.T) {
	GetKvmVmState(vmName2)
}
