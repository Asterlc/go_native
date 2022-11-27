package routes

import (
	"go_native/controllers"
	"go_native/middleware"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", middleware.ContentJSON(controllers.EmptyPath))
	http.HandleFunc("/hello", middleware.ContentJSON(controllers.Hello))
}
