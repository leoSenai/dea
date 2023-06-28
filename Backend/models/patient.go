package models

type Patient struct {
	IdPatient int    `gorm:"column:idpaciente;primaryKey" json:"IdPatient"`
	Name      string `gorm:"column:nome" json:"Name"`
	Email     string `gorm:"column:Email" json:"Email"`
	Cpf       string `gorm:"column:cpf" json:"Cpf"`
	Address   string `gorm:"column:endereco" json:"Address"`
	Phone     string `gorm:"column:telefone" json:"Phone"`
	BornDate  string `gorm:"column:dataNascimento" json:"BornDate"`
	Sex       string `gorm:"column:sexo" json:"Sex"`
	NewBorn   int    `gorm:"column:recemNascido" json:"NewBorn"`
	DadName   string `gorm:"column:nomePai" json:"DadName"`
	MomName   string `gorm:"column:nomeMae" json:"MomName"`
	Cns       string `gorm:"column:CNS" json:"Cns"`
	Cid10     int    `gorm:"column:CID10" json:"Cid10"`
	Active    int    `gorm:"column:ativo" json:"Active"`
	Password  string `gorm:"column:senha" json:"Password"`
	Salt      string `gorm:"column:salt" json:"Salt"`
}

func (p *Patient) TableName() string {
	return "paciente"
}
