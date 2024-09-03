package employee

import (
	"net/http"
	"recruitment/types"
	"recruitment/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	var employee types.CreateEmployeePayload

	if err := utils.ParseJson(r, &employee); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.CreateEmployee(types.Employee{
		ID:        primitive.NewObjectID(),
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		BirthDate: employee.BirthDate,
		Telephone: employee.Telephone,
		Address:   employee.Address,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, employee)
}
func (h *EmployeeHandler) getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	employee, err := h.store.GetEmployee(employeeId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.WriteError(w, http.StatusNotFound, err)
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
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

	if len(employees) == 0 {
		utils.WriteJson(w, http.StatusOK, []types.Employee{})
	}

	utils.WriteJson(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := primitive.ObjectIDFromHex(id)
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
