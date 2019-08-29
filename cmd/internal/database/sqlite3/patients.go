package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"vehicles-microservice/internal/vehicle"
)

const VehicleTable = "Vehicles"

type vehicleRepository struct {
	db *sql.DB
}

func NewSqlite3VehicleRepository(db *sql.DB) vehicle.VehicleRepository {
	return &vehicleRepository{
		db,
	}
}

func (r *vehicleRepository) Create(vehicle *vehicle.Vehicle) error {

	statement, err := r.db.Prepare("INSERT INTO vehicles(id, vehicle_number, year, make, model, vin, created, updated, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
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
	}
	return nil
}

func (r *vehicleRepository) FindByID(id string) (*vehicle.Vehicle, error) {
	vehicle := new(vehicle.Vehicle)
	err := r.db.QueryRow("SELECT id, vehicle_number, year, make, model, vin, created, updated, deleted FROM vehicles where id=$1", id).Scan(&vehicle.ID, &vehicle.VehicleNumber, &vehicle.Year, &vehicle.Make, &vehicle.Model, &vehicle.VIN, &vehicle.Created, &vehicle.Updated, &vehicle.Deleted)
	if err != nil {
		panic(err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) FindAll() (vehicles []*vehicle.Vehicle, err error) {
	rows, err := r.db.Query("SELECT id, vehicle_number, year, make, model, vin, created, updated, deleted FROM vehicles")
	defer rows.Close()

	for rows.Next() {
		vehicle := new(vehicle.Vehicle)
		tmCreated := ""
		tmUpdated := ""
		tmDeleted := ""

		if err = rows.Scan(&vehicle.ID, &vehicle.VehicleNumber, &vehicle.Year, &vehicle.Make, &vehicle.Model, &vehicle.VIN, &tmCreated, &tmUpdated, &tmDeleted); err != nil {
			log.Print(err)
			return nil, err
		}

		t, err := time.Parse(time.RFC3339Nano, tmCreated)
		fmt.Println(t, err)
		vehicle.Created = t
		t, err = time.Parse(time.RFC3339Nano, tmUpdated)
		fmt.Println(t, err)
		vehicle.Updated = t
		t, err = time.Parse(time.RFC3339Nano, tmDeleted)
		fmt.Println(t, err)
		vehicle.Deleted = t

		vehicles = append(vehicles, vehicle)

	}
	return vehicles, nil
}

// DeleteByID attempts to delete a vehicle in a sqlite3 repository
func (r *vehicleRepository) DeleteByID(id string) error {
	_, err := r.db.Exec("DELETE FROM vehicles where id=$1", id)

	if err != nil {
		panic(err)
	}

	return nil
}
