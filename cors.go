package helper_GO

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CORS_EnableCors(r *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	return c.Handler(r)
}
