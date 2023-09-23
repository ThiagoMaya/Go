package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
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

	// Imprimir información de cada archivo
	for _, archivo := range archivos {
		if !archivo.IsDir() {
			rutaAbsoluta := filepath.Join(directorio, archivo.Name())
			info, err := os.Stat(rutaAbsoluta)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Nombre:", archivo.Name())
			nombres = append(nombres, archivo.Name()+"\n")
			fmt.Println("Tamaño:", info.Size(), "bytes")
			fmt.Println("Directorio:", rutaAbsoluta)
			fmt.Println("Fecha de última modificación:", info.ModTime().Format(time.RFC3339))
			fmt.Println("Permisos:", info.Mode().String())
			fmt.Println("------------------------")
		}
	}

	fmt.Print("Nombres guardados en un array \n")
	fmt.Print(nombres)
}
