package customer

import "time"

type Customer struct {
	ID          string    `json:"id" db:"id"`
	LastName    string    `json:"lastName" db:"last_name"`
	FirstName   string    `json:"firstName" db:"first_name"`
	Phone       string    `json:"phone" db:"phone"`
	CreatedTime time.Time `json:"createdTime" db:"created_time"`
	UpdatedTime time.Time `json:"updatedTime" db:"updated_time"`
	DeletedTime time.Time `json:"deletedTime" db:"deleted_time"`
}
