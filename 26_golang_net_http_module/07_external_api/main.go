package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type Reaction struct {
	Likes uint `json:"likes"`
	Dislikes uint `json:"dislikes"`
}

type Post struct{
	Id uint `json:"id"`
	Title string `json:"title"`
	Reaction Reaction `json:"reaction"`
	UserId uint `json:"user_id"`
}

type ApiResponse struct {
	Posts []Post `json:"posts,omitempty"`
	ErrorMsg string `json:"error_msg,omitempty"`
}

func writeJson(w http.ResponseWriter,statusCode int,data ApiResponse){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)


	err:=json.NewEncoder(w).Encode(data)
	if err!=nil{
		fmt.Printf("Error occurred: %v\n", err)
	}
}


func externalHandler(w http.ResponseWriter, r *http.Request) {
		url:="https://dummyjson.com/posts?limit=10&select=title,reactions,userId"


		if r.Method!=http.MethodGet{
			writeJson(w, http.StatusMethodNotAllowed, ApiResponse{ErrorMsg: "Method not allowed"})
			return
		}

		res,err:=http.Get(url)
		if err!=nil{
		   writeJson(w, http.StatusInternalServerError, ApiResponse{ErrorMsg: "Failed to fetch posts"})
			return
		}

		if res.StatusCode!=http.StatusOK{
		    writeJson(w, http.StatusInternalServerError, ApiResponse{ErrorMsg: "Failed to fetch posts"})
			return
		}

		defer res.Body.Close()

		bodyBytes,err:=io.ReadAll(res.Body)
		if err!=nil{
			writeJson(w, http.StatusInternalServerError, ApiResponse{ErrorMsg: "Failed to fetch posts"})
			return
		}

		var apiResponse ApiResponse

		err=json.Unmarshal(bodyBytes,&apiResponse)
		if err!=nil{
			writeJson(w, http.StatusInternalServerError, ApiResponse{ErrorMsg: "Failed to fetch posts"})
			return
		}

		// for terminal printing
		for _,val:=range apiResponse.Posts{
			fmt.Printf("%+v\n\n", val)
		}


		// for sending the response back to the client in json
		writeJson(w, http.StatusOK, apiResponse)
}

func main() {
	http.HandleFunc("/",externalHandler)


	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Printf("Error occurred: %v\n", err)
	}

	fmt.Println("Server is running on :8080")
}
