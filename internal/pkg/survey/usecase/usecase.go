package usecase

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/survey"
	uuid "github.com/satori/go.uuid"
)

type SurveyUsecase struct {
	repo survey.SurveyRepo
}

func NewSurveyUsecase(repo survey.SurveyRepo) *SurveyUsecase {
	return &SurveyUsecase{
		repo: repo,
	}
}

//func (uc *SurveyUsecase) SaveSurvey(ctx context.Context, surveyInfo models.Survey) error {
//	err := uc.repo.WriteSurvey(ctx, surveyInfo)
//	surveyInfo.ID = uuid.NewV4()
//	if err != nil {
//		//if errors.Is(err, repo.ErrProductNotFound) {
//		//	return err
//		//}
//		err = fmt.Errorf("error happened in repo.WriteSurvey: %w", err)
//
//		return err
//	}
//
//	return nil
//
//}

func (uc *SurveyUsecase) GetSurvey(ctx context.Context, surveyID uuid.UUID, userID uuid.UUID) (models.Survey, error) {
	survey, err := uc.repo.ReadSurvey(ctx, surveyID)
	if err != nil {
		//if errors.Is(err, repo.ErrProductNotFound) {
		//	return err
		//}
		err = fmt.Errorf("error happened in repo.GetSurvey: %w", err)

		return models.Survey{}, err
	}

	resultID, err := uc.repo.CreateResult(ctx, surveyID, userID)
	if err != nil {
		//if errors.Is(err, repo.ErrProductNotFound) {
		//	return err
		//}
		err = fmt.Errorf("error happened in repo.CreateResult: %w", err)

		return models.Survey{}, err
	}
	survey.ResultID = resultID

	return survey, nil

}
