package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RequestBody struct{
	Title string `json:"title"`
	Done bool `json:"done"`
}

func ReadJson(w http.ResponseWriter,r *http.Request)(RequestBody, bool){
	var data RequestBody
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data,false
	}

	// validations of request body

	data.Title=strings.TrimSpace(data.Title)

	if data.Title == ""{
		return data,false
	}

	return data,true
}
