package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Sharad")
	fmt.Println(message)
}