package survey

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type SurveyHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type surveyHandler struct {
	surveyService SurveyService
}

func NewSurveyHandler(surveyService SurveyService) SurveyHandler {
	return &surveyHandler{
		surveyService,
	}
}

func (h *surveyHandler) Get(w http.ResponseWriter, r *http.Request) {
	surveys, err := h.surveyService.FindAllSurveys()
	if err != nil {
		logrus.WithField("error", err).Error("Unable to find all surveys")
		http.Error(w, "Unable to find all surveys", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(surveys)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to get smsReminder", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *surveyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	smsReminder, err := h.surveyService.FindSurveyByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Unable to find survey")
		http.Error(w, "Unable to find survey", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(smsReminder)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to fetch survey", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *surveyHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {

}

func (h *surveyHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler.go DeleteByID() called")
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.surveyService.DeleteSurveyByID(id)
	if err != nil {
		logrus.WithField("error", err).Error("Error calling surveyService.DeleteSurveyByID")
		http.Error(w, "Unable to delete survey", http.StatusInternalServerError)
		return
	}
}

func (h *surveyHandler) Create(w http.ResponseWriter, r *http.Request) {

	var survey Survey
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&survey); err != nil {
		logrus.Error("Unable to decode survey", err)
		http.Error(w, "Bad format for survey", http.StatusBadRequest)
		return
	}

	if err := h.surveyService.CreateSurvey(&survey); err != nil {
		logrus.WithField("error", err).Error("Unable to create smsReminder")
		http.Error(w, "Unable to create smsReminder", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(survey)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to create smsReminder", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}
