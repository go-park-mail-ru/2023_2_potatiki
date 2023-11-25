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

func (uc *SurveyUsecase) SaveResponse(ctx context.Context, surveyInfo models.SurveyResponse) error {
	err := uc.repo.SaveResults(ctx, surveyInfo)
	if err != nil {
		//if errors.Is(err, repo.ErrProductNotFound) {
		//	return err
		//}
		err = fmt.Errorf("error happened in repo.SaveResults: %w", err)

		return err
	}

	return nil

}

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

func (uc *SurveyUsecase) GetStat(ctx context.Context, surveyID uuid.UUID) (models.Statistics, error) {
	survey, err := uc.repo.ReadSurvey(ctx, surveyID)
	if err != nil {
		//if errors.Is(err, repo.ErrProductNotFound) {
		//	return err
		//}
		err = fmt.Errorf("error happened in repo.GetSurvey: %w", err)

		return models.Statistics{}, err
	}
	statistics := make([]models.Stat, 0, len(survey.Questions))
	for _, q := range survey.Questions {
		answers, err := uc.repo.GetAnswers(ctx, q.ID)
		if err != nil {
			//if errors.Is(err, repo.ErrProductNotFound) {
			//	return err
			//}
			err = fmt.Errorf("error happened in repo.GetSurvey: %w", err)

			return models.Statistics{}, err
		}
		sum := 0.0
		for _, a := range answers {
			sum += float64(a.Answer)
		}

		statistics = append(statistics, models.Stat{QuestionName: q.Name, StatValue: sum / float64(len(answers))})
	}

	return statistics, nil

}
