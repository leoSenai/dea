package models

type Question struct {
	IdQuestion int    `gorm:"column:idquestao;primaryKey" json:"IdQuestion"`
	Desc       string `gorm:"column:descricao" json:"Desc"`
	IdQuiz     int    `gorm:"column:idquestionario" json:"IdQuiz"`
}

func (p *Question) TableName() string {
	return "questao"
}
