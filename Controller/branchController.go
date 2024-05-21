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

func BranchGetByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	employee, err := service.NewBranchOps(r.Context()).BranchGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, employee)
}

func BranchCreateController(w http.ResponseWriter, r *http.Request) {

	var newBranch ent.Branch
	err := json.NewDecoder(r.Body).Decode(&newBranch)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdBranch, err := service.NewBranchOps(r.Context()).BranchCreate(newBranch)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}

	utils.Return(w, true, http.StatusOK, nil, createdBranch)
}

func BranchUpdateController(w http.ResponseWriter, r *http.Request) {

	var newBranchData ent.Branch
	err := json.NewDecoder(r.Body).Decode(&newBranchData)
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
	newBranchData.ID = id

	updatedBranch, err := service.NewBranchOps(r.Context()).BranchUpdate(newBranchData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedBranch)
}

func BranchDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewBranchOps(r.Context()).BranchDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}
