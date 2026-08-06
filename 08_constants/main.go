package main

import "fmt"

func main() {
	const appName="GoLang Learning"
	fmt.Println(appName)

	// typed constants

	const maxUpload int=25;
	const discount float64=33.33;
	const maxStorage string ="20gb"
	fmt.Println(maxUpload, discount, maxStorage)
}