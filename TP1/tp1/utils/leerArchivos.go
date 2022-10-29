package utils

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/errores"
	"rerepolez/sort"
	"rerepolez/votos"
	"strconv"
	"strings"
)

func leerArchivo(ruta string, cb func(string) bool) {
	var (
		flgContinuo bool = true
		lineaLeida  string
	)

	archivo, err := os.Open(ruta)

	if err != nil {
		return
	}

	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	for s.Scan() && flgContinuo {
		lineaLeida = s.Text()
		flgContinuo = cb(lineaLeida)
	}

	err = s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	}
}

func validarExistenciaArchivo(ruta string) bool {
	archivo, err := os.Open(ruta)

	if err != nil {
		return false
	}

	defer archivo.Close()

	return true
}

func ProcesarArchivoBoletas(ruta string) ([]votos.Partido, error) {
	if !validarExistenciaArchivo(ruta) {
		err := new(errores.ErrorLeerArchivo)
		return nil, err
	}

	listaPartidos := []votos.Partido{}

	agregarPartido := func(partidoString string) bool {
		partidoArray := strings.Split(partidoString, ",")

		nombrePartido := partidoArray[0]
		candidatos := [votos.CANT_VOTACION]string{partidoArray[1], partidoArray[2], partidoArray[3]}

		partido := votos.CrearPartido(nombrePartido, candidatos)

		listaPartidos = append(listaPartidos, partido)
		return true
	}

	leerArchivo(ruta, agregarPartido)

	return listaPartidos, nil
}

func ProcesarArchivoPadrones(ruta string) ([]int, error) {

	if !validarExistenciaArchivo(ruta) {
		err := new(errores.ErrorLeerArchivo)
		return nil, err
	}

	var (
		listaPadronesString []string
		listaPadronesInt    []int
		mayorLongitud       int
	)

	agregarPadron := func(padron string) bool {
		if len(padron) > mayorLongitud {
			mayorLongitud = len(padron)
		}
		listaPadronesString = append(listaPadronesString, padron)
		return true
	}

	leerArchivo(ruta, agregarPadron)

	listaPadronesString = sort.RadixSort(listaPadronesString, mayorLongitud)

	// Convertimos el array a array de enteros
	for _, el := range listaPadronesString {
		elementoConvertido, err := strconv.Atoi(el)

		if err != nil {
			break
		}

		listaPadronesInt = append(listaPadronesInt, elementoConvertido)
	}

	return listaPadronesInt, nil
}
