package survey

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type SurveyService interface {
	CreateSurvey(survey *Survey) error
	FindSurveyByID(id string) (*Survey, error)
	FindAllSurveys() ([]*Survey, error)
	DeleteSurveyByID(id string) error
}

type surveyService struct {
	repo SurveyRepository
}

// NewSurveyService will bind a repository
func NewSurveyService(repo SurveyRepository) SurveyService {
	return &surveyService{
		repo,
	}
}

func (s *surveyService) CreateSurvey(survey *Survey) error {
	survey.ID = uuid.New().String()
	survey.CreatedTime = time.Now()
	survey.UpdatedTime = time.Now()

	if err := s.repo.Create(survey); err != nil {
		logrus.WithField("error", err).Error("Error creating survey")
		return err
	}

	logrus.WithField("id", survey.ID).Info("Created new survey")
	return nil
}

func (s *surveyService) FindSurveyByID(id string) (*Survey, error) {
	survey, err := s.repo.FindByID(id)

	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error finding survey")
		return nil, err
	}
	logrus.WithField("id", id).Info("Found survey")
	return survey, nil
}

func (s *surveyService) FindAllSurveys() ([]*Survey, error) {
	surveys, err := s.repo.FindAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error finding all surveys")
		return nil, err
	}
	logrus.Info("Found all surveys")
	return surveys, nil
}

func (s *surveyService) DeleteSurveyByID(id string) error {
	err := s.repo.DeleteByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error deleting survey")
		return err
	}
	return err
}
