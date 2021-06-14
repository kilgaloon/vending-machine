package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSONResponse write to response writter
func JSONResponse(w http.ResponseWriter, j interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}

	if status != 0 {
		w.WriteHeader(status)
	}

	w.Write(response)
}
