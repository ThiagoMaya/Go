package main

import "fmt"

func main() {
	nombre := "santiago"
	edad := 18

	mensaje := fmt.Sprintf("Hola, %s  su edad es  %d", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("Nombre: %s , edad %d \n", nombre, edad)

	fmt.Print("ingrese un segundo nombre:")
	fmt.Scanln(&nombre)
	println("Segundo nombre: ", nombre)
}
