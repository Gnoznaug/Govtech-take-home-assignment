package models

import (
	"database/sql"
	"github.com/gnoznaug/src/config"
	"github.com/gnoznaug/src/util"
	"fmt"
	// "github.com/gnoznaug/src/errors"
	"log"
)

var db *sql.DB

type RegisterRequest struct {
	Teacher string `json:"teacher"`
	Students []string `json:"students"`
}

type SuspendStudentRequest struct {
	Student string `json:"student"`
}

type CommonStudentsResponse struct {
	Students []string `json:"students"`
}

func init() {
	config.Connect()
	db = config.GetDB();
}

func RegisterTeacher(teacherEmail string, studentEmails []string) error {
	_, _ = db.Exec(util.GetAddTeacherQuery(teacherEmail));

	for _,element := range studentEmails {
		_,_ = db.Exec(util.GetAddStudentQuery(element))
	}
	_, err := db.Exec(util.GetRegisterStudentsUnderTeacherQuery(teacherEmail, studentEmails))
	if err != nil {
        return fmt.Errorf("One of the students is already registed to this teacher.")
    }
	return nil
}

func FindCommonStudents(teacherEmails []string) ([]string,error) {
	// maybe query if the teachers exist or not first
	rows, _ := db.Query(util.GetCommonStudentsQuery(teacherEmails));
	defer rows.Close()
	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			log.Fatal(err)
		}
		emails = append(emails, email)
	}
	return emails, nil
}

func SuspendStudent(studentEmail string) error {
	Result, _:= db.Exec(util.GetSuspendStudentQuery(studentEmail));
	rows, _ := Result.RowsAffected()
	if (rows == 0) {
		return fmt.Errorf("The student is already suspended.")
	}
	return nil
}