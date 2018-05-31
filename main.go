package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// The TODO Type (more like an object)
type TODO struct {
    Name string   `json:"Name,omitempty"`
    Completed bool   `json:"Completed,omitempty"`
    Due  time.Time `json:"Due,omitempty"`
}

var todos []TODO

// Display all from the todos var
func Gettodos(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(todos)
}

// create a new item
func CreateTODO(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var TODO TODO
    _ = json.NewDecoder(r.Body).Decode(&TODO)
    TODO.Name = params["Name"]
    todos = append(todos, TODO)
    json.NewEncoder(w).Encode(todos)
}

// Delete an item
func DeleteTODO(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range todos {
        if item.Name == params["Name"] {
            todos = append(todos[:index], todos[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(todos)
    }
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    todos = append(todos, TODO{ID: "1", Name: "Complete Assignment: Machine Learning", Completed: "Yes", Due: "31/5/2018"})
    todos = append(todos, TODO{ID: "2", Name: "Gym registration", Completed: "No", Due: "30/6/2018"})
    router.HandleFunc("/todos", Gettodos).Methods("GET")
    router.HandleFunc("/todos/{id}", GetTODO).Methods("GET")
    router.HandleFunc("/todos/{id}", CreateTODO).Methods("POST")
    router.HandleFunc("/todos/{id}", DeleteTODO).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
