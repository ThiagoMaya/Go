package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListaTareas struct {
	tareas []Tarea
}

// agregar tareas
func (l *ListaTareas) agregarTarea(t Tarea) {

	l.tareas = append(l.tareas, t)

}

// Marcar tarea como completada
func (l *ListaTareas) marcarCompletado(index int) {

	l.tareas[index].completado = true

}

//Actualizar Tarea

func (l *ListaTareas) actualizarTarea(index int, t Tarea) {

	l.tareas[index] = t
}

//eliminar Tarea

func (l *ListaTareas) eliminarTarea(index int) {

	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {

	//instanciar lista de tareas

	lista := ListaTareas{}

	//Entrada de datos

	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int

		fmt.Println("Seleccione una opcion:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Println("ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {

		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Println("Ingrese la descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')

			lista.agregarTarea(t)
			fmt.Println("La tarea se ha agregado")

		case 2:

			var index int
			fmt.Println("Ingrese el indice de la tarea que quiere marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completa")

		case 3:

			var index int
			var t Tarea

			fmt.Println("Ingrese el indice de la tarea que quiere actualizar: ")
			fmt.Scanln(&index)

			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Println("Ingrese la descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')

			lista.actualizarTarea(index, t)
			fmt.Println("Tarea actualizada")

		case 4:

			var index int
			fmt.Println("Ingrese el indice de la tarea que quiere eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)

		case 5:
			fmt.Println("Saliendo del programa")
			return

		default:
			fmt.Println("Opcion invalida")
		}

		fmt.Println("Lista de tareas ")
		fmt.Println("===========================")

		for i, t := range lista.tareas {

			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)

		}
		fmt.Println("===========================")

	}

	//listar tareas

}
