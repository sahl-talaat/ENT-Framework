package router

import (
	controller "entdemo/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerProductRouter(r *mux.Router) {
	employeeRouter := r.PathPrefix("api/v1/prodycts").Subrouter()
	employeeRouter.HandleFunc("/{id}", controller.ProductGetByIDController).Methods(http.MethodGet)
	employeeRouter.HandleFunc("/", controller.ProductCreateController).Methods(http.MethodPost)
	employeeRouter.HandleFunc("/{id}", controller.ProductUpdateController).Methods(http.MethodPut)
	employeeRouter.HandleFunc("/{id}", controller.ProductDeleteController).Methods(http.MethodDelete)
}
