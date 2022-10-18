package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

type Student struct {
	Nama string
	Data []string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/purwita", satu).Methods("GET")
	r.HandleFunc("/purwita/{id}", dua).Methods("GET")

	log.Fatal(http.ListenAndServe(PORT, r))
}

func satu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	if r.Method == "GET" {
		d1 := strconv.Itoa(123)
		d2 := strconv.FormatBool(true)

		var Student = Student{
			Nama: "Purwitasari",
			Data: []string{d1, "memasak", d2},
		}

		json.NewEncoder(w).Encode(Student)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func dua(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("Invalid")
	}

	if r.Method == "GET" {
		d1 := strconv.Itoa(123)
		d2 := strconv.FormatBool(true)

		var Student = Student{
			Nama: "Purwitasari",
			Data: []string{d1, "memasak", d2},
		}
		if ID == 1 {
			msg := "Hallo Saya Purwita"
			json.NewEncoder(w).Encode(msg)
			return
		} else if ID == 2 {
			json.NewEncoder(w).Encode(Student.Data)
			return
		} else {
			json.NewEncoder(w).Encode(Student)
			return
		}
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)

}
