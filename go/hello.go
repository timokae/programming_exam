package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var contacts = make(map[int]string)

func getContact(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)
	id, _ := strconv.Atoi(v["id"])
	writer.Write([]byte(contacts[id]))
}

func main() {
	contacts[0] = "Anton"
	contacts[1] = "Berta"
	contacts[2] = "Ceasar"

	r := mux.NewRouter()
	r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	http.ListenAndServe(":8080", r)
}