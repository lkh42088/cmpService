package nas

import "testing"

func TestCopyBackupFileToNas(t *testing.T) {
	src := "/home/nubes/go/src/a.txt"
	dst := "/home/nubes/nas/backup/."
	CopyBackupFileToNas(src, dst)
}