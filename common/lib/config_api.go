package lib

import (
	"bytes"
	"encoding/json"
	"os"
)

func CreateConfig(path string, cfg interface{}) error {
	var file * os.File
	var err error
	var b []byte

	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	if err != nil {
		LogWarnln("Failed to create collector default config file", err)
		return err
	}
	defer file.Close()

	// Get default config
	b, err = json.Marshal(cfg)
	if err != nil {
		return err
	}

	b, err = PrettyPrint(b)
	if err != nil {
		return err
	}

	// write file
	_, err = file.WriteString(string(b))
	if err != nil {
		return err
	}

	return nil
}

func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
