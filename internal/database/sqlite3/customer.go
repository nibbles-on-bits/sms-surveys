package sqlite3

import (
	"database/sql"
	"sms-surveys/internal/customer"
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

	/*statement, err := r.db.Prepare("INSERT INTO vehicles(id, vehicle_number, year, make, model, vin, created, updated, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		panic(err)
	}

	tmCreated := vehicle.Created.UTC().Format(time.RFC3339)
	tmUpdated := vehicle.Updated.UTC().Format(time.RFC3339)
	tmDeleted := vehicle.Deleted.UTC().Format(time.RFC3339)
	res, err := statement.Exec(vehicle.ID, vehicle.VehicleNumber, vehicle.Year, vehicle.Make, vehicle.Model, vehicle.VIN, tmCreated, tmUpdated, tmDeleted)

	fmt.Printf("res=%#v\n", res)

	if err != nil {
		panic(err)
	}*/

	return nil
}

func (r *customerRepository) FindByID(id string) (*customer.Customer, error) {
	customer := new(customer.Customer)
	// err := r.db.QueryRow("SELECT id, vehicle_number, year, make, model, vin, created, updated, deleted FROM vehicles where id=$1", id).Scan(&vehicle.ID, &vehicle.VehicleNumber, &vehicle.Year, &vehicle.Make, &vehicle.Model, &vehicle.VIN, &vehicle.Created, &vehicle.Updated, &vehicle.Deleted)
	// if err != nil {
	// 	panic(err)
	// }
	return customer, nil
}

func (r *customerRepository) FindAll() (customers []*customer.Customer, err error) {
	// rows, err := r.db.Query("SELECT id, vehicle_number, year, make, model, vin, created, updated, deleted FROM vehicles")
	// defer rows.Close()

	// for rows.Next() {
	// 	vehicle := new(vehicle.Vehicle)
	// 	tmCreated := ""
	// 	tmUpdated := ""
	// 	tmDeleted := ""

	// 	if err = rows.Scan(&vehicle.ID, &vehicle.VehicleNumber, &vehicle.Year, &vehicle.Make, &vehicle.Model, &vehicle.VIN, &tmCreated, &tmUpdated, &tmDeleted); err != nil {
	// 		log.Print(err)
	// 		return nil, err
	// 	}

	// 	t, err := time.Parse(time.RFC3339Nano, tmCreated)
	// 	fmt.Println(t, err)
	// 	vehicle.Created = t
	// 	t, err = time.Parse(time.RFC3339Nano, tmUpdated)
	// 	fmt.Println(t, err)
	// 	vehicle.Updated = t
	// 	t, err = time.Parse(time.RFC3339Nano, tmDeleted)
	// 	fmt.Println(t, err)
	// 	vehicle.Deleted = t

	// 	vehicles = append(vehicles, vehicle)

	// }
	// return vehicles, nil
	return nil, nil
}

// DeleteByID attempts to delete a vehicle in a sqlite3 repository
func (r *customerRepository) DeleteByID(id string) error {
	// _, err := r.db.Exec("DELETE FROM vehicles where id=$1", id)

	// if err != nil {
	// 	panic(err)
	// }

	return nil
}
