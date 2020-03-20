package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Driver struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Drivers struct {
	Drivers []Driver
}

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

func GetDriverById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadDrivers()

	var drivers Drivers
	json.Unmarshal(data, &drivers)

	for _, v := range drivers.Drivers {
		if v.Uuid == vars["id"] {
			driver, _ := json.Marshal(v)
			w.Write([]byte(driver))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", ListDrivers)
	r.HandleFunc("/drivers/{id}", GetDriverById)
	http.ListenAndServe(":8081", r)
}
