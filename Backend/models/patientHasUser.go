package models

type PatientHasUser struct {
	IdPatient int64 `gorm:"column:paciente_idpaciente" json:"IdPatient"`
	IdUser    int64 `gorm:"column:usuario_idusuario" json:"IdUser"`
}

func (p *PatientHasUser) TableName() string {
	return "paciente_has_usuario"
}
