package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"sms-surveys/internal/customer"
	"strconv"
	"time"
)

type customerRepository struct {
	db *sql.DB
}

func NewSqlite3CustomerRepository(db *sql.DB) customer.CustomerRepository {
	return &customerRepository{
		db,
	}
}

func (r *customerRepository) Create(customer *customer.Customer) error {

	statement, err := r.db.Prepare("INSERT INTO customer(id, last_name, first_name, phone, created_time, updated_time, deleted_time) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err)
	}

	tmCreated := strconv.FormatInt(customer.CreatedTime.Unix(), 10)
	tmUpdated := strconv.FormatInt(customer.UpdatedTime.Unix(), 10)
	tmDeleted := strconv.FormatInt(customer.DeletedTime.Unix(), 10)
	res, err := statement.Exec(customer.ID, customer.LastName, customer.FirstName, customer.Phone, tmCreated, tmUpdated, tmDeleted)

	fmt.Printf("res=%#v\n", res)

	if err != nil {
		panic(err)
	}

	return nil
}

func (r *customerRepository) FindByID(id string) (*customer.Customer, error) {
	customer := new(customer.Customer)
	tmCreated := ""
	tmUpdated := ""
	tmDeleted := ""
	err := r.db.QueryRow("SELECT id, last_name, first_name, phone, created_time, updated_time, deleted_time FROM customer where id=$1", id).Scan(&customer.ID, &customer.LastName, &customer.FirstName, &customer.Phone, &tmCreated, &tmUpdated, &tmDeleted)
	if err != nil {
		panic(err)
	}
	return customer, nil
}

func (r *customerRepository) FindAll() (customers []*customer.Customer, err error) {
	rows, err := r.db.Query("SELECT id, last_name, first_name, phone, created_time, updated_time, deleted_time FROM customer")
	if err != nil {
		panic("customer.go FindAll() error = " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		customer := new(customer.Customer)
		var tmCreated int64
		var tmUpdated int64
		var tmDeleted int64

		if err = rows.Scan(&customer.ID, &customer.LastName, &customer.FirstName, &customer.Phone, &tmCreated, &tmUpdated, &tmDeleted); err != nil {
			log.Print(err)
			return nil, err
		}

		t := time.Unix(tmCreated, 0)
		fmt.Println(t, err)
		customer.CreatedTime = t
		t = time.Unix(tmUpdated, 0)
		fmt.Println(t, err)
		customer.UpdatedTime = t
		t = time.Unix(tmDeleted, 0)
		fmt.Println(t, err)
		customer.DeletedTime = t

		customers = append(customers, customer)

	}
	return customers, nil
}

// DeleteByID attempts to delete a vehicle in a sqlite3 repository
func (r *customerRepository) DeleteByID(id string) error {
	_, err := r.db.Exec("DELETE FROM customer where id=$1", id)

	if err != nil {
		panic(err)
	}

	return nil
}
