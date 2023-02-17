package models

import (
	"database/sql"
	"github.com/gnoznaug/src/config"
	"github.com/gnoznaug/src/util"
	"fmt"
	"github.com/gnoznaug/src/errors"
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

type RecipientRequest struct {
	Teacher string `json:"teacher"`
	Notification string `json:"notification"`
}

type RecipientsResponse struct {
	Recipients []string `json:"recipients"`
}

func init() {
	config.Connect()
	db = config.GetDB();
}

func RegisterTeacher(teacherEmail string, studentEmails []string) error {
	if (!TeacherExists(teacherEmail)) {
		return errors.TeacherDoesNotExistError(teacherEmail)
	}
	for _,element := range studentEmails {
		if (!StudentExists(element)) {
			return errors.StudentDoesNotExistError(element)
		}
	}
	_, err := db.Exec(util.GetRegisterStudentsUnderTeacherQuery(teacherEmail, studentEmails))
	if err != nil {
        return fmt.Errorf("One of the students is already registed to this teacher.")
    }
	return nil
}

func FindCommonStudents(teacherEmails []string) ([]string,error) {
	for _, element := range teacherEmails {
		if (!TeacherExists(element)) {
			return nil, errors.TeacherDoesNotExistError(element)
		}
	}
	rows, _ := db.Query(util.GetCommonStudentsQuery(teacherEmails))
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
	if (!StudentExists(studentEmail)) {
		return errors.StudentDoesNotExistError(studentEmail)
	}
	Result, _:= db.Exec(util.GetSuspendStudentQuery(studentEmail))
	rows, _ := Result.RowsAffected()
	if (rows == 0) {
		return fmt.Errorf("The student is already suspended.")
	}
	return nil
}

func GetRecipients(teacherEmail string, notification string) ([]string, error) {
	if (!TeacherExists(teacherEmail)) {
		return nil, errors.TeacherDoesNotExistError(teacherEmail)
	}
	explicitRecipients := util.ExtractEmails(notification)
	teacher := []string{teacherEmail}
	implicitRecipients,_ := FindCommonStudents(teacher)
	s := map[string]bool{}
	for _, element := range explicitRecipients {
		if (!StudentExists(element)) {
			return nil, errors.StudentDoesNotExistError(element)
		}
		s[element] = true
	}
	for _, element := range implicitRecipients {
		s[element] = true
	}
	var recipients []string
	for key,_ := range s {
		if (StudentIsNotSuspended(key)) {
			recipients = append(recipients, key)
		}
	}
	return recipients,nil
}

func TeacherExists(email string) bool {
	rows, _ := db.Query(util.GetDoesTeacherExistQuery(email))
	return rows.Next()
}

func StudentExists(email string) bool {
	rows, _ := db.Query(util.GetDoesStudentExistQuery(email))
	return rows.Next()
}

func StudentIsNotSuspended(email string) bool {
	rows, _ := db.Query(util.GetIsStudentSuspendedQuery(email))
	return rows.Next()
}