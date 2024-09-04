package vacancy

import (
	"fmt"
	"net/http"
	"recruitment/types"
	"recruitment/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VacancyHandler struct {
	Store *Store
}

func NewVacancyHandler(Store *Store) *VacancyHandler {
	return &VacancyHandler{Store: Store}
}

func (h *VacancyHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/vacancy", h.createVacancy).Methods("POST")
	router.HandleFunc("/vacancies", h.getVacancies).Methods("GET")
	router.HandleFunc("/vacancy/{id}", h.getVacancy).Methods("GET")
	router.HandleFunc("/vacancy/{id}", h.deleteVacancy).Methods("DELETE")
}

func (h *VacancyHandler) createVacancy(w http.ResponseWriter, r *http.Request) {
	var vacancy types.CreateVacancyPayload

	if err := utils.ParseJson(r, &vacancy); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(vacancy); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.Store.CreateVacancy(types.Vacancy{
		EmployerID:           vacancy.EmployerID,
		Position:             vacancy.Position,
		Responsibilities:     vacancy.Responsibilities,
		QualificationReqs:    vacancy.QualificationReqs,
		WorkSchedule:         vacancy.WorkSchedule,
		WorkingConditions:    vacancy.WorkingConditions,
		SocialBenefits:       vacancy.SocialBenefits,
		EmployeeExpectations: vacancy.EmployeeExpectations,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, vacancy)

}

func (h *VacancyHandler) getVacancies(w http.ResponseWriter, r *http.Request) {
	var vacancies []types.Vacancy

	vacancies, err := h.Store.GetVacancies()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, vacancies)
}

func (h *VacancyHandler) getVacancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vacancyId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	vacancy, err := h.Store.GetVacancy(vacancyId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.WriteError(w, http.StatusNotFound, err)
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	utils.WriteJson(w, http.StatusOK, vacancy)
}

func (h *VacancyHandler) deleteVacancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vacancyId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.Store.DeleteVacancy(vacancyId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, nil)
}
