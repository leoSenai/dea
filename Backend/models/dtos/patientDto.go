package dtos

type PatientDTO struct {
	IdPatient int64
	Name      string
	Cpf       string
	Address   string
	Phone     string
	BornDate  string
	Sex       string
	NewBorn   int
	DadName   string
	MomName   string
	Cns       string
	Cid10     int
	Active    int
}
