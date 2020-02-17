// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//User struct to hold user info
type User struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	Bio  string `json:"Bio"`
}

//Users users list
var Users []User

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Users Directory! \n ---- APIs available ---- \n /users GET\n /user POST \n /user/{id} DELETE \n /user{id} GET \n")
	fmt.Println("homePage")
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Return all Users")
	json.NewEncoder(w).Encode(Users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, user := range Users {
		if user.ID == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	// update our global Articles array to include
	// our new Article
	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, user := range Users {
		if user.ID == id {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/dir/", homePage)
	myRouter.HandleFunc("/dir/users", returnAllUsers)
	myRouter.HandleFunc("/dir/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/dir/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/dir/user/{id}", returnSingleUser)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Users = []User{
		User{ID: "1", Name: "Hrishi", Bio: "Software Engineer"},
		User{ID: "2", Name: "Ashu", Bio: "Network Engineer"},
	}
	handleRequests()
}
