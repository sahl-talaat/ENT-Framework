package router

import (
	controller "entdemo/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerEmployeeRouter(r *mux.Router) {
	employeeRouter := r.PathPrefix("api/v1/empoyees").Subrouter()
	employeeRouter.HandleFunc("/{id}", controller.EmployeeGetByIDController).Methods(http.MethodGet)
	employeeRouter.HandleFunc("/", controller.EmployeeCreateController).Methods(http.MethodPost)
	employeeRouter.HandleFunc("/{id}", controller.EmployeeUpdateController).Methods(http.MethodPut)
	employeeRouter.HandleFunc("/{id}", controller.EmployeeDeleteController).Methods(http.MethodDelete)
}
