/*
q1. Create a handler which accepts any kind of json and prints it
    check if client is still connected or not
    if client is connected then return json processed otherwise just move on
	r.body
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    UserData `json:"data"`
}

func main() {
	http.HandleFunc("/json-parse", JsonParse)
	fmt.Println("Server started at :8086")
	http.ListenAndServe(":8086", nil)
}

func JsonParse(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var userData UserData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Parsed JSON: %+v\n", userData)

	// select {
	// case <-r.Context().Done():
	// 	fmt.Println("Client disconnected. Halting further processing.")
	// 	return
	// default:
		
	// }

	response := Response{
		Status:  "success",
		Message: "Processed successfully",
		Data:    userData,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
