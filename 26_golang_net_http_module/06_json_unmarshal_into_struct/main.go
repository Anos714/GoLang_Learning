package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSON Unmarshaling in Go means converting raw JSON data (like a string, bytes, or a network request payload) into a Go data structure, such as a struct, map, or slice.

// To do this, Go provides the json.Unmarshal() function inside the standard encoding/json package.


type User struct {
	FirstName string `json:"firstName"`
	Age        int    `json:"age"`

}

// API response wrapper
type ApiResponse struct {
	Users []User `json:"users"`
}

func main(){
		url:="https://dummyjson.com/users?limit=5&select=firstName,age"

		res,err:=http.Get(url)
		if err!=nil{
			fmt.Println("Error fetching from url: %v",err)
		}
		defer res.Body.Close()

		if res.StatusCode!=http.StatusOK{
			fmt.Printf("Server returned an error status: %d\n", res.StatusCode)
			return
		}


		bodyBytes,err:=io.ReadAll(res.Body)
		if err!=nil{
			fmt.Println("Error reading response body:", err)
			return
		}

		var apiRes ApiResponse

		// unmarshal the JSON response into the ApiResponse struct
		error:=json.Unmarshal(bodyBytes,&apiRes)
		if error!=nil{
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		for _,val:=range apiRes.Users{
			fmt.Printf("%+v\n", val)
		}
}
