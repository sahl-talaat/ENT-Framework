package router

import (
	controller "entdemo/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerDepartmentRouter(r *mux.Router) {
	departmentRouter := r.PathPrefix("api/v1/departments").Subrouter()
	departmentRouter.HandleFunc("/{id}", controller.DepartmentGetByIDController).Methods(http.MethodGet)
	departmentRouter.HandleFunc("/", controller.DepartmentCreateController).Methods(http.MethodPost)
	departmentRouter.HandleFunc("/{id}", controller.DepartmentUpdateController).Methods(http.MethodPut)
	departmentRouter.HandleFunc("/{id}", controller.DepartmentDeleteController).Methods(http.MethodDelete)
}
