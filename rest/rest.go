package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JJerBum/nomadcoin/blockchain"
	"github.com/JJerBum/nomadcoin/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string `json:"message"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See all block",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See a block",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
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
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)

	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
