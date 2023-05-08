package models

type Anamnese struct {
	IdAnamnese int    `gorm:"column:idanamnese;primaryKey" json:"IdAnamnese"`
	IdPatient  int    `gorm:"column:idpaciente" json:"IdPatient"`
	IdUser     int    `gorm:"column:idusuario" json:"IdUser"`
	Notes      string `gorm:"column:anotacoes" json:"Notes"`
	Indicative int    `gorm:"column:indicativo" json:"Indicative"`
}

func (p *Anamnese) TableName() string {
	return "anamnese"
}
