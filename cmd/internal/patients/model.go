package patients

type Patient struct {
	LastName  string `json:"lastName" db:"last_name"`
	FirstName string `json:"firstName" db:"first_name"`
	Phone     string `json:"phone" db:"phone"`
}
