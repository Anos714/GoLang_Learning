package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-rest-api/config"
	"todo-rest-api/routes"
	"todo-rest-api/utils"

	"github.com/joho/godotenv"
)


func testHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodGet{
		utils.WriteJson(w,http.StatusMethodNotAllowed,utils.DataBody{
			Success: false,
			Message: "Only GET method allowed",
			Data:nil,
		})
		return
	}

	utils.WriteJson(w,http.StatusOK,utils.DataBody{
		Success:true,
		Message:"pong",
		Data:nil,
		})
}


func main(){

	// loading the env
	err:=godotenv.Load()
	if err!=nil{
		log.Fatalf("Warning: No .env file found, relying on system environment variables: %v\n",err)
	}

	config.ConnectDatabase()

	// a test route
	http.HandleFunc("GET /ping",testHandler)
	// register all routes
	routes.RegisterTodoRoutes()

	port:=os.Getenv("PORT")
	if port==""{
		port="8000"
	}
	fmt.Printf("Server running smoothly on port %s...\n", port)
	err=http.ListenAndServe(":"+port,nil)
	if err!=nil{
		log.Fatalf("Server crash error: %v\n", err)
	}

}
