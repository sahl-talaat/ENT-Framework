package controller

import (
	"encoding/json"
	service "entdemo/Service"
	utils "entdemo/Utils"
	"entdemo/ent"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DepartmentGetByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	employee, err := service.NewDepartmentOps(r.Context()).DepartmentGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, employee)
}

func DepartmentCreateController(w http.ResponseWriter, r *http.Request) {

	var newDepartment ent.Department
	err := json.NewDecoder(r.Body).Decode(&newDepartment)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdDepartment, err := service.NewDepartmentOps(r.Context()).DepartmentCreate(newDepartment)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}

	utils.Return(w, true, http.StatusOK, nil, createdDepartment)
}

func DepartmentUpdateController(w http.ResponseWriter, r *http.Request) {

	var newDepartmentData ent.Department
	err := json.NewDecoder(r.Body).Decode(&newDepartmentData)
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
	newDepartmentData.ID = id

	updatedDepartment, err := service.NewDepartmentOps(r.Context()).DepartmentUpdate(newDepartmentData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedDepartment)
}

func DepartmentDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewDepartmentOps(r.Context()).DepartmentDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}
