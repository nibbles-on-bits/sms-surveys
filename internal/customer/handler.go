package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type CustomerHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	customerService CustomerService
}

func NewCustomerHandler(customerService CustomerService) CustomerHandler {
	return &customerHandler{
		customerService,
	}
}

func (h *customerHandler) Get(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.FindAllCustomers()
	if err != nil {
		logrus.WithField("error", err).Error("Unable to find all customers")
		http.Error(w, "Unable to find all customers", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(customers)
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

func (h *customerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	smsReminder, err := h.customerService.FindCustomerByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Unable to find customer")
		http.Error(w, "Unable to find customer", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(smsReminder)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to fetch customer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *customerHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {

}

func (h *customerHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler.go DeleteByID() called")
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.customerService.DeleteCustomerByID(id)
	if err != nil {
		logrus.WithField("error", err).Error("Error calling customerService.DeleteCustomerByID")
		http.Error(w, "Unable to delete customer", http.StatusInternalServerError)
		return
	}
}

func (h *customerHandler) Create(w http.ResponseWriter, r *http.Request) {

	var customer Customer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		logrus.Error("Unable to decode customer", err)
		http.Error(w, "Bad format for customer", http.StatusBadRequest)
		return
	}

	if err := h.customerService.CreateCustomer(&customer); err != nil {
		logrus.WithField("error", err).Error("Unable to create smsReminder")
		http.Error(w, "Unable to create smsReminder", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(customer)
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
