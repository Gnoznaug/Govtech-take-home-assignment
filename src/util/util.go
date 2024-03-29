package util

import (
	"fmt"
	"strings"
	"regexp"
)

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

func GetDoesTeacherExistQuery(email string) string {
	return fmt.Sprintf(`SELECT email FROM teacher WHERE email = "%s";`, email)
}

func GetDoesStudentExistQuery(email string) string {
	return fmt.Sprintf(`SELECT email FROM student WHERE email = "%s";`, email)
}

func GetIsStudentSuspendedQuery(email string) string {
	return fmt.Sprintf(`SELECT email FROM student WHERE email = "%s" AND suspended_status = FALSE;`, email)
}

func ExtractEmails(s string) []string {
	re := regexp.MustCompile(`@[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
    matches := re.FindAllString(s, -1)
	for i, email := range matches {
        matches[i] = email[1:]
    }
    return matches
}
