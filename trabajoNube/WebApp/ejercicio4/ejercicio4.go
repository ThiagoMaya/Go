package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Especifica la ruta del directorio que deseas listar
	directorio := "/home/santiago/carpetaGo" // Cambia esta ruta a la que desees listar

	var nombres []string

	// Leer el contenido del directorio
	archivos, err := ioutil.ReadDir(directorio)
	if err != nil {
		log.Fatal(err)
	}

	// Imprimir los nombres de los archivos
	for _, archivo := range archivos {
		if !archivo.IsDir() {
			fmt.Println(archivo.Name())
			nombres = append(nombres, archivo.Name()+"\n")
		}
	}

	fmt.Print("Nombres guardados en un array \n")
	fmt.Print(nombres)
}
