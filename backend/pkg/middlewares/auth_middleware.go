package middlewares

import (
	"github.com/gorilla/mux"
	"net/http"
)

// AuthMiddleware is an Authorization check middleware
func AuthMiddleware(r *mux.Router) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			auth := getRequestHeaderAuth(r, req)
			if auth == "" {
				_, _ = w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}

func getRequestHeaderAuth(r *mux.Router, req *http.Request) string {
	reqAuth := req.Header.Get("Authorization")

	return reqAuth
}
