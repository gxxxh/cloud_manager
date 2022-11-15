package main

import (
	"cloud_manager/src/service"
	"log"
)

func main() {
	mcm, err := service.NewMultiCloudService(nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(mcm)
}
