package router

import (
	controller "entdemo/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerOrderRouter(r *mux.Router) {
	orderRouter := r.PathPrefix("api/v1/orders").Subrouter()
	orderRouter.HandleFunc("", controller.BranchCreateController).Methods(http.MethodPost)
}
