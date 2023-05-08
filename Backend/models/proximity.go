package models

type Proximity struct {
	IdPatient int    `gorm:"column:paciente_idpaciente;primaryKey" json:"IdPatient"`
	IdPerson  int    `gorm:"column:pessoa_idpessoa" json:"IdPerson"`
	Desc      string `gorm:"column:descricao" json:"Desc"`
}

func (p *Proximity) TableName() string {
	return "proximidade"
}
