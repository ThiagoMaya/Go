package main

import (
	"fmt"
	"log"

	"github.com/ThiagoMaya/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0) //si se quita muestra el error con fecha y hora y detalles

	names := []string{"Santiago", "Thiago", "Cristiano", "Messi"}
	messages, err := greetings.Hellos(names)

	if err != nil {

		log.Fatal(err)
	}

	/*
		message, err := greetings.Hello("Santiago")

		if err != nil {

			log.Fatal(err)
		}
	*/
	fmt.Println(messages)

}
