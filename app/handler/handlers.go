package handler

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type test struct {
	Status string
}

//GreetHandler respojns
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the The Green Dragon !")
	decoder := json.NewDecoder(r.Body)
	var t test
	err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    log.Printf("%+v",t)
}