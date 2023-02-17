package util

import (
	"fmt"
	"strings"
)

func GetAddTeacherQuery(email string) string {
	return fmt.Sprintf(`INSERT INTO teacher(email) VALUES("%s")`, email)
}

func GetAddStudentsQuery(emails []string) string {
	var EmailValues = GetFormattedStudentEmails(emails)
	var s string = fmt.Sprintf(`INSERT INTO student(email) VALUES%s;`,EmailValues)
	return s
}

func GetFormattedStudentEmails(emails []string) string {
	var EmailValues string
	for _, element := range emails {
		EmailValues += fmt.Sprintf(`("%s"),`, element)
	}
	return strings.TrimSuffix(EmailValues, ",")
}

func GetRegisterStudentsUnderTeacherQuery(teacherEmail string, studentEmails []string) string {
	var StudentEmailValues = GetFormattedStudentEmails(studentEmails)
	return fmt.Sprintf(`INSERT INTO teaching(teacher_id,student_id) SELECT teacher_id, student_id FROM teacher, student WHERE teacher.email = "%s" AND student.email IN %s;`, teacherEmail, StudentEmailValues)
}

func GetSuspendStudentQuery(studentEmail string) string {
	return fmt.Sprintf(`UPDATE student SET suspended_status = 1 WHERE email = "%s";`, studentEmail)
}