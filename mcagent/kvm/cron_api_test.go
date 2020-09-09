package kvm

import (
	"sync"
	"testing"
)

func TestAddSnapshotCron(t *testing.T) {
	AddSnapshotCron("00", "15")
}

func TestAddSnapshotCronByMin(t *testing.T) {
	AddSnapshotCronByMin("1")
}

func TestAddSnapshotCronBySecond(t *testing.T) {
	AddSnapshotCronBySecond("5")
}

func TestStartCron(t *testing.T) {
	var wg sync.WaitGroup
	n := NewCronSnapshot(5)
	SetCronSnapshot(n)

	wg.Add(1)
	AddSnapshotCronSecond("10", "vm01")
	go n.Start(&wg)
	wg.Wait()
}