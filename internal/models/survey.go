package models

import uuid "github.com/satori/go.uuid"

type SurveysCompleted struct {
	CompletedSurveys []string `json:"completedSurveys"`
}

type Survey struct {
	ID        uuid.UUID `json:"surveyID"`
	Name      string    `json:"surveyName"`
	Questions []Question
}

type Question struct {
	ID   uuid.UUID `json:"questionID"`
	Name string    `json:"questionName"`
	QuestionType
}

type QuestionType struct {
	ID   int    `json:"-"`
	Type string `json:"questionType"`
}

type Answer struct {
	ID     int `json:"-"`
	Answer int `json:"-"`
}

type Stat struct {
	QuestionName string  `json:"questionName"`
	StatValue    float64 `json:"statValue"`
}

type Statistics []Stat
