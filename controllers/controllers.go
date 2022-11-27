package controllers

import (
	"encoding/json"
	"go_native/models"
	"log"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	message := models.Message{
		Message: "Hello World !",
	}
	jsonResp, err := json.Marshal(message)
	if err != nil {
		response.WriteHeader(http.StatusBadGateway)
		log.Println("[HELLOWORLD]", request.Method, http.StatusBadGateway, request.URL.Path)
	} else {
		response.WriteHeader(200)
		response.Write(jsonResp)
		log.Println(request.Method, request.URL.Path, http.StatusOK)
	}
}

func EmptyPath(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	message := models.Message{
		Message: "NO",
	}

	jsonResp, err := json.Marshal(message)

	if err != nil {
		log.Println("[HELLOWORLD]", request.Method, request.URL.Path, http.StatusBadGateway)
	} else {
		response.WriteHeader(http.StatusAccepted)
		response.Write(jsonResp)
		log.Println(request.Method, request.URL.Path, http.StatusOK)
	}
}
