package operaciones

import (
	"fmt"
	"main/TDAs"
	errores "main/errores"
	"main/votos"
	"os"
)

func Deshacer(
	colaVotantes TDAs.Cola[votos.Votante],
) {
	if colaVotantes.EstaVacia() {
		err := new(errores.FilaVacia)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	err := colaVotantes.VerPrimero().Deshacer()

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}
