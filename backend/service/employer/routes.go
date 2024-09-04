package employer

import (
	"fmt"
	"net/http"
	"recruitment/types"
	"recruitment/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployerHandler struct {
	Store *Store
}

func NewEmployerHandler(store *Store) *EmployerHandler {
	return &EmployerHandler{Store: store}
}

func (h *EmployerHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/employer", h.createEmployer).Methods("POST")
	router.HandleFunc("/employers", h.getEmployers).Methods("GET")
	router.HandleFunc("/employer/{id}", h.getEmployer).Methods("GET")
	// router.HandleFunc("/employer/{id}", h.updateEmployer).Methods("PUT")
	router.HandleFunc("/employer/{id}", h.deleteEmployer).Methods("DELETE")
}

func (h *EmployerHandler) createEmployer(w http.ResponseWriter, r *http.Request) {
	var employer types.Employer
	
	if err := utils.ParseJson(r, &employer); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(employer); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.Store.CreateEmployer(employer); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, employer)
}

func (h *EmployerHandler) getEmployers(w http.ResponseWriter, r *http.Request) {
	var employers []types.Employer

	employers, err := h.Store.GetEmployers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, employers)
}

func (h *EmployerHandler) getEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	employer, err := h.Store.GetEmployer(employerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, employer)
}

func (h *EmployerHandler) deleteEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.Store.DeleteEmployer(employerId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, nil)
}
