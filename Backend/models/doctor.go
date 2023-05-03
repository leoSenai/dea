package models

type Doctor struct {
	IdDoctor int `gorm:"column:idmedico" json:"IdDoctor"`
	Crm string `gorm:"column:crm" json:"Crm"`
	Name string `gorm:"column:nome" json:"Name"`
	IdCbo int `gorm:"column:cbo_idcbo" json:"IdCbo"`
}

func (p *Doctor) TableName() string {
	return "medico"
}