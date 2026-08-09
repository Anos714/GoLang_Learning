package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		io.WriteString(w,"Hello World!\n")
	}

	http.HandleFunc("/",helloHandler)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))

}
