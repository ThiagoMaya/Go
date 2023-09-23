package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtener el nombre de host
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error al obtener el nombre de host:", err)
		return
	}

	fmt.Println("Nombre de host:", hostname)
}
