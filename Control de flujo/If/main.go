package main

import "fmt"

func main() {

	//Se definen las variables antes del ;  que marca el inicio de la condicion

	//DEFINICION DE VARIABLES EN EL IF

	if nombre, edad := "rubencho", 18; nombre == "thiago" {

		fmt.Println("Hola, ", nombre)
	} else {

		fmt.Printf("Nombre : %s - Edad %d\n", nombre, edad)
	}

	users := make(map[string]string)

	//Cada clave retorna dos valores: el valor y un boolean

	users["Alex"] = "alex@mail.com"
	users["Santiago"] = "santiago@mail.com"

	if correo, ok := users["Alex"]; ok {

		fmt.Println(correo)
	} else {

		fmt.Println("No se obtuvo el correo")
	}

}
