package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DataBody struct{
 Success bool `json:"success"`
 Message string `json:"messsage,omitempty"`
 Data any `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter,statusCode int, data DataBody){

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(statusCode)

	err:=json.NewEncoder(w).Encode(data)
	if err!=nil{
		fmt.Printf("Error encoding the data in json: %v\n",err)
	}
}
