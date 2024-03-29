package dtos

type PersonDTO struct {
	IdPerson   int64
	Name       string
	BornDate   string
	DocNumber  string
	DocType    string
	Password   string
	Salt       string
	Phone      string
	IdPatient  int64
	DescPerson string
	Email      string
}

type PersonResetPasswordDTO struct {
	IdPerson int64
	Email    string
}
