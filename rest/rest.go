package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JJerBum/nomadcoin/blockchain"
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

type AddBlockBody struct {
	Message string `json:"message"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks/{id}"),
			Method:      "GET",
			Description: "See a block",
			Payload:     "data:string",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case http.MethodPost:
		// request client example body -> {"message": "my block data"}
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)

	}
}

func Start() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
