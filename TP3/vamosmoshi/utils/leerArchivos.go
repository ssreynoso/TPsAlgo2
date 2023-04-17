package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAGrafo "vamosmoshi/TDAs/Grafo"
	TDAHash "vamosmoshi/TDAs/Hash"
	"vamosmoshi/customTDAs"
)

func LeerArchivoCiudades(ruta string, grafoCiudades TDAGrafo.Grafo, hashCiudades TDAHash.Diccionario[string, customTDAs.Ciudad]) {
	var (
		lineaLeida  string
		nroLineas   int
		cont        int
		flgCiudades bool = true
	)

	archivo, err := os.Open(ruta)

	if err != nil {
		return
	}

	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	for s.Scan() {
		lineaLeida = s.Text()
		if nroLineas == 0 {
			nroLineas, _ = strconv.Atoi(lineaLeida)
		} else {
			cont++
		}
		if flgCiudades {
			if cont != 0 {
				datos := strings.Split(lineaLeida, ",")
				ciudad := customTDAs.CrearCiudad(datos[0], datos[1], datos[2])
				hashCiudades.Guardar(datos[0], ciudad)
				grafoCiudades.AgregarVertice(datos[0])
			}
		} else {
			if cont != 0 {
				datos2 := strings.Split(lineaLeida, ",")
				peso, _ := strconv.Atoi(datos2[2])
				grafoCiudades.AgregarArista(datos2[0], datos2[1], peso)
			}
		}

		if cont == nroLineas {
			nroLineas = 0
			cont = 0
			flgCiudades = false
		}
	}

	err = s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	}
}
