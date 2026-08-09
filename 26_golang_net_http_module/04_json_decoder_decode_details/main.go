package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type JsonResponse struct {
	Success  string         `json:"status"`
	Message string      `json:"message"`
	Data    any         `json:"data,omitempty"` // Making data optional too is good practice!
}

type RequestBody struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

// 2. Create the Type Constraint Interface
// This acts as a gatekeeper that allows ONLY these two types
type AllowedResponses interface {
	JsonResponse | RequestBody
}


// 3. Create the Generic Function
// [T AllowedResponses] binds the function to your constraint
func writeJson[T AllowedResponses](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method!="POST" {
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// return

		response:=JsonResponse{
			Success:  "failed",
			Message: "Method not allowed",
			Data: nil,
		}

		// use the helper method writeJson to send a response
		writeJson(w,http.StatusMethodNotAllowed,response)
		return
	}

	// you can use this but best practise is to not use this cause go will automatically close the body after the handler returns, but always close the body on the client side http.Get()
	defer r.Body.Close() // Close the body after reading

	var data RequestBody
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		writeJson(w,http.StatusBadRequest,JsonResponse{
			Success: "failed",
			Message: "Invalid Request Body",
			Data:    nil,
		})
		return
	}

	// validations of request body

	data.Name=strings.TrimSpace(data.Name)
	data.Email=strings.TrimSpace(data.Email)

	if data.Name == "" || data.Email == "" {
		writeJson(w, http.StatusBadRequest, JsonResponse{
			Success: "failed",
			Message: "Name  or Email is required",
			Data:    nil,
		})
		return
	}

	response:=JsonResponse{
		Success:  "success",
		Message: "Data received",
		Data: data,
	}

	// use the helper method writeJson to send a response
	writeJson(w,http.StatusOK,response)
}

func main(){
	http.HandleFunc("/test", testHandler)

	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
	fmt.Println("Server started on :8080")
}
