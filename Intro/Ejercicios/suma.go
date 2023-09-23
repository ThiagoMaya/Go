package main

import (
	"fmt"
)

func sumar(a, b int) {

	resultado := a + b

	fmt.Println("El resultado es :", resultado)

}
func division(a, b int) {

	c := a / b
	r := a % b
	fmt.Printf("El cociente es: %d y el residuo es %d", c, r)
}

func precioVenta(igv, pv float64) {
	r := igv * pv

	fmt.Printf("El precio de venta será: %f", r)

}

func main() {

	var opcion, a, b int
	pv := 0.18
	var igv float64

	fmt.Println("Escriba el numero de la funcion que desea realizar\n1.Sumar dos números enteros\n2.Hallar cociente y residuo de dos números\n3.Hallar precio de venta")
	fmt.Scanln(&opcion)

	switch opcion {

	case 1:
		fmt.Printf("Ingrese el primer numero: ")
		fmt.Scanln(&a)

		fmt.Printf("Ingrese el Segundo numero: ")
		fmt.Scanln(&b)

		sumar(a, b)

	case 2:
		fmt.Printf("Ingrese el primer numero: ")
		fmt.Scanln(&a)

		fmt.Printf("Ingrese el Segundo numero: ")
		fmt.Scanln(&b)
		division(a, b)

	case 3:
		fmt.Printf("Ingrese el valor de venta del producto:")
		fmt.Scanln(&igv)
		precioVenta(igv, pv)
	}

}
