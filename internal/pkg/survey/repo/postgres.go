package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	createResult = `
	INSERT INTO results (id, user_id, survey_id) VALUES ($1, $2, $3);
	`

	createResponse = `
	INSERT INTO answer (id, question, result_id, answer) VALUES ($1, $2, $3, $4);
	`

	getSurveyInfo = `
	SELECT id as survey_id, name
	FROM survey
	WHERE id = $1;`

	getSurveyQuestions = `
	SELECT q.id as question_id, q.name as question_name, qt.type as question_type
	FROM question q
	JOIN question_type qt ON q.type = qt.id
	WHERE q.id_survey = $1;`
)

var (
	ErrSurveyNotFound          = errors.New("survey not found")
	ErrSurveyQuestionsNotFound = errors.New("survey questions not found")
)

type SurveyRepo struct {
	db pgxtype.Querier
}

func NewSurveyRepo(db pgxtype.Querier) *SurveyRepo {
	return &SurveyRepo{
		db: db,
	}
}

func (r *SurveyRepo) CreateResult(ctx context.Context, survey_id uuid.UUID, user_id uuid.UUID) (uuid.UUID, error) {
	resultID := uuid.NewV4()
	_, err := r.db.Exec(ctx, createResult,
		resultID,
		user_id,
		survey_id,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return uuid.UUID{}, err
	}
	return resultID, nil
}

func (r *SurveyRepo) SaveResults(ctx context.Context, surveyResponse models.SurveyResponse) error {
	responseID := uuid.NewV4()
	_, err := r.db.Exec(ctx, createResponse,
		responseID,
		surveyResponse.QuestionID,
		surveyResponse.ResultID,
		surveyResponse.Answer,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}
	return nil
}

func (r *SurveyRepo) ReadSurvey(ctx context.Context, surveyID uuid.UUID) (models.Survey, error) {
	survey := models.Survey{}
	err := r.db.QueryRow(ctx, getSurveyInfo, surveyID).
		Scan(&survey.ID, &survey.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Survey{}, ErrSurveyNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Survey{}, err
	}

	rows, err := r.db.Query(ctx, getSurveyQuestions, surveyID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Survey{}, ErrSurveyQuestionsNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return models.Survey{}, err
	}
	questionsSlice := []models.Question{}
	question := models.Question{}
	for rows.Next() {
		err = rows.Scan(
			&question.ID,
			&question.Name,
			&question.QuestionType.Type,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return models.Survey{}, err
		}
		questionsSlice = append(questionsSlice, question)
	}
	defer rows.Close()

	survey.Questions = questionsSlice
	return survey, nil
}
