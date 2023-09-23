package main

import (
	"fmt"
	"math/rand"
)

func main() {

	jugar()

}

func jugar() {

	numAleatorio := rand.Intn(100)
	var numeroIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {

		intentos++

		fmt.Printf("Ingresa un numero (intentos restantes: %d)", maxIntentos-intentos)
		fmt.Scanln(&numeroIngresado)

		if numeroIngresado == numAleatorio {

			fmt.Println("Felicidades, has acertado")
			retry()
			return
		} else if numeroIngresado < numAleatorio {
			fmt.Println("El número a adivinar es mayor")
		} else if numeroIngresado > numAleatorio {
			fmt.Println("El número a adivinar es menor")
		}

	}

	fmt.Printf("Numero de intentos superado, el numero a adivinar era : %d", numAleatorio)
	retry()
}

func retry() {

	var eleccion string
	fmt.Println("Desea jugar nuevamente? (s/n)")
	fmt.Scanln(&eleccion)

	switch eleccion {

	case "s":
		jugar()

	case "n":
		fmt.Println("Gracias por jugar sapo")
	}

}
