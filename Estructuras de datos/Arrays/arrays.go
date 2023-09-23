package main

import (
	"fmt"
)

func main() {

	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 110
	numeros[2] = 450
	fmt.Println(numeros)

	//Arrays con valores

	colores := [...]string{
		"Rojo",
		"Verde",
		"Naranja",
		"Negro",
	}

	fmt.Println(colores, len(colores))

	//Indices indefinidos

	monedas := [...]string{
		0: "Dolar",
		2: "Euros",
		3: "Pesos",
	}

	fmt.Println(monedas, len(monedas))

	//subArray

	submoneda := monedas[0:3]
	fmt.Println(submoneda)

	var matriz [5]int

	fmt.Println(matriz)
}
