package survey

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/survey_mock.go -package mock

type SurveyUsecase interface {
	// SaveSurvey(context.Context, models.Survey)
	GetSurvey(context.Context, uuid.UUID) (models.Survey, error)
	GetStat(context.Context, uuid.UUID) (models.Statistics, error)
}

type SurveyRepo interface {
	ReadSurvey(context.Context, uuid.UUID) (models.Survey, error)
	GetAnswers(context.Context, uuid.UUID) ([]models.Answer, error)
	// WriteSurvey(context.Context, models.Survey) error
	// ReadCompletedSurveys(context.Context, uuid.UUID) (models.SurveysCompleted, error)
}
