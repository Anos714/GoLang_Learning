package main

import (
	"fmt"
	"net/http"
)



func startHandler(w http.ResponseWriter, r *http.Request){
	// val,err:=w.Write([]byte("Hello from go server"))
	// if err != nil {
	// 	fmt.Errorf("Write Failed: %v",err)
	// } else {
	// 	fmt.Println("Wrote",val,"bytes")
	// }

	// we can use w.Write() but fmt.Fprintln() is easier to use
	// syntax: fmt.Fprintln(w, "Hello from go server")
	// fmt.Fprintln() is a convenience function that writes a string to the response writer

	// fmt.Fprintln(w, "\nwe write using fmt.Fprintln()")

	// if you also want to send customs header and status code then always follow this
	// 1 headers -> 2 status code -> 3 response body
	// always in this order otherwise it throws an error

	// set an header
	w.Header().Set("X-Custom-Header", "My server")

	// set a startHandler -> you can directly set the status code or use http.StatusOK
	w.WriteHeader(http.StatusOK)

	// response body
	fmt.Fprintln(w,"Hello from go server")
}


func helloHandler(w http.ResponseWriter, r *http.Request){
 name:=r.URL.Query().Get("name")
 if name=="" {
 	name="Guest"
 }
 fmt.Fprintln(w, "Hello", name)
}


func main(){

	http.HandleFunc("/",startHandler)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/home",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w, "Home Page")
	})
	http.HandleFunc("/about",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w, "About Page")
	})
	http.HandleFunc("/contact",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w, "Contact Page")
	})


	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Errorf("Serve Failed: %v",err)
	} else {
		fmt.Println("Server started on :8080")
	}
}
