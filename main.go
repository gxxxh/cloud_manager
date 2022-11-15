package main

import (
	"github.com/kube-stack/multicloud_service/src/service"
	"log"
)

func main() {
	mcm, err := service.NewMultiCloudService(nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(mcm)
}
