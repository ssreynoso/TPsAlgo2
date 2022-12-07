package operaciones

import (
	"algogram/customTDAs"
	"algogram/errores"
	"fmt"
	"os"
	"strings"
)

func Login(
	input []string,
	usuariosRegistrados customTDAs.Usuarios,
	usuario_sesion *customTDAs.Usuario,
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

	*usuario_sesion = usuariosRegistrados.Login(usuario)

	fmt.Fprintf(os.Stdout, "Hola %s\n", usuario)
}
