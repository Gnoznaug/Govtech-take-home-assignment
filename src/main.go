package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gnoznaug/src/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterTeacherRoutes(router)
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe("localhost:8080",router))
}

// func goDotEnvVariable(key string) string {

// 	// load .env file
// 	err := godotenv.Load("../.env")
  
// 	if err != nil {
// 	  log.Fatalf("Error loading .env file")
// 	}
  
// 	return os.Getenv(key)
//   }