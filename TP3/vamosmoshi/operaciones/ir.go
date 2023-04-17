package operaciones

import (
	"fmt"
	"os"
	TDAGrafo "vamosmoshi/TDAs/Grafo"
	TDAHash "vamosmoshi/TDAs/Hash"
	customTDAs "vamosmoshi/customTDAs"
	"vamosmoshi/errores"
)

func Ir(
	input []string,
	grafoCiudades TDAGrafo.Grafo,
	hashCiudades TDAHash.Diccionario[string, customTDAs.Ciudad],
) {
	desde := input[0][:len(input[0])-1]
	hasta := input[1][:len(input[1])-1]
	// archivo := input[2][:len(input[2])-1]

	if !hashCiudades.Pertenece(desde) || !hashCiudades.Pertenece(hasta) || !grafoCiudades.SonAdyacentes(desde, hasta) {
		error := new(errores.ErrorNoSeEncontroRecorrido)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	// Hacer camino minimo desde hasta

	// Guardar en archivo kml el resultado del camino minimo

	fmt.Fprintf(os.Stdout, "Post likeado\n")
}
