package main

import (
	httpActivities "activities/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("HELLO I AM VERY VERY BORED"))
	})
	r.HandleFunc("/activities", httpActivities.ReturnReadActivities).Methods("GET")              // READ
	r.HandleFunc("/activities", httpActivities.ReturnCreateActivities).Methods("POST")   // CREATE
	r.HandleFunc("/activities/{index}", httpActivities.ReturnUpdateActivities).Methods("PATCH")          // UPDATE
	r.HandleFunc("/activities/{index}", httpActivities.ReturnDeleteActivities).Methods("DELETE") // DELETE

	fmt.Println("Run Port:9090")
	log.Fatal(http.ListenAndServe(":9090", r))

}
