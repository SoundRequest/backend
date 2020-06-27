package checker

import (
	"net/http"

	"gopkg.in/oauth2.v3/server"
)

// CheckIsTokenValid success=200, failed=400
func CheckIsTokenValid(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		f.ServeHTTP(w, r)
	})
}
