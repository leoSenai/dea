package dtos

type PersonDTO struct {
	IdPerson   int64
	Name       string
	BornDate   string
	DocNumber  string
	DocType    string
	Password   string
	Salt       string
	IdPatient  int64
	DescPerson string
	Email      string
}
