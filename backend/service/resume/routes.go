package resume

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

type ResumeHandler struct {
	store types.ResumeStore
}

func NewResumeHandler(store types.ResumeStore) *ResumeHandler {
	return &ResumeHandler{store: store}
}

func (h *ResumeHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/resume", h.createResume).Methods("POST")
	router.HandleFunc("/resumes", h.getResumes).Methods("GET")
	router.HandleFunc("/resume/{id}", h.getResume).Methods("GET")
	router.HandleFunc("/resume/{id}", h.deleteResume).Methods("DELETE")
}

func (h *ResumeHandler) createResume(w http.ResponseWriter, r *http.Request) {
	var resume types.CreateResumePayload

	if err := utils.ParseJson(r, &resume); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(resume); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.store.CreateResume(types.Resume{
		EmployeeID:     resume.EmployeeID,
		Education:      resume.Education,
		Experience:     resume.Experience,
		Skills:         resume.Skills,
		DisabilityInfo: resume.DisabilityInfo,
		AdditionalInfo: resume.AdditionalInfo,
		JobExpectation: resume.JobExpectation,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, resume)
}

func (h *ResumeHandler) getResume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resumeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resume, err := h.store.GetResume(resumeId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.WriteError(w, http.StatusNotFound, err)
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJson(w, http.StatusOK, resume); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *ResumeHandler) getResumes(w http.ResponseWriter, r *http.Request) {
	resumes, err := h.store.GetResumes()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJson(w, http.StatusOK, resumes); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *ResumeHandler) deleteResume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resumeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.DeleteResume(resumeId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, nil)
}
