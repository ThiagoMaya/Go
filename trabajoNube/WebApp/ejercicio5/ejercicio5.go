package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
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

			if filepath.Ext(archivo.Name()) == ".jpg" || filepath.Ext(archivo.Name()) == ".png" {

				nombres = append(nombres, archivo.Name())
			}

		}
	}
	fmt.Println("El directorio cuenta con: ", len(nombres), "imágenes")
	fmt.Println("Imágenes en array")
	fmt.Println(nombres)
}
