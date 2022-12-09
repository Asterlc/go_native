package controllers

import (
	"encoding/json"
	"go_native/models"
	"log"
	"net/http"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	message := models.Message{
		Message: "Hello World !",
	}
	jsonResp, err := json.Marshal(message)
	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		log.Println("[HELLOWORLD]", req.Method, http.StatusBadGateway, req.URL.Path)
	} else {
		res.WriteHeader(200)
		res.Write(jsonResp)
		log.Println(req.Method, req.URL.Path, http.StatusOK)
	}
}

func EmptyPath(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("Content-type", "application/json")
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
