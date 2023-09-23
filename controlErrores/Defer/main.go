package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hola.txt")
	if err != nil {

		fmt.Println(err)
		return
	}

	defer file.Close() //Se ejecuta antes de terminar la ejecucion del main, cierra el
	//  flujo en cualquier situación antes de dar un return

	_, err = file.Write([]byte("Hola, thiago"))
	if err != nil {

		fmt.Println(err)
		return
	}

}
