package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Nombre   string
	Email    string
	Telefono string
}

// guardar contactos a un archivo json
// serializa
func saveContact2File(contacts []Contact) error {

	file, err := os.Create("contacts.json")
	if err != nil {

		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts) //codifica el slice de contactos y lo almacena en formato json

	if err != nil {

		return err
	}
	return nil

}

// cargar contactos desde un archivo json
// deserealiza
func loadContactsFromFile(contacts *[]Contact) error {

	file, err := os.Open("contacts.json")
	if err != nil {

		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {

		return err
	}

	return nil
}

func main() {

	var contacts []Contact

	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("=== GESTOR DE CONTACTOS ===\n",
			"1. Agregar un contacto \n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Elige una opción: ")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción")
		}

		switch option {

		case 1:

			var c Contact
			fmt.Print("Nombre: ")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Telefono, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := saveContact2File(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}

		case 2:

			fmt.Println("==================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Nombre, contact.Email, contact.Telefono)
			}
			fmt.Println("==================================")

		case 3:
			return

		}
	}

}
