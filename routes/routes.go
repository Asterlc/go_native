package routes

import (
	"go_native/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controllers.EmptyPath)
	http.HandleFunc("/hello", controllers.Hello)
}
