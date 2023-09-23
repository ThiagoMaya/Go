package main

import "fmt"

func saludar(nombre string) {

	//fmt.Println("HelloWorld")
	fmt.Println("Hola", nombre)
}

func sumar(a, b int) int {

	return a + b

}

func main() {
	saludar("Santiago")
	r := sumar(10, 52)
	fmt.Print(r)
}
