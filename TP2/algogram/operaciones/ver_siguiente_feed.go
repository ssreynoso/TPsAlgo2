package operaciones

import (
	"algogram/customTDAs"
	"algogram/errores"
	"fmt"
	"os"
)

func VerSiguienteFeed(
	usuario_sesion *customTDAs.Usuario,
) {
	if *usuario_sesion == nil || !(*usuario_sesion).HayPostsParaVer() {
		error := new(errores.ErrorVerSigFeed)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	ptrPostAVer := (*usuario_sesion).VerSiguientePost()

	(*ptrPostAVer).Informacion()
}
