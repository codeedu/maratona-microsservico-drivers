package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func loadDrivers() []byte {

	jsonFile, err := os.Open("drivers.json")
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func ListDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := loadDrivers()
	w.Write([]byte(drivers))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", ListDrivers)
	http.ListenAndServe(":8081", r)
}
