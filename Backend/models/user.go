package models

type User struct {
	IdUser int `gorm:"column:idusuario" json:"IdUser"`
	Name string `gorm:"column:nome" json:"Name"`
	Password string `gorm:"column:senha" json:"Password"`
	Salt string `gorm:"column:salt" json:"Salt"`
	TypeUser byte `gorm:"column:tipousuario" json:"TypeUser"`
	IdCbo int `gorm:"column:cbo_idcbo" json:"IdCbo"`
	Active byte `gorm:"column:ativo" json:"Active"`
}

func (p *User) TableName() string {
	return "usuario"
}