package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAGrafo "vamosmoshi/TDAs/Grafo"
	TDAHash "vamosmoshi/TDAs/Hash"
	customTDAs "vamosmoshi/customTDAs"
	"vamosmoshi/errores"
	"vamosmoshi/operaciones"
	"vamosmoshi/utils"
)

func main() {
	params := os.Args[1:] // No me importa el nombre del programa

	if len(params) < 1 {
		errorParamentos := new(errores.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", errorParamentos.Error())
		return
	}

	// Leo el archivo de ciudades:
	rutaListaCiudades := params[0] // Ruta del archivo donde estarán las ciudades
	grafoCiudades := TDAGrafo.CrearGrafo(false)
	hashCiudades := TDAHash.CrearHash[string, customTDAs.Ciudad]()
	utils.LeerArchivoCiudades(rutaListaCiudades, grafoCiudades, hashCiudades)

	// Leo input del usuario
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		menu(s.Text(), grafoCiudades, hashCiudades)
	}

	err := s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	}
}

func menu(input string, grafoCiudades TDAGrafo.Grafo, hashCiudades TDAHash.Diccionario[string, customTDAs.Ciudad]) {
	params := strings.Split(input, " ") // Comando a evaluar
	cmd := params[0]
	data := params[1:]

	switch cmd {
	case "ir":
		operaciones.Ir(data, grafoCiudades, hashCiudades)
	case "itinerario":
		operaciones.Itinerario(data, grafoCiudades, hashCiudades)
	case "viaje":
		operaciones.Ir(data, grafoCiudades, hashCiudades)
	case "reducir_caminos":
		operaciones.Ir(data, grafoCiudades, hashCiudades)
	default:
		fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("ERROR: [%s] el parámetro ingresado no es válido.", cmd))
	}
}
