package entities

import (
	"errors"
	"github.com/Sina-karimi81/gopham-general/db"
	"github.com/Sina-karimi81/gopham-general/utils"
)

type Staff struct {
	Id        int64
	FirstName string
	LastName  string
	Password  string
	Job       string
	Role      string
}

func NewStaff(id int64, firstName, lastName, job, Role string) *Staff {
	return &Staff{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Job:       job,
		Role:      Role,
	}
}

func (staff *Staff) Save() error {
	query := "INSERT INTO STAFF(FIRST_NAME, LAST_NAME, PASSWORD, JOB, ROLE) VALUES (?, ?, ?, ?, ?)"

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmnt.Close()

	hashedPass, err := utils.EncryptPassword(staff.Password)
	if err != nil {
		return err
	}

	result, err := stmnt.Exec(staff.FirstName, staff.LastName, hashedPass, staff.Job, staff.Role)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	staff.Id = id

	return err
}

func (staff *Staff) ValidateCredentials(pass string) error {
	query := "SELECT ID , PASSWORD FROM STAFF WHERE FIRST_NAME = ? AND LAST_NAME = ?"

	row := db.DB.QueryRow(query, staff.FirstName, staff.LastName)

	var retrievedPassword string
	err := row.Scan(&staff.Password, &retrievedPassword)
	// no user was found
	if err != nil {
		return err
	}

	isValid := utils.CheckPasswordHash(retrievedPassword, pass)

	if !isValid {
		return errors.New("credentials invalid")
	}

	return nil
}
