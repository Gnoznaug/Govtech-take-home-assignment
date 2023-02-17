package models

import (
	"database/sql"
	"github.com/gnoznaug/src/config"
	"github.com/gnoznaug/src/util"
	"fmt"
	// "github.com/gnoznaug/src/errors"
)

var db *sql.DB

type RegisterRequest struct {
	Teacher string `json:"teacher"`
	Students []string `json:"students"`
}

type SuspendStudentRequest struct {
	Student string `json:"student"`
}

func init() {
	config.Connect()
	db = config.GetDB();
}

func RegisterTeacher(teacherEmail string, studentEmails []string) error {
	_, _ = db.Exec(util.GetAddTeacherQuery(teacherEmail));

	_, _ = db.Exec(util.GetAddStudentsQuery(studentEmails))

	_, err := db.Exec(util.GetRegisterStudentsUnderTeacherQuery(teacherEmail, studentEmails))
	if err != nil {
        return fmt.Errorf("One of the students is already registed to this teacher.")
    }
	return nil
}

func SuspendStudent(studentEmail string) error {
	Result, _:= db.Exec(util.GetSuspendStudentQuery(studentEmail));
	rows, _ := Result.RowsAffected()
	if (rows == 0) {
		return fmt.Errorf("The student is already suspended.")
	}
	return nil
}