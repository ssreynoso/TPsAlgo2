package operaciones

import (
	"fmt"
	"main/TDAs"
	"main/customTDAs"
	errores "main/errores"
	"main/votos"
	"os"
)

func FinVotar(
	colaVotantes TDAs.Cola[votos.Votante],
	listaPartidos customTDAs.ListaPartidos,
	contadorInpugnados *int,
) {
	if colaVotantes.EstaVacia() {
		err := new(errores.FilaVacia)
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

	colaVotantes.Desencolar()

	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}
