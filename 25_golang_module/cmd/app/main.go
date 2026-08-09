package main

import (
	greet "go-modules/internal"
)

func main() {
	msg := greet.Greet("rahul sain")
	println(msg)
}
