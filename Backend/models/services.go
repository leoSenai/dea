package models

type Services struct {
	IdServices int `gorm:"column:idservicos" json:"IdServices"`
	Desc string `gorm:"column:descricao" json:"Desc"`
}

func (p *Services) TableName() string {
	return "servicos"
}