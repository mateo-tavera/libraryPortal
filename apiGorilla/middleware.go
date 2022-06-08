package apigorilla

import (
	"net/http"

	"github.com/google/uuid"
)

//Generates a Header in every request
func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestId := r.Header.Get("X-Request-ID") //Create header id
		if len(requestId) == 0 {
			requestId = uuid.New().String()
		}

		w.Header().Set("X-Request-ID", requestId)

		next.ServeHTTP(w, r) //When finish this function, we proced to the next one
	})

}

//Define validUsers for endpoints
var validUsers = map[string]string{"user1": "password1", "user2": "password2"}

//Control authorization for certian endpoints
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := r.Header.Get("Authorization")

		//Check if user is in the list of valid users
		if validUsers[user] == "" {
			w.WriteHeader(http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r) //When finish this function, we proced to the next one

	})
}
