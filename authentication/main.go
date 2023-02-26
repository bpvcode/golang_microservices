package main

import (
	"fmt"

	"github.com/bpvcode/golang_microservices/authentication/initializers/db"
	"github.com/bpvcode/golang_microservices/authentication/initializers/environment"
	"github.com/bpvcode/golang_microservices/authentication/initializers/http"
)

// Init function will set environment variables and test database connection
func init() {
	environment.LoadEnvVariables()
	db := db.GetDB()
	fmt.Print(db.Statement.Vars...)
	// add db migrations here if needed
}

func main() {
	http.New()
	// TODO: ADD router.Register() here
	http.Listen()
}
