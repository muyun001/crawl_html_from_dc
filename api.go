package main

import (
	"crawl_html_from_dc/main/routers"
	"log"
)

func main() {
	router := routers.Load()

	err := router.Run(":9010")
	if err != nil {
		log.Fatalln(err)
	}
}
