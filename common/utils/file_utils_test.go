package utils

import (
	"fmt"
	"testing"
)

func TestGetFileListInFolder(t *testing.T) {
	files := GetFileListInFolder("/opt/images")
	fmt.Println(files)
}

func TestGetQcowFileInFolder(t *testing.T) {
	files := GetQcowFileInFolder("/opt/images")
	fmt.Println("files:", files)
}
