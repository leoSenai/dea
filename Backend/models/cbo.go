package models 

type Cbo struct {
	IdCbo int `gorm:"column:idcbo" json:"IdCbo"`
	Code int `gorm:"column:codigo" json:"Code"`
	Desc string `gorm:"column:descricao" json:"Desc"`
}

func (p *Cbo) TableName() string {
	return "cbo"
}