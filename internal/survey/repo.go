package survey

type SurveyRepository interface {
	Create(survey *Survey) error
	FindAll() ([]*Survey, error)
	FindByID(id string) (*Survey, error)
	DeleteByID(id string) error
	//PatchById(id string) error
}
