package main

import "github.com/kube-stack/multicloud_service/src/generator/cmd"

func main() {
	cmd.Init()
	cmd.Execute()
	//mcm, err := service.NewMultiCloudService(nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(mcm)
}
