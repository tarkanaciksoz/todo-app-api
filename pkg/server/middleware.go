package server

import (
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Z-Key, Access-Control-Allow-Headers, Access-Control-Request-Method, Access-Control-Request-Headers")
		rw.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, PATCH, DELETE")
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
		rw.WriteHeader(http.StatusOK)

		ctx := r.Context()
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func OptionsHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Z-Key, Access-Control-Allow-Headers, Access-Control-Request-Method, Access-Control-Request-Headers")
	rw.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, PATCH, DELETE")
	rw.Header().Add("Access-Control-Allow-Credentials", "true")
	rw.WriteHeader(http.StatusOK)
}
