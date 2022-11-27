package middleware

import "net/http"

// EXEMPLO DO GORILLA MUX
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// TENTANDO FAZER NATIVO
func ContentJSON(controller http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		controller.ServeHTTP(w, r)
	}
}
