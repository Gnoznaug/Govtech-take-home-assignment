package controllers

import(
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/gnoznaug/src/models"
	"github.com/gnoznaug/src/errors"
)

func RegisterTeacher(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadAll(r.Body)
	var req models.RegisterRequest
	json.Unmarshal(rawData, &req)
	err = models.RegisterTeacher(req.Teacher, req.Students)
	if (err == nil) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		ErrorHanlder(w,err)
	}
}

func SuspendStudent(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadAll(r.Body)
	var req models.SuspendStudentRequest
	err = json.Unmarshal(rawData, &req)
	if (err != nil) {
		SynxtaxErrorHandler(w, err)
	}
	err = models.SuspendStudent(req.Student)
	if (err == nil) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		ErrorHanlder(w,err)
	}
}

func GetCommonStudents(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	teachers, ok := values["teacher"]
	if (ok) {
		emails, err := models.FindCommonStudents(teachers)
		if (err != nil) {
			ErrorHanlder(w, err)
		} else {
			w.WriteHeader(http.StatusOK)
			res := models.CommonStudentsResponse{Students: emails}
			respJSON, _ := json.Marshal(res)
			w.Write(respJSON)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		res := errors.APIError{Message: "No Teachers chosen."}
		respJSON, _ := json.Marshal(res)
		w.Write(respJSON)
	}
}

func GetRecipients(w http.ResponseWriter, r *http.Request) {
	rawData, err := ioutil.ReadAll(r.Body)
	var req models.RecipientRequest
	err = json.Unmarshal(rawData, &req)
	if (err != nil) {
		SynxtaxErrorHandler(w, err)
	}
	Recipients,err := models.GetRecipients(req.Teacher, req.Notification)
	if (err != nil) {
		ErrorHanlder(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		res := models.RecipientsResponse{Recipients: Recipients}
		respJSON, _ := json.Marshal(res)
		w.Write(respJSON)
	}
}

func ErrorHanlder(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	res := errors.APIError{Message: err.Error()}
	respJSON, _ := json.Marshal(res)
	w.Write(respJSON)
}

func SynxtaxErrorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	res := errors.APIError{Message: "Bad request, please follow the request format."}
	respJSON, _ := json.Marshal(res)
	w.Write(respJSON)
}