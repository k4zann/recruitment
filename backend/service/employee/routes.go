package employee

import (
	"net/http"
	"recruitment/types"
	"recruitment/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	store types.EmployeeStore
}

func NewEmployeeHandler(store types.EmployeeStore) *EmployeeHandler {
	return &EmployeeHandler{store: store}
}

func (h *EmployeeHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/employee", h.createEmployee).Methods("POST")
	router.HandleFunc("/employees", h.getEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", h.getEmployee).Methods("GET")
	// router.HandleFunc("/employee/{id}", h.updateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", h.deleteEmployee).Methods("DELETE")
}

func (h *EmployeeHandler) createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee types.Employee

	if err := utils.ParseJson(r, &employee); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.CreateEmployee(employee); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, employee)
}

func (h *EmployeeHandler) getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	employee, err := h.store.GetEmployee(employeeId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) getEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.store.GetEmployees()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.DeleteEmployee(employeeId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, nil)
}
