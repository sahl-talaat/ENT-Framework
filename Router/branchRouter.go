package router

import (
	controller "entdemo/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerBranchRouter(r *mux.Router) {
	employeeRouter := r.PathPrefix("api/v1/branches").Subrouter()
	employeeRouter.HandleFunc("/{id}", controller.BranchGetByIDController).Methods(http.MethodGet)
	employeeRouter.HandleFunc("/", controller.BranchCreateController).Methods(http.MethodPost)
	employeeRouter.HandleFunc("/{id}", controller.BranchUpdateController).Methods(http.MethodPut)
	employeeRouter.HandleFunc("/{id}", controller.BranchDeleteController).Methods(http.MethodDelete)
}
