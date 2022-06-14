package apigorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {

	public := r.NewRoute().Subrouter()
	protected := r.NewRoute().Subrouter()

	r.Use(RequestIdMiddleware)             //It means that any router uses this middleware
	protected.Use(AuthorizationMiddleware) //Only the routes 'protected' uses this middleware

	public.HandleFunc("/books", a.GetBooks).Methods(http.MethodGet)
	public.HandleFunc("/book/{id}", a.GetBook).Methods(http.MethodGet)

	protected.HandleFunc("/book/{id}", a.PatchBook).Methods(http.MethodPatch)
	protected.HandleFunc("/books", a.CreateBook).Methods(http.MethodPost)

}
