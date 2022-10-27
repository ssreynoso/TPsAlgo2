package main

import (
	// "bufio"
	"fmt"
	"main/errores"
	"main/votos"

	// "strconv"

	// errores "main/errores"
	"os"
	// "strings"
)

func main() {

	var (
		listaPartidos    []votos.Partido
		listaPadrones    []int
		errorLeerArchivo *errores.ErrorLeerArchivo
	)

	params := os.Args[1:] // No me importa el nombre del programa
	fmt.Println(params)

	if len(params) < 2 {
		errorParamentos := new(errores.ErrorParametros)
		fmt.Println(errorParamentos.Error())
		return
	}

	rutaListaPartido := params[0]  //Ruta del archivo csv donde están los partidos y candidatos
	rutaListaPadrones := params[1] // Ruta del archivo txt con la lista de padrones válidos

	// Leo archivo de candidatos y lleno structs
	listaPartidos, errorLeerArchivo = ProcesarArchivoBoletas(rutaListaPartido)
	if errorLeerArchivo != nil {
		fmt.Println(errorLeerArchivo.Error())
		return
	}

	// Tengo que leer el archivo txt y cargarlo en memoria para poder realizar búsquedas binarias después
	listaPadrones, errorLeerArchivo = ProcesarArchivoPadrones(rutaListaPadrones)
	if errorLeerArchivo != nil {
		fmt.Println(errorLeerArchivo.Error())
		return
	}

	fmt.Println(listaPadrones)
	fmt.Println(listaPartidos)

	// Creo una cola.

	// input()
}

// func input() {
// 	// fmt.Println("Ingrese de a uno, ingrese una linea vacia para terminar")
// 	// Creamos un scanner que vea el input en la consola.
// 	s := bufio.NewScanner(os.Stdin)
// 	// Mientras no se ingrese una línea vacía, se leerá todo input.
// 	for s.Scan() {
// 		if s.Text() == "" {
// 			break
// 		}
// 		menu(s.Text())
// 	}
// 	// Verificamos que no haya un error con respecto a la lectura.
// 	err := s.Err()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func menu(input string) {

// 	params := strings.Split(input, " ") // Comando a evaluar
// 	cmd := params[0]
// 	data := params[1:]

// 	fmt.Println(cmd)
// 	fmt.Println(data)

// 	switch cmd {
// 	case "ingresar":
// 		//
// 	case "votar":
// 		// Apilo la acción
// 	case "deshacer":
// 		// Desapilo la última acción que hizo el votante
// 	case "fin-votar":
// 		// Proceso la pila del primer votante
// 	default:
// 		fmt.Println("\n", "ERROR: ", cmd, "El parámetro ingresado no es válido")
// 	}
// }

// func EmitirVoto(tipoDato string, lista string) {

// 	dniInt, err := strconv.Atoi(dniString)

// 	switch tipoDato {
// 	case "Presidente":
// 		fmt.Println("OK")
// 	case "Gobernador":
// 		fmt.Println("OK")
// 	case "Intendente":
// 		fmt.Println("OK")
// 	default:
// 		fmt.Println("ERROR: Tipo de voto inválido")

// 	}

// }
