package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//hello devuelve un saludo

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("El nombre está vacio")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil

}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {

		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message

	}

	return messages, nil

}

//devuelve varios formatos de saludo de manera aleatoria

func randomFormat() string {

	formats := []string{

		"Hola, %v, puto",
		"Hola, %v, sapo",
		"Qué bueno verte, %v",
		"De nuevo esta gonorrea de %v",
	}

	return formats[rand.Intn(len(formats))]
}
