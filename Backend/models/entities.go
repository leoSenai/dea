package models

type Pessoa struct {
	IdPessoa        int64  `gorm:"column:idpessoa" json:"idpessoa"`
	Nome            string `gorm:"column:nome" json:"nome"`
	DataNascimento  string `gorm:"column:dataNascimento" json:"dataNascimento"`
	NumeroDocumento string `gorm:"column:numeroDocumento" json:"numeroDocumento"`
	TipoDocumento   string `gorm:"column:tipoDocumento" json:"tipoDocumento"`
	Senha           string `gorm:"column:senha" json:"senha"`
	Salt            string `gorm:"column:salt" json:"salt"`
}
