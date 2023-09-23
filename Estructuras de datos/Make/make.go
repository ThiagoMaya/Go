package main

import "fmt"

func main() {

	numeros := make([]int, 3, 3)

	numeros[0] = 100
	numeros[1] = 150
	numeros[2] = 190

	numeros = append(numeros, 745)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
