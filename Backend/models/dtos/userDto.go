package dtos

type UserDTO struct {
	IdUser   int    `gorm:"column:idusuario;primaryKey" json:"IdUser,omitempty"`
	Name     string `gorm:"column:nome" json:"Name"`
	TypeUser string `gorm:"column:tipoUsuario" json:"TypeUser"`
	IdCbo    int    `gorm:"column:cbo_idcbo" json:"IdCbo"`
	Active   int    `gorm:"column:ativo" json:"Active"`
}

func (p *UserDTO) TableName() string {
	return "usuario"
}
