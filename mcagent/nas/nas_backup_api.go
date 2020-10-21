package nas

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func CopyBackupFileToNas(srcPath string, dstPath string) error {
	//cfg := config.GetMcGlobalConfig()
	//org := fmt.Sprintf("/home/nubes/go/src/a.txt")
	//target := fmt.Sprintf("/home/nubes/nas/backup/.")
	if srcPath == "" || dstPath == "" {
		return errors.New("Path parameter is invalid.\n")
	}
	if _, err := os.Stat(srcPath); err != nil {
		return errors.New("Source file is not exist.\n")
	}

	args := []string{
		srcPath,
		dstPath,
	}

	fmt.Println("args: ", args)

	binary := "cp"
	cmd := exec.Command(binary, args...)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}