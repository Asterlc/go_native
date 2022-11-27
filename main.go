package main

import (
	"go_native/routes"
	"log"
	"net/http"
)

func main() {
	routes.Routes()
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Panic(err.Error())
	} else {
		log.Println("Started")
	}
}
