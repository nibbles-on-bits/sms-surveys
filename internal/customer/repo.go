package customer

type CustomerRepository interface {
	Create(customer *Customer) error
	FindAll() ([]*Customer, error)
	FindByID(id string) (*Customer, error)
	DeleteByID(id string) error
	//PatchById(id string) error
}
