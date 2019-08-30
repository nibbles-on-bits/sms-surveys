package customer

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type CustomerService interface {
	CreateCustomer(customer *Customer) error
	FindCustomerByID(id string) (*Customer, error)
	FindAllCustomers() ([]*Customer, error)
	DeleteCustomerByID(id string) error
}

type customerService struct {
	repo CustomerRepository
}

// NewCustomerService will bind a repository
func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerService{
		repo,
	}
}

func (s *customerService) CreateCustomer(customer *Customer) error {
	customer.ID = uuid.New().String()
	customer.CreatedTime = time.Now()
	customer.UpdatedTime = time.Now()

	if err := s.repo.Create(customer); err != nil {
		logrus.WithField("error", err).Error("Error creating customer")
		return err
	}

	logrus.WithField("id", customer.ID).Info("Created new customer")
	return nil
}

func (s *customerService) FindCustomerByID(id string) (*Customer, error) {
	customer, err := s.repo.FindByID(id)

	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error finding customer")
		return nil, err
	}
	logrus.WithField("id", id).Info("Found customer")
	return customer, nil
}

func (s *customerService) FindAllCustomers() ([]*Customer, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error finding all customers")
		return nil, err
	}
	logrus.Info("Found all customers")
	return customers, nil
}

func (s *customerService) DeleteCustomerByID(id string) error {
	err := s.repo.DeleteByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error deleting customer")
		return err
	}
	return err
}
