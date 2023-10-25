package models

type PatientHasQuiz struct {
	IdPatient  int64  `gorm:"column:paciente_idpaciente" json:"IdPatient"`
	IdQuiz     int64  `gorm:"column:questionario_idquestionario" json:"IdQuiz"`
	Finished   int64  `gorm:"finalizado" json:"Finished"`
	Answers    string `gorm:"column:respostas" json:"Answers"`
	AnsweredIn string `gorm:"column:respondido_em" json:"AnsweredIn"`
}

func (p *PatientHasQuiz) TableName() string {
	return "paciente_has_questionario"
}
