package http

import (
	"activities/controller"
	"activities/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var data controller.DataGlobal

func ReturnReadActivities(w http.ResponseWriter, r *http.Request) {

	result := data.ReadActivities()
	jsonData, err := json.Marshal(result)

	if err != nil {
		log.Println("err convert json on http: ", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ReturnCreateActivities(w http.ResponseWriter, r *http.Request) {

	var request model.Activities

	json.NewDecoder(r.Body).Decode(&request)

	data.CreateActivities(request)
	jsonData, err := json.Marshal(request)

	if err != nil {
		log.Println("err convert json on http: ", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ReturnUpdateActivities(w http.ResponseWriter, r *http.Request) {

	var request model.Activities

	_index := mux.Vars(r)
	index := _index["index"]

	dataIndex, err := strconv.Atoi(index)

	if err != nil {
		log.Println("error convert string to int :",err.Error())
	}

	json.NewDecoder(r.Body).Decode(&request)

	data.UpdateActivities(dataIndex,request)
	jsonData, err := json.Marshal(request)

	if err != nil {
		log.Println("err convert json on http: ", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func ReturnDeleteActivities(w http.ResponseWriter, r *http.Request) {

	_index := mux.Vars(r)
	index := _index["index"]

	dataIndex, err := strconv.Atoi(index)

	if err != nil {
		log.Println("error convert string to int :",err.Error())
	}

	data.DeleteActivities(dataIndex)

	if err != nil {
		log.Println("err convert json on http: ", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Index: "+index))

}
