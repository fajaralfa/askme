package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var userRepo userInterface = repo{}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	all, err := userRepo.All()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(all)

	res, err := json.Marshal(all)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(res))
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := userRepo.Find(vars["id"])
	if err != nil {
		fmt.Println("Error while querying", err)
	}

	userjson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while parsing json", err)
	}

	fmt.Fprint(w, string(userjson))
}
