package models

type Asking struct {
	IdAsking int    `gorm:"column:idpergunta;primaryKey" json:"idAsking"`
	Desc     string `gorm:"column:descricao" json:"Desc"`
}

func (p *Asking) TableName() string {
	return "pergunta"
}
