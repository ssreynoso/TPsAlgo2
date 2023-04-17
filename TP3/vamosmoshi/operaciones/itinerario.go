package operaciones

import (
	"fmt"
	"os"
	TDAGrafo "vamosmoshi/TDAs/Grafo"
	TDAHash "vamosmoshi/TDAs/Hash"
	customTDAs "vamosmoshi/customTDAs"
	"vamosmoshi/errores"
)

func Itinerario(
	input []string,
	grafoCiudades TDAGrafo.Grafo,
	hashCiudades TDAHash.Diccionario[string, customTDAs.Ciudad],
) {
	if *usuario_sesion != nil {
		error := new(errores.ErrorYaHayLoggeado)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	usuario := strings.Join(input, " ")

	if !usuariosRegistrados.Pertenece(usuario) {
		error := new(errores.ErrorUsuarioInexistente)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	*usuario_sesion = usuariosRegistrados.GetUsuario(usuario)

	fmt.Fprintf(os.Stdout, "Hola %s\n", usuario)
}
