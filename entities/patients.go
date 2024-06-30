package entities

type Patient struct {
	Id        int64
	FirstName string
	LastName  string
	IsInsured int
	Diseases  string
}

func NewPatient(id int64, firstName, lastName string, isInsured int, diseases string) *Patient {
	return &Patient{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		IsInsured: isInsured,
		Diseases:  diseases,
	}
}
