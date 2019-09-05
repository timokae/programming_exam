package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Kontakt struct {
	Vorname		string
	Nachname	string
}

var kontaktMap = make(map[int]Kontakt)

func main() {
	r := mux.NewRouter()
	kontaktMap[1] = Kontakt{ Vorname: "Carsten", Nachname: "Köhn" }

	r.HandleFunc("/kontakte", getContactsHandler).Methods("GET")
	r.HandleFunc("/kontakte/{id}", getContactHandler).Methods("GET")
	r.HandleFunc("/kontakte", addContactHandler).Methods("POST")
	r.HandleFunc("/kontakte/{id}", updateContactHandler).Methods("PUT")
	r.HandleFunc("/kontakte/{id}", deleteContactHandler).Methods("DELETE")

	log.Printf("Service listening on http://localhost:8080...")
	_ = http.ListenAndServe(":8080", r)
}

func getContactsHandler(writer http.ResponseWriter, _ *http.Request) {
	var kontakte []Kontakt
	for _, kontakt := range kontaktMap {
		kontakte = append(kontakte, kontakt)
	}

	binary, err := json.Marshal(kontakte)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(binary)
}

func getContactHandler(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)
	id, _ := strconv.Atoi(v["id"])

	if _, ok := kontaktMap[id]; !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	binary, err := json.Marshal(kontaktMap[id])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(binary)
}

func addContactHandler(writer http.ResponseWriter, request *http.Request) {
	id := nextId()
	
	var kontakt Kontakt
	err := json.NewDecoder(request.Body).Decode(&kontakt)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	
	kontaktMap[id] = kontakt
	url := request.URL.String()
	writer.Header().Set("Location", fmt.Sprintf("%s/%d", url, id))
	writer.WriteHeader(http.StatusCreated)
}

func updateContactHandler(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)
	id, _ := strconv.Atoi(v["id"])
	
	var kontakt Kontakt
	err := json.NewDecoder(request.Body).Decode(&kontakt)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	kontaktMap[id] = kontakt
	writer.WriteHeader(http.StatusNoContent)
}

func deleteContactHandler(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		log.Print(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, ok := kontaktMap[id]; !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	delete(kontaktMap, id) 
	writer.WriteHeader(http.StatusNoContent)

}

func nextId() int {
	id := 1
	for k := range kontaktMap {
		id = k + 1
	}

	return id
}