package routes

import (
	"go_native/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/hello", controllers.Hello)
}
