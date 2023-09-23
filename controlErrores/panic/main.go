package main

import "fmt"

func divide(dividendo, divisor int) {

	defer func() { //funcion anónima, se declara sin nombre y se pone parentesis al final para ejecutar automaticamente
		if r := recover(); r != nil {

			fmt.Println(r) // si ocurre algún panic se muestra sin detner la ejecución del programa
		}
	}()
	validarCero(divisor)
	fmt.Println(dividendo / divisor)

}

func validarCero(divisor int) {

	if divisor == 0 {

		panic("No se puede dividir entre 0")
	}
}

func main() {
	divide(100, 25)
	divide(100, 0)

}
