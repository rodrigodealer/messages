package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/messages/handlers"
	"github.com/rodrigodealer/messages/mysql"
)

func main() {

	log.SetOutput(os.Stdout)
	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", Server())
	if err != nil {
		log.Panic("Something is wrong : " + err.Error())
	}
}

func Server() http.Handler {
	mysql := &mysql.MySQLConnection{}
	mysql.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler(mysql)).Name("/healthcheck").Methods("GET")
	return r
}
