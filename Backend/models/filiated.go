package models

type Filiated struct {
	IdFiliated int64  `json:"IdFiliated"`
	IdQuiz     int64  `json:"IdQuiz"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
}
