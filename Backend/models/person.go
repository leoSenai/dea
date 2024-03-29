package models

type Person struct {
	IdPerson  int64  `gorm:"primaryKey;autoIncrement;column:idpessoa" json:"IdPerson"`
	Name      string `gorm:"column:nome" json:"Name"`
	BornDate  string `gorm:"column:dataNascimento" json:"BornDate"`
	DocNumber string `gorm:"column:numeroDocumento" json:"DocNumber"`
	DocType   string `gorm:"column:tipoDocumento" json:"DocType"`
	Email     string `gorm:"column:email" json:"Email"`
	Password  string `gorm:"column:senha" json:"Password"`
	Salt      string `gorm:"column:salt" json:"Salt"`
	Phone     string `gorm:"column:telefone" json:"Phone"`
}

func (p *Person) TableName() string {
	return "pessoa"
}
