package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Martínez": "Edad: \nNombre completo: \nApodo: \nFecha de nacimiento: \nEstatura: \nPeso: ",
		// Agregar un jugador de la selección argentina
		"Garnacho": "Edad: 19\nNombre completo: Alejandro Garnacho Ferreyra\nApodo: Bichito\nFecha de nacimiento:  1 de julio de 2004 \nEstatura: 1.80 m\nPeso: 80 kg",
	}

	// Verificar si se proporciona un argumento (el apellido del jugador a buscar).
	if len(os.Args) != 2 {
		fmt.Println("Uso: tp26.exe apellido_jugador")
		return
	}

	// Obtener el apellido proporcionado como argumento.
	apellido := os.Args[1]

	// Utilizar una declaración switch para buscar y mostrar la descripción.
	switch descripcion := diccionario[apellido]; descripcion {
	case "":
		fmt.Printf("El jugador '%s' no se encuentra.\n", apellido)
	default:
		fmt.Printf("%s:\n%s\n", apellido, descripcion)
	}
}
