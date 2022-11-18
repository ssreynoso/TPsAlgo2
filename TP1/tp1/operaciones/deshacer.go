package operaciones

import (
	"fmt"
	"os"
	"rerepolez/cola"
	"rerepolez/customTDAs"
	errores "rerepolez/errores"
	"rerepolez/votos"
)

func Deshacer(
	colaVotantes cola.Cola[votos.Votante],
	listaDNIsYaVotaron customTDAs.ListaDNIs,
) {
	if colaVotantes.EstaVacia() {
		err := new(errores.FilaVacia)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	if listaDNIsYaVotaron.PadronFraudulento(colaVotantes.VerPrimero().LeerDNI()) {
		err := new(errores.ErrorVotanteFraudulento)
		err.Dni = colaVotantes.Desencolar().LeerDNI()
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
