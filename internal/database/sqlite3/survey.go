package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"sms-surveys/internal/survey"
	"strconv"
	"time"
)

type surveyRepository struct {
	db *sql.DB
}

func NewSqlite3SurveyRepository(db *sql.DB) survey.SurveyRepository {
	return &surveyRepository{
		db,
	}
}

func (r *surveyRepository) Create(survey *survey.Survey) error {

	statement, err := r.db.Prepare("INSERT INTO survey(id, description, from_ph_num, to_ph_num, flow_params, created_time, updated_time, deleted_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		panic(err)
	}

	tmCreated := strconv.FormatInt(survey.CreatedTime.Unix(), 10)
	tmUpdated := strconv.FormatInt(survey.UpdatedTime.Unix(), 10)
	tmDeleted := strconv.FormatInt(survey.DeletedTime.Unix(), 10)
	res, err := statement.Exec(survey.ID, survey.Description, survey.FromPhNum, survey.ToPhNum, survey.FlowParams, tmCreated, tmUpdated, tmDeleted)

	fmt.Printf("res=%#v\n", res)

	if err != nil {
		panic(err)
	}

	return nil
}

func (r *surveyRepository) FindByID(id string) (*survey.Survey, error) {
	survey := new(survey.Survey)
	tmCreated := ""
	tmUpdated := ""
	tmDeleted := ""
	err := r.db.QueryRow("SELECT id, description, from_ph_num, to_ph_num, flow_params, created_time, updated_time, deleted_time FROM survey where id=$1", id).Scan(&survey.ID, &survey.Description, &survey.FromPhNum, &survey.ToPhNum, &survey.FlowParams, &tmCreated, &tmUpdated, &tmDeleted)
	if err != nil {
		panic(err)
	}
	return survey, nil
}

func (r *surveyRepository) FindAll() (surveys []*survey.Survey, err error) {
	rows, err := r.db.Query("SELECT id, description, from_ph_num, to_ph_num, flow_params, created_time, updated_time, deleted_time FROM survey")
	if err != nil {
		panic("survey.go FindAll() error = " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		survey := new(survey.Survey)
		var tmCreated int64
		var tmUpdated int64
		var tmDeleted int64

		if err = rows.Scan(&survey.ID, &survey.Description, &survey.FromPhNum, &survey.ToPhNum, &survey.FlowParams, &tmCreated, &tmUpdated, &tmDeleted); err != nil {
			log.Print(err)
			return nil, err
		}

		t := time.Unix(tmCreated, 0)
		fmt.Println(t, err)
		survey.CreatedTime = t
		t = time.Unix(tmUpdated, 0)
		fmt.Println(t, err)
		survey.UpdatedTime = t
		t = time.Unix(tmDeleted, 0)
		fmt.Println(t, err)
		survey.DeletedTime = t

		surveys = append(surveys, survey)

	}
	return surveys, nil
}

// DeleteByID attempts to delete a vehicle in a sqlite3 repository
func (r *surveyRepository) DeleteByID(id string) error {
	_, err := r.db.Exec("DELETE FROM survey where id=$1", id)

	if err != nil {
		panic(err)
	}

	return nil
}
