package main

import (
	"fmt"
	"log"

	"com.example/greetings"
	// "golang.org/x/example/stringutil"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	// set flag to  0 to hide the time, source file, and line number
	log.SetFlags(0)

	// fmt.Println(stringutil.Reverse("Hello"))

	var message string
	var err error
	message, err = greetings.Hello("Hydra")

	// if an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range messages {
		fmt.Printf("%v: %v\n", k, v)
	}
}
