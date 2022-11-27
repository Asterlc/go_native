package controllers

import (
	"encoding/json"
	"go_native/models"
	"log"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("Content-type", "application/json")
	message := models.Message{
		Message: "Hello World !",
	}

	// json.NewEncoder(response).Encode(message)

	jsonResp, err := json.Marshal(message)
	if err != nil {
		response.WriteHeader(http.StatusBadGateway)
		log.Println("[HELLOWORLD]", request.Method, http.StatusBadGateway, request.URL.Path)
	} else {
		response.WriteHeader(200)
		response.Write(jsonResp)
		log.Println(request.Method, http.StatusOK, request.URL.Path)
	}

	// error := json.NewEncoder(response).Encode(message)
	// if error != nil {
	// 	response.WriteHeader(400)
	// 	log.Println("Fait to hello world", error.Error(), "status: 400")
	// } else {
	// 	response.WriteHeader(200)
	// 	json.NewEncoder(response).Encode(message)
	// 	log.Println(request.Method, request.URL.Path)
	// }
}

func EmptyPath() {

}
