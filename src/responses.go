package main

import (
	"net/http"
	"encoding/json"
)
/*Handling the response for errors*/
/**
 * respondWithError - Handles the error writring.
 * @w: The object writer.
 * @code: The http code for the type of error.
 * @message: The string containing the specific error.
 */
func respondWithError(w http.ResponseWriter, code int, message string){
	respondWithJSON(w, code, map[string]string{"error": message})
}

/**
 * respondWithJSON - Handles the response by converting info to json format.
 * @w: The object writer.
 * @code: The http code for the request.
 * @payload: Contents of the book object.
 */
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
