package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
)

func main() {
	// Especifica la ruta del directorio que deseas listar
	directorio := "/home/santiago/carpetaGo" // Cambia esta ruta a la que desees listar

	// Leer el contenido del directorio
	archivos, err := ioutil.ReadDir(directorio)
	if err != nil {
		log.Fatal(err)
	}

	var nombres []string

	// Imprimir los nombres de los archivos
	for _, archivo := range archivos {
		if !archivo.IsDir() {

			nombres = append(nombres, archivo.Name())
		}
	}

	indiceAleatorio := rand.Intn(len(nombres))
	// Acceder al elemento aleatorio
	elementoAleatorio := nombres[indiceAleatorio]

	fmt.Println("El archivo escogido al azar es:", elementoAleatorio)

}
