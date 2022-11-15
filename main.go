package main

import (
	"log"
	"multicloud_service/src/service"
)

func main() {
	mcm, err := service.NewMultiCloudService(nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(mcm)
}
