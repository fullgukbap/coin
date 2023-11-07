package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JJerBum/nomadcoin/utils"
)

var port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}

	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
