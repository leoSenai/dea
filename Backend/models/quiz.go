package models

type Quiz struct {
	IdQuiz   int64  `gorm:"column:idquestionario;primaryKey" json:"IdQuiz" `
	Name     string `gorm:"column:nome" json:"Name"`
	Created  string `gorm:"column:criado" json:"Created"`
	Updated  string `gorm:"column:atualizado" json:"Updated"`
	Interval string `gorm:"column:intervalo" json:"Interval"`
}

func (p *Quiz) TableName() string {
	return "questionario"
}
