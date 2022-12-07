package operaciones

import (
	"algogram/customTDAs"
	"algogram/errores"
	"fmt"
	"os"
)

func Logout(
	usuario_sesion *customTDAs.Usuario,
) {
	if *usuario_sesion == nil {
		error := new(errores.ErrorNoHayLoggeado)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	*usuario_sesion = nil

	fmt.Fprintf(os.Stdout, "Adios\n")
}
