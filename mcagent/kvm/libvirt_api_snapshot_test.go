package kvm

import (
	"fmt"
	"testing"
	"time"
)

var vmName = "vm10"

func TestCreateSnapshot(t *testing.T) {
	CreateSnapshot(vmName, "snap 02", "test snap 02")
}

func TestGetSnapshot(t *testing.T) {
	GetAllSnapshots(vmName)
}

func TestGetSnapshotListName(t *testing.T) {
	GetSnapshotsListName(vmName)
}

func TestDeleteAllSnapshot(t *testing.T) {
	DeleteAllSnapshot(vmName)
}

func TestTimeDate(a *testing.T) {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%d%s%d-%d-%d-%d\n",
		t.Year(), t.Month().String()[:3], t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func TestSlice(t *testing.T) {
	a := []string{"h1", "h2", "h3"}
	fmt.Println(a)
	fmt.Println(a[:0])
	fmt.Println(a[1:])
	a = append(a[:1], a[2:]...)
	fmt.Println(a)

	findit := -1
	b := []string{"h1", "h2", "h3"}
	for index, en := range b {
		if en == "h2" {
			findit = index
			break
		}
	}
	if findit > 0 {
		b = append(b[:findit], b[findit+1:]...)
	}
	fmt.Println(b)
}
