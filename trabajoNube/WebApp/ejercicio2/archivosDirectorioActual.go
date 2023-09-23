package main

import (
	"fmt"
	"os"
)

func main() {

	// Obtén el directorio actual
	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	var nombres []string

	// Itera a través de los archivos en el directorio actual
	for _, entry := range dir {
		// Verifica si la entrada es un archivo (no un directorio)
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			nombres = append(nombres, entry.Name()+"\n")
		}
	}

	fmt.Print("Nombres guardados en un array \n")
	fmt.Print(nombres)
}
