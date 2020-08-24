package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFileListInFolder(dir string) (files []string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	return files
}

func GetQcowFileInFolder(dir string) (images []string) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		//fmt.Println(file)
		if strings.HasSuffix(file, ".qcow2") &&
			strings.Contains(file, "G") {
			images = append(images, file)
		}
	}
	return images
}

func IsExistFile(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}

