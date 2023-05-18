package models

type AnamneseHasAsking struct {
	IdAnamnese int64  `gorm:"column:anamnese_idanamnese" json:"idAnamnese"`
	IdAsking   int64  `gorm:"column:pergunta_idpergunta" json:"idAsking"`
	Awnser     string `gorm:"column:resposta" json:"awnser"`
}

func (p *AnamneseHasAsking) TableName() string {
	return "anamnese_has_pergunta"
}

/*

Business rules:

1 - When deleting an anamneseHasAsking register, the only parameter necessary is the 'IdAnamnese'.
2 - This table should be used to verify if a asking can be deleted or not.
3 - Must be able to get the records passing IdAdamnese or IdAsking as a parameter.

*/
