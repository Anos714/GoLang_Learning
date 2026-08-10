package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	writeJson "todo-rest-api/utils"

	"github.com/joho/godotenv"
)


func testHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodGet{
		writeJson.WriteJson(w,http.StatusMethodNotAllowed,writeJson.DataBody{
			Success: false,
			Message: "Only GET method allowed",
			Data:nil,
		})
		return 
	}
	
	writeJson.WriteJson(w,http.StatusOK,writeJson.DataBody{
		Success:true,
		Message:"Api is healthy or live",
		Data:nil,
		})
	return 
}


func main(){

	// loading the env
	err:=godotenv.Load()
	if err!=nil{
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	// a test route
	http.HandleFunc("/",testHandler)

	port:=os.Getenv("PORT")
	if port==""{
		port="8000"
	}
	err=http.ListenAndServe(":"+port,nil)
	if err!=nil{
		fmt.Printf("Server crash error: %v\n", err)
	}
	fmt.Printf("Server running smoothly on port %s...\n", port)
}
