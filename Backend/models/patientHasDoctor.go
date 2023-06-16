package models

type PatientHasDoctor struct {
	IdPatient int64 `gorm:"column:paciente_idpaciente" json:"IdPatient"`
	IdDoctor  int64 `gorm:"column:medico_idmedico" json:"IdDoctor"`
}

func (p *PatientHasDoctor) TableName() string {
	return "paciente_has_medico"
}
