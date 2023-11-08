package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JJerBum/nomadcoin/utils"
)

var port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	utils.HandleErr(json.NewEncoder(rw).Encode(data))
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
