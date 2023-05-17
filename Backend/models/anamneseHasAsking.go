package models

type AnamneseHasAsking struct {
	IdAnamnese int64  `gorm:"column:anamnese_idanamnese" json:"idAnamnese"`
	IdAsking   int64  `gorm:"column:pergunta_idpergunta" json:"idAsking"`
	Awnser     string `gorm:"column:resposta" json:"awnser"`
}

func (p *AnamneseHasAsking) TableName() string {
	return "anamnese_has_pergunta"
}
