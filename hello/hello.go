package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//set properties of predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time , source file, and line number.
	log.SetPrefix("greetings:");
	log.SetFlags(0);
	message,err := greetings.Hello("")
	//Request a greeting message.
	// if an error was returned, print it to the console and exit the program
	if err!=nil{
		log.Fatal(err);
	}
	fmt.Println(message)
}