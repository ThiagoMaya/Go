package main

import (
	"errors"
	"fmt"
)

func divide(dividendo, divisor int) (int, error) {

	if divisor == 0 {

		return 0, errors.New("No se puede dividir entre 0")

	}
	return dividendo / divisor, nil
}

func main() {

	result, err := divide(10, 6)

	if err != nil {

		fmt.Println("Error", err)
		return
	}
	fmt.Println("Resultado:", result)
	/*
		str := "123a"

		num, err := strconv.Atoi(str) //El strconv devuelve el numero en la variable y un error

		if err != nil { //Si la variable error devuelve nil es porque no hay error
			fmt.Println("error:", err) // De lo contrario se muestra el error recibido
			return
		}

		fmt.Println("Numero: ", num)
	*/
}
