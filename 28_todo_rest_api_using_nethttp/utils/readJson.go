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

func ReadJson(w http.ResponseWriter,r *http.Request){
	var data RequestBody
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		WriteJson(w,http.StatusBadRequest,DataBody{
			Success: false,
			Message: "Invalid Request Body",
			Data:    nil,
		})
		return
	}

	// validations of request body

	data.Title=strings.TrimSpace(data.Title)

	if data.Title == ""{
		WriteJson(w, http.StatusBadRequest, DataBody{
			Success: false,
			Message: "Title is required",
			Data:    nil,
		})
		return
	}

	response:=DataBody{
		Success:  true,
		Message: "Todo created",
		Data: data,
	}

	// use the helper method writeJson to send a response
	WriteJson(w,http.StatusOK,response)
}
