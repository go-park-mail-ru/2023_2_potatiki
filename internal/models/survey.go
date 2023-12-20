package models

import uuid "github.com/satori/go.uuid"

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/survey.go

type SurveysCompleted struct {
	CompletedSurveys []string `json:"completedSurveys"`
}

type Survey struct {
	ID        uuid.UUID `json:"surveyId"`
	Name      string    `json:"surveyName"`
	ResultID  uuid.UUID `json:"resultId"`
	Questions []Question
}

type Question struct {
	ID   uuid.UUID `json:"questionId"`
	Name string    `json:"questionName"`
	QuestionType
}

type QuestionType struct {
	ID   int    `json:"-"`
	Type string `json:"questionType"`
}

type Answer struct {
	ID     uuid.UUID `json:"-"`
	Answer int       `json:"-"`
}

type SurveyResponse struct {
	QuestionID uuid.UUID `json:"questionId"`
	ResultID   uuid.UUID `json:"resultId"`
	Answer     int       `json:"answer"`
	UserID     uuid.UUID `json:"-"`
}

type Stat struct {
	QuestionName string  `json:"questionName"`
	StatValue    float64 `json:"statValue"`
}

type Statistics []Stat
