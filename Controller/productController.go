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

func ProductGetByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	employee, err := service.NewProductOps(r.Context()).ProductGetByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, employee)
}

func ProductCreateController(w http.ResponseWriter, r *http.Request) {

	var newProduct ent.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdProduct, err := service.NewProductOps(r.Context()).ProductCreate(newProduct)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}

	utils.Return(w, true, http.StatusOK, nil, createdProduct)
}

func ProductUpdateController(w http.ResponseWriter, r *http.Request) {

	var newProductData ent.Product
	err := json.NewDecoder(r.Body).Decode(&newProductData)
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
	newProductData.ID = id

	updatedProduct, err := service.NewProductOps(r.Context()).ProductUpdate(newProductData)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, updatedProduct)
}

func ProductDeleteController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewProductOps(r.Context()).ProductDelete(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, deletedID)
}
