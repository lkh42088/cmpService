package cron

import (
	"sync"
	"testing"
)

func TestAddSnapshotCron(t *testing.T) {
	AddSnapshotCron("00", "15")
}

func TestAddSnapshotCronByMin(t *testing.T) {
	AddSnapshotCronByMin("vm01", "1")
}

func TestAddSnapshotCronBySecond(t *testing.T) {
	AddSnapshotCronBySecond("5")
}

func TestStartCron(t *testing.T) {
	var wg sync.WaitGroup
	n := NewCronScheduler(5)
	SetCronScheduler(n)

	wg.Add(1)
	AddSnapshotCronSecond("10", "vm01")
	go n.Start(&wg)
	wg.Wait()
}