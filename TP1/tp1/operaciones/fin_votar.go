package operaciones

import (
	"fmt"
	"os"
	"rerepolez/cola"
	"rerepolez/customTDAs"
	errores "rerepolez/errores"
	"rerepolez/votos"
)

func FinVotar(
	colaVotantes cola.Cola[votos.Votante],
	listaPartidos customTDAs.ListaPartidos,
	contadorInpugnados *int,
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

	voto, err := colaVotantes.VerPrimero().FinVoto()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
	}

	if !voto.Impugnado {
		listaPartidos.SumarVoto(voto)
	} else {
		*contadorInpugnados++
	}

	listaDNIsYaVotaron.InsertarDNI(colaVotantes.Desencolar().LeerDNI())

	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}
