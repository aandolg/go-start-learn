package main

import (
	"fmt"
	"log"

	"good-work.in/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}