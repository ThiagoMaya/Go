package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Manejo de la solicitud HTTP aquí
	fmt.Fprintf(w, "Hola Mundo en el servidor de go!")
}

func main() {
	// Definir un manejador (handler) para las solicitudes HTTP
	http.HandleFunc("/", handler)

	fmt.Println("Servidor corriendo en el puerto 8080")
	// Iniciar el servidor HTTP en el puerto 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}

}
