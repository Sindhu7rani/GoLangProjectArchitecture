package router

import (
	"net/http"

	"GoProjectArche/internal/handler"
	"GoProjectArche/internal/middleware"
)

/*
SetupRoutes registers all API endpoints.
When a request comes in, the router:
1. Matches the URL path
2. Executes middleware (security checks)
3. Calls the corresponding handler
*/
func SetupRoutes(h *handler.AccountHandler) http.Handler {
	// Create a new custom ServeMux (router)
	// ServeMux is responsible for routing incoming requests
	mux := http.NewServeMux()

	//depends on request middlerware(security check)+handler works
	mux.Handle("/accounts/create",
		middleware.AuthMiddleware(http.HandlerFunc(h.CreateAccount)),
	)

	mux.Handle("/accounts/get",
		middleware.AuthMiddleware(http.HandlerFunc(h.GetAccount)),
	)
	return mux
}
