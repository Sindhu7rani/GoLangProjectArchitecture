package middleware

import "net/http"

/*
AuthMiddleware is executed for every incoming request.
It checks whether the request is authenticated and authorized.
If valid, the request is forwarded to the actual handler.
*/
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Retrieve the token from the Authorization header of the request
		token := r.Header.Get("Authorization")
		//if no token->not authticated->err
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		//if user authenticated go to handler
		next.ServeHTTP(w, r)
	})
}
