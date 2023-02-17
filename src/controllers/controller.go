package controllers

import(
	"encoding/json"
	"net/http"
	// "fmt"
	// "strconv"
	"io/ioutil"
	"github.com/gnoznaug/src/models"
	"github.com/gnoznaug/src/errors"
)

// var APIError errors.APIError


func AddTeacher(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadAll(r.Body)
	var req models.RegisterRequest
	json.Unmarshal(rawData, &req)
	err = models.RegisterTeacher(req.Teacher, req.Students)
	if (err == nil) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		res := errors.APIError{Message: err.Error()}
		respJSON, _ := json.Marshal(res)
		w.Write(respJSON)
	}
}

func SuspendStudent(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadAll(r.Body)
	var req models.SuspendStudentRequest
	json.Unmarshal(rawData, &req)
	err = models.SuspendStudent(req.Student)
	if (err == nil) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		res := errors.APIError{Message: err.Error()}
		respJSON, _ := json.Marshal(res)
		w.Write(respJSON)
	}
}
