package kvm

import "testing"

var vmName2 = "vm01"

func TestShutdownVm2(t *testing.T) {
	LibvirtShutdownVm(vmName2)
}

func TestKvmSuspendVm(t *testing.T) {
	LibvirtSuspendVm(vmName2)
}

func TestKvmResumeVm(t *testing.T) {
	LibvirtResumeVm(vmName2)
}

func TestKvmDestroyVm(t *testing.T) {
	LibvirtDestroyVm(vmName2)
}

func TestKvmUndefineVm(t *testing.T) {
	LibvirtUndefineVm(vmName2)
}

func TestKvmStartVm(t *testing.T) {
	LibvirtStartVm(vmName2)
}

func TestGetKvmVmState(t *testing.T) {
	GetLibvirtVmState(vmName2)
}
