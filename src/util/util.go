package util

import (
	"fmt"
	"strings"
)

func GetAddTeacherQuery(email string) string {
	return fmt.Sprintf(`INSERT INTO teacher(email) VALUES("%s")`, email)
}

func GetAddStudentQuery(emails string) string {
	var s string = fmt.Sprintf(`INSERT INTO student(email) VALUES("%s");`,emails)
	return s
}

func GetFormattedEmails(emails []string) string {
	var EmailValues string = "("
	for _, element := range emails {
		EmailValues += fmt.Sprintf(`"%s",`, element)
	}
	return strings.TrimSuffix(EmailValues, ",") + ")"
}

func GetRegisterStudentsUnderTeacherQuery(teacherEmail string, studentEmails []string) string {
	var StudentEmailValues = GetFormattedEmails(studentEmails)
	return fmt.Sprintf(`INSERT INTO teaching(teacher_id,student_id) SELECT teacher_id, student_id FROM teacher, student WHERE teacher.email = "%s" AND student.email IN %s;`, teacherEmail, StudentEmailValues)
}

func GetCommonStudentsQuery(teacherEmails []string) string {
	return fmt.Sprintf(`SELECT DISTINCT student.email FROM student INNER JOIN teaching ON student.student_id = teaching.student_id
		 INNER JOIN teacher ON teaching.teacher_id = teacher.teacher_id WHERE teacher.email IN %s;`, GetFormattedEmails(teacherEmails))
}

func GetSuspendStudentQuery(studentEmail string) string {
	return fmt.Sprintf(`UPDATE student SET suspended_status = 1 WHERE email = "%s";`, studentEmail)
}