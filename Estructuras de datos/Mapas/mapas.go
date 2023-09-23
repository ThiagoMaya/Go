package main

import "fmt"

func main() {

	dias := make(map[int]string)

	fmt.Println(dias)

	//Agregar datos

	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	dias[7] = "SABADUKI"

	fmt.Println(dias)

	delete(dias, 3)

	fmt.Println(dias, len(dias))

	//Nueva map

	estudiantes := make(map[string][]int)

	estudiantes["Santiago"] = []int{13, 15, 19}
	estudiantes["Thiago"] = []int{3, 1, 49}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["Santiago"][0])

}
