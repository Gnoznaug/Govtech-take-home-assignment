package routes

import (
	"github.com/gorilla/mux"
	"github.com/gnoznaug/src/controllers"
)

var RegisterTeacherRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/register", controllers.AddTeacher).Methods("POST")
}