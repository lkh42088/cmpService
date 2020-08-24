package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"fmt"
	"strconv"
	"strings"
)

func GetMgoImageByName (name string) mcmodel.MgoImage {
	var image mcmodel.MgoImage
	image.Name = name[:strings.LastIndexAny(name,".")]
	list := strings.Split(name, "-")
	fmt.Println(list)
	if list[0] == "windows10" {
		image.Variant = "win10"
	} else {
		image.Variant = "ubuntu18"
	}
	image.Hdd, _ = strconv.Atoi(list[1][:strings.LastIndexAny(list[1],"G")])
	//fmt.Println(image)
	return image
}

func InitImages() {
	cfg := config.GetGlobalConfig()
	images := utils.GetQcowFileInFolder(cfg.VmImageDir)
	for _, image := range images {
		img := GetMgoImageByName(image[len(cfg.VmImageDir)+1:])
		fmt.Printf("image: %v", img)
		//mcmongo.McMongo.AddImage(&img)
	}
}

func GetImages() (list []mcmodel.MgoImage) {
	cfg := config.GetGlobalConfig()
	images := utils.GetQcowFileInFolder(cfg.VmImageDir)
	for _, image := range images {
		img := GetMgoImageByName(image[len(cfg.VmImageDir)+1:])
		//fmt.Printf("image: %v\n", img)
		list = append(list, img)
	}
	return list
}
