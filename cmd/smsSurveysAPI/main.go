package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sms-surveys/internal/customer"
	"sms-surveys/internal/database/sqlite3"
	"sms-surveys/internal/env"
	"sms-surveys/internal/survey"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

const (
	DefaultHTTPPort      = "3004"
	DefaultRedisURL      = "localhost:6379"
	DefaultRedisPassword = "example"

	DefaultPostgresURL      = "postgres://example:postgres@localhost/vehicle?sslmode=disable"
	DefaultPostgresUser     = "postgres"
	DefaultPostgresHost     = "db"
	DefaultPostgresPort     = "5432"
	DefaultPostgresPassword = "example"
	DefaultPostgresDBName   = "customer"

	DefaultSqlite3File = "./sms-surveys.db"
)

func main() {
	errChan := make(chan error)

	Logger := zap.NewExample()
	defer Logger.Sync()
	Logger.Info("Welcome to sms-surveys-microservice")

	TwilioSID := env.EnvString("SVY_TWILIO_SID", "")
	TwilioToken := env.EnvString("SVY_TWILIO_TOKEN", "")
	HTTPPort := env.EnvString("SVY_HTTP_PORT", DefaultHTTPPort)

	fmt.Printf("TwilioSID = %s\n", TwilioSID)
	fmt.Printf("TwilioToken = %s\n", TwilioToken)

	var CustomerRepo customer.CustomerRepository
	var SurveyRepo survey.SurveyRepository

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	var dbURL string
	dbType := "sqlite3" // for now, we can add others later

	switch dbType {
	case "sqlite3":
		dbURL = env.EnvString("DATABASE_URL", DefaultSqlite3File)
		db, err := sql.Open("sqlite3", dbURL)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS customer (id TEXT PRIMARY KEY, last_name TEXT, first_name TEXT, phone TEXT, created_time INTEGER, updated_time INTEGER, deleted_time INTEGER)")
		statement.Exec()
		statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS survey (id TEXT PRIMARY KEY, description TEXT, from_ph_num TEXT, to_ph_num TEXT, flow_params TEXT, created_time INTEGER, updated_time INTEGER, deleted_time INTEGER)")
		statement.Exec()

		
		SurveyRepo = sqlite3.NewSqlite3SurveyRepository(db)
		CustomerRepo = sqlite3.NewSqlite3CustomerRepository(db)
	default:
		panic("Unknown database")
	}

	CustomerService := customer.NewCustomerService(CustomerRepo)
	CustomerHandler := customer.NewCustomerHandler(CustomerService)
	SurveyService := survey.NewSurveyService(SurveyRepo)
	SurveyHandler := survey.NewSurveyHandler(SurveyService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", CustomerHandler.Get).Methods("GET")
	router.HandleFunc("/customers/{id}", CustomerHandler.GetByID).Methods("GET")
	router.HandleFunc("/customers/{id}", CustomerHandler.DeleteByID).Methods("DELETE")
	router.HandleFunc("/customers", CustomerHandler.Create).Methods("POST")
	router.HandleFunc("/surveys", SurveyHandler.Get).Methods("GET")
	router.HandleFunc("/surveys/{id}", SurveyHandler.GetByID).Methods("GET")
	router.HandleFunc("/surveys/{id}", SurveyHandler.DeleteByID).Methods("DELETE")
	router.HandleFunc("/surveys", SurveyHandler.Create).Methods("POST")

	errs := make(chan error, 2)

	go func() {
		logrus.Info(fmt.Sprintf("Listening server mode on port : %s", HTTPPort))
		p := ":" + HTTPPort
		errs <- http.ListenAndServe(p, router)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("Err Chan %s", <-c)
	}()

	logrus.Error("sms-surveys microservice terminated", <-errs)

}
