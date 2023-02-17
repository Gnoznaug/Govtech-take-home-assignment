package util

import (
    "testing"
)

func TestGetFormattedEmails(t *testing.T) {
    testCases := []struct {
        name     string
        emails   []string
        expected string
    }{
        {
            name:     "Empty input",
            emails:   []string{},
            expected: "()",
        },
        {
            name:     "Single email",
            emails:   []string{"test@example.com"},
            expected: `("test@example.com")`,
        },
        {
            name:     "Multiple emails",
            emails:   []string{"test1@example.com", "test2@example.com", "test3@example.com"},
            expected: `("test1@example.com","test2@example.com","test3@example.com")`,
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := GetFormattedEmails(tc.emails)
            if result != tc.expected {
                t.Errorf("Expected %v, but got %v", tc.expected, result)
            }
        })
    }
}

func TestExtractEmails(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected []string
    }{
        {
            name:     "No emails",
            input:    "Hello world",
            expected: []string{},
        },
        {
            name:     "Single email",
            input:    "@test@example.com",
            expected: []string{"test@example.com"},
        },
		{
            name:     "Email without mention",
            input:    "test@example.com",
            expected: []string{},
        },
        {
            name:     "Multiple emails",
            input:    "@test1@example.com @test2@example.com @test3@example.com",
            expected: []string{"test1@example.com", "test2@example.com", "test3@example.com"},
        },
		{
            name:     "Message in front of email",
            input:    "Hello! @test1@example.com",
            expected: []string{"test1@example.com"},
        },
		{
            name:     "Message behind email",
            input:    "@test1@example.com Hello!",
            expected: []string{"test1@example.com"},
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := ExtractEmails(tc.input)
            if len(result) != len(tc.expected) {
                t.Errorf("Expected %v, but got %v", tc.expected, result)
                return
            }
            
            for i, email := range result {
                if email != tc.expected[i] {
                    t.Errorf("Expected %v, but got %v", tc.expected[i], email)
                    return
                }
            }
        })
    }
}

func TestGetSuspendStudentQuery(t *testing.T) {
    studentEmail := "studenta@gmail.com"
    expectedQuery := `UPDATE student SET suspended_status = 1 WHERE email = "studenta@gmail.com";`
    
    query := GetSuspendStudentQuery(studentEmail)
    if query != expectedQuery {
        t.Errorf("Expected query '%s', but got '%s'", expectedQuery, query)
    }
}

func TestGetDoesTeacherExistQuery(t *testing.T) {
    email := "teacher@example.com"
    expectedQuery := `SELECT email FROM teacher WHERE email = "teacher@example.com";`
    
    query := GetDoesTeacherExistQuery(email)
    if query != expectedQuery {
        t.Errorf("Expected query '%s', but got '%s'", expectedQuery, query)
    }
}

func TestGetDoesStudentExistQuery(t *testing.T) {
    email := "student@example.com"
    expectedQuery := `SELECT email FROM student WHERE email = "student@example.com";`
    
    query := GetDoesStudentExistQuery(email)
    if query != expectedQuery {
        t.Errorf("Expected query '%s', but got '%s'", expectedQuery, query)
    }
}

func TestGetIsStudentSuspendedQuery(t *testing.T) {
    email := "student@example.com"
    expectedQuery := `SELECT email FROM student WHERE email = "student@example.com" AND suspended_status = FALSE;`
    
    query := GetIsStudentSuspendedQuery(email)
    if query != expectedQuery {
        t.Errorf("Expected query '%s', but got '%s'", expectedQuery, query)
    }
}