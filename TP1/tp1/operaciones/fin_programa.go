package operaciones

import (
	"fmt"
	"main/TDAs"
	"main/customTDAs"
	errores "main/errores"
	"main/votos"
	"os"
)

func FinPrograma(
	colaVotantes TDAs.Cola[votos.Votante],
	listaPartidos customTDAs.ListaPartidos,
	contadorInpugnados *int,
) {
	if !colaVotantes.EstaVacia() {
		err := new(errores.ErrorCiudadanosSinVotar)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", "")
	fmt.Fprintf(os.Stdout, "%s\n", "Presidente:")
	listaPartidos.ImprimirResultados(votos.PRESIDENTE)

	fmt.Fprintf(os.Stdout, "%s\n", "")
	fmt.Fprintf(os.Stdout, "%s\n", "Gobernador:")
	listaPartidos.ImprimirResultados(votos.GOBERNADOR)

	fmt.Fprintf(os.Stdout, "%s\n", "")
	fmt.Fprintf(os.Stdout, "%s\n", "Intendente:")
	listaPartidos.ImprimirResultados(votos.INTENDENTE)

	var txtVoto string

	if *contadorInpugnados == 0 || *contadorInpugnados > 1 {
		txtVoto = "votos"
	} else {
		txtVoto = "voto"
	}

	fmt.Fprintf(os.Stdout, "%s\n", "")
	fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("Votos Impugnados: %d %s", *contadorInpugnados, txtVoto))
}
