package models

type Proximity struct {
	IdPatient int64  `gorm:"column:paciente_idpaciente" json:"IdPatient"`
	IdPerson  int64  `gorm:"column:pessoa_idpessoa" json:"IdPerson"`
	Desc      string `gorm:"column:descricao" json:"Desc"`
}

func (p *Proximity) TableName() string {
	return "proximidade"
}
