package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// json response struct
type UserResponse struct {
	Status string `json "status"`
	Message string `json "message"`
	DateTime string `json "datetime"`
}

// controllers
func jsonEncoderhandler(w http.ResponseWriter, r *http.Request) {

	// setting headers
	w.Header().Set("Content-Type", "application/json")

	// setting status code
	w.WriteHeader(http.StatusOK)

	// first set json struct
	UserResponse := UserResponse{
		Status:  "success",
		Message: "Hello from json encoder",
		DateTime: time.Now().UTC().String(),
	}
	err := json.NewEncoder(w).Encode(UserResponse)
	if err != nil {
		fmt.Printf("Failed to encode: %v", err)
	}
}


func main(){


	// routes
	http.HandleFunc("/",jsonEncoderhandler)



	// server starts
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		fmt.Errorf("Server failed: %v",err)
	}
	fmt.Println("Server started on :8080")
}
