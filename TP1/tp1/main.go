package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/TDAs"
	"rerepolez/customTDAs"
	errores "rerepolez/errores"
	"rerepolez/operaciones"
	votos "rerepolez/votos"
	"strings"
)

func main() {

	params := os.Args[1:] // No me importa el nombre del programa

	if len(params) < 2 {
		errorParamentos := new(errores.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", errorParamentos.Error())
		return
	}

	rutaListaPartido := params[0]  //Ruta del archivo csv donde están los partidos y candidatos
	rutaListaPadrones := params[1] // Ruta del archivo txt con la lista de padrones válidos

	listaPartidos, errorLeerArchivo := customTDAs.CrearListaPartidos(rutaListaPartido)
	if errorLeerArchivo != nil {
		return
	}

	listaPadrones, errorLeerArchivo := customTDAs.CrearListaPadrones(rutaListaPadrones)
	if errorLeerArchivo != nil {
		return
	}

	colaVotantes := TDAs.CrearColaEnlazada[votos.Votante]()

	var contadorInpugnados int

	// Lista de personas que ya votaron
	listaDNIsYaVotaron := customTDAs.CrearListaDNIsYaVotaron()

	// Leo input del usuario
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		menu(s.Text(), listaPartidos, listaPadrones, colaVotantes, &contadorInpugnados, listaDNIsYaVotaron)
	}

	operaciones.FinPrograma(colaVotantes, listaPartidos, &contadorInpugnados)

	err := s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	}
}

func menu(
	input string,
	listaPartidos customTDAs.ListaPartidos,
	listaPadrones customTDAs.ListaPadrones,
	colaVotantes TDAs.Cola[votos.Votante],
	contadorInpugnados *int,
	listaDNIsYaVotaron customTDAs.ListaDNIs,
) {
	params := strings.Split(input, " ") // Comando a evaluar
	cmd := params[0]
	data := params[1:]

	switch cmd {
	case "ingresar":
		operaciones.Ingresar(data, listaPadrones, colaVotantes)
	case "votar":
		operaciones.Votar(data, listaPartidos, colaVotantes, listaDNIsYaVotaron)
	case "deshacer":
		operaciones.Deshacer(colaVotantes, listaDNIsYaVotaron)
	case "fin-votar":
		operaciones.FinVotar(colaVotantes, listaPartidos, contadorInpugnados, listaDNIsYaVotaron)
	default:
		fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("ERROR: [%s] el parámetro ingresado no es válido.", cmd))
	}
}
