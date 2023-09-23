package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

func buscarArchivo(nombreArchivo string, directorioInicial string) (string, error) {
	var rutaEncontrada string

	err := filepath.Walk(directorioInicial, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == nombreArchivo {
			rutaEncontrada = ruta
			return nil
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if rutaEncontrada == "" {
		return "", fmt.Errorf("Archivo no encontrado: %s", nombreArchivo)
	}

	return rutaEncontrada, nil
}

func main() {

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

	indiceAleatorio := rand.Intn(len(nombres))
	// Acceder al elemento aleatorio
	elementoAleatorio := nombres[indiceAleatorio]

	fmt.Println("El archivo escogido al azar es:", elementoAleatorio)

	rutaImagen, err := buscarArchivo(elementoAleatorio, directorio)
	if err != nil {
		log.Fatal(err)
	}

	contenidoImagen, err := ioutil.ReadFile(rutaImagen)
	if err != nil {
		fmt.Println("Error al leer la imagen:", err)
		return
	}

	// Codificar el contenido de la imagen en base64
	imagenBase64 := base64.StdEncoding.EncodeToString(contenidoImagen)

	archivo, err := os.Create("ImagenCodificada.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(imagenBase64)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	// Imprimir la imagen codificada en base64
	fmt.Println("Imagen codificada en base64.")

}
