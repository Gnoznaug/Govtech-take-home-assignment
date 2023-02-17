package errors

import (
    "fmt"
)

type APIError struct {
    Message string `json:"message"`
}

func StudentDoesNotExistError(studentEmail string) error {
    return fmt.Errorf("There is no such student with the email %s.", studentEmail)
}

func TeacherDoesNotExistError(teacherEmail string) error {
    return fmt.Errorf("There is no such teacher with the email %s.", teacherEmail)
}