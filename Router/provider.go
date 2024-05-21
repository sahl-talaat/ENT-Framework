package router

import "github.com/gorilla/mux"

func RegisterRouter(r *mux.Router) {
	registerEmployeeRouter(r)
	registerDepartmentRouter(r)
	registerProductRouter(r)
	registerOrderRouter(r)
	registerBranchRouter(r)
}
