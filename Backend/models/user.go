package models

type User struct {
	IdUser     int    `gorm:"column:idusuario;primaryKey" json:"IdUser,omitempty"`
	Name       string `gorm:"column:nome" json:"Name"`
	Email      string `gorm:"column:email" json:"Email"`
	Password   string `gorm:"column:senha" json:"Password"`
	Salt       string `gorm:"column:salt" json:"Salt"`
	TypeUser   string `gorm:"column:tipoUsuario" json:"TypeUser"`
	IdCbo      int    `gorm:"column:cbo_idcbo" json:"IdCbo"`
	IdServices int64  `gorm:"column:idservicos;primaryKey" json:"IdServices"`
	Active     int    `gorm:"column:ativo" json:"Active"`
	Phone      string `gorm:"column:telefone" json:"Phone"`
}

func (p *User) TableName() string {
	return "usuario"
}
