package models

type Doctor struct {
	IdDoctor int64  `gorm:"column:idmedico;primaryKey" json:"IdDoctor"`
	Crm      string `gorm:"column:crm" json:"Crm"`
	Name     string `gorm:"column:nome" json:"Name"`
	IdCbo    int64  `gorm:"column:cbo_idcbo" json:"IdCbo"`
}

func (p *Doctor) TableName() string {
	return "medico"
}
