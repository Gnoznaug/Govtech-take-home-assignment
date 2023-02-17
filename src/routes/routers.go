package routes

import (
	"github.com/gorilla/mux"
	"github.com/gnoznaug/src/controllers"
)

var RegisterTeacherRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/register", controllers.RegisterTeacher).Methods("POST")
	router.HandleFunc("/api/commonstudents", controllers.GetCommonStudents).Methods("GET")
	router.HandleFunc("/api/suspend", controllers.SuspendStudent).Methods("POST")
	router.HandleFunc("/api/retrievefornotifications", controllers.GetRecipients).Methods("POST")
}