package models

import "time"

type ProximityHasQuiz struct {
	ProximityIdPatient int64     `gorm:"column:proximidade_paciente_idpaciente" json:"ProximityIdPatient"`
	ProximityIdPerson  int64     `gorm:"column:proximidade_pessoa_idpessoa" json:"ProximityIdPerson"`
	IdQuiz             int64     `gorm:"column:questionario_idquestionario" json:"IdQuiz"`
	Finished           string    `gorm:"finalizado" json:"Finished"`
	Answers            string    `gorm:"column:respostas" json:"Answers"`
	AnsweredIn         time.Time `gorm:"column:respondido_em" json:"AnsweredIn"`
}

func (p *ProximityHasQuiz) TableName() string {
	return "proximidade_has_questionario"
}
