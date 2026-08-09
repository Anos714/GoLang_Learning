package main

import (
	"fmt"
	"io"
	"net/http"
)




func main() {
	url:="https://dummyjson.com/users?limit=5&select=firstName,age"

	// 1. Send the HTTP GET request
	res,err:=http.Get(url)
	if err!=nil{
		fmt.Printf("Network error fetching URL: %v\n", err)
				return
	}

	// 2. CRITICAL: Always close the body to prevent memory leaks
	// defer ensures it runs at the very end of the main function
	defer res.Body.Close()

	// 3. Check the HTTP status code
	if res.StatusCode!=http.StatusOK{
		fmt.Printf("Server returned an error status: %d\n", res.StatusCode)
				return
	}

	// 4. Read the response body bytes
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// it gives you the raw bytes of the response body
	fmt.Print("Raw response: ",bodyBytes)

	// 5. Convert the response body bytes to a string
	bodyString:=string(bodyBytes)
	fmt.Println("Response Data:")
	fmt.Println(bodyString)

}
