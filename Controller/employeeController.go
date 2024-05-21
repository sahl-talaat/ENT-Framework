package controller

import (
	"encoding/json"
	service "entdemo/Service"
	utils "entdemo/Utils"
	"entdemo/ent"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func EmployeeGetByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	employee, err := service.NewEmployeeOps(r.Context()).EmployeeGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, employee)
}

func EmployeeCreateController(w http.ResponseWriter, r *http.Request) {

	var newEmployee ent.Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdEmployee, err := service.NewEmployeeOps(r.Context()).EmployeeCreate(newEmployee)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}

	utils.Return(w, true, http.StatusOK, nil, createdEmployee)
}

func EmployeeUpdateController(w http.ResponseWriter, r *http.Request) {

	var newEmployeeData ent.Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployeeData)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	newEmployeeData.ID = id

	updatedEmployee, err := service.NewEmployeeOps(r.Context()).EmployeeUpdate(newEmployeeData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedEmployee)
}

func EmployeeDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewEmployeeOps(r.Context()).EmployeeDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}
