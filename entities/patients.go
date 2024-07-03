package entities

import (
	"github.com/Sina-karimi81/gopham-general/db"
)

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

func (p *Patient) Save() error {
	query := "INSERT INTO STAFF(FIRST_NAME, LAST_NAME, ISINSURED, DISEASES) VALUES (?, ?, ?, ?)"

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmnt.Close()

	result, err := stmnt.Exec(p.FirstName, p.LastName, p.IsInsured, p.Diseases)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	p.Id = id

	return err
}
