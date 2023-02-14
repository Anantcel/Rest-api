package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/getemployees", GetEmployees).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/getemployee/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/updateemployee/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/deleteemployee/{eid}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
