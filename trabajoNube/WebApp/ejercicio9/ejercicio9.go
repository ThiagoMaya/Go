package main

import (
	"fmt"
	"os"
)

func main() {
	// Los argumentos se encuentran en os.Args, a partir del segundo elemento.
	// El primer elemento (os.Args[0]) es el nombre del programa en sí.
	// Por lo tanto, los argumentos comienzan desde os.Args[1].
	argumentos := os.Args[1:]

	// Puedes imprimir la lista de argumentos o realizar operaciones en ellos.
	fmt.Println("Número de argumentos:", len(argumentos))
	fmt.Println("Argumentos pasados:")
	for i, arg := range argumentos {
		fmt.Printf("%d: %s\n", i+1, arg)
	}
}
