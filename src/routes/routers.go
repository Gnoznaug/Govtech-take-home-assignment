package routes

import (
	"github.com/gorilla/mux"
	"github.com/gnoznaug/src/controllers"
)

var RegisterTeacherRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/register", controllers.AddTeacher).Methods("POST")
	router.HandleFunc("/api/suspend", controllers.SuspendStudent).Methods("POST")
	router.HandleFunc("/api/commonstudents", controllers.GetCommonStudents).Methods("GET")
}