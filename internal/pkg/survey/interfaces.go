package survey

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/survey_mock.go -package mock

type SurveyUsecase interface {
	SaveResponse(context.Context, models.SurveyResponse) error
	GetSurvey(context.Context, uuid.UUID, uuid.UUID) (models.Survey, error)
}

type SurveyRepo interface {
	SaveResults(context.Context, models.SurveyResponse) error
	ReadSurvey(context.Context, uuid.UUID) (models.Survey, error)
	CreateResult(context.Context, uuid.UUID, uuid.UUID) (uuid.UUID, error)
	// WriteSurvey(context.Context, models.Survey) error
	// ReadCompletedSurveys(context.Context, uuid.UUID) (models.SurveysCompleted, error)
}
