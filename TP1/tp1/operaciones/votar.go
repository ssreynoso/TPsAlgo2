package operaciones

import (
	"fmt"
	"os"
	"rerepolez/TDAs"
	"rerepolez/customTDAs"
	errores "rerepolez/errores"
	"rerepolez/votos"
	"strconv"
)

func Votar(
	data []string,
	listaPartidos customTDAs.ListaPartidos,
	colaVotantes TDAs.Cola[votos.Votante],
	listaDNIsYaVotaron customTDAs.ListaDNIs,
) {
	if len(data) < 2 {
		err := new(errores.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	if colaVotantes.EstaVacia() {
		err := new(errores.FilaVacia)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	var tipoVoto votos.TipoVoto
	tipoVotoString := data[0]

	switch tipoVotoString {
	case "Presidente":
		tipoVoto = votos.PRESIDENTE
	case "Gobernador":
		tipoVoto = votos.GOBERNADOR
	case "Intendente":
		tipoVoto = votos.INTENDENTE
	default:
		err := new(errores.ErrorTipoVoto)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	numeroListaString := data[1]
	numeroLista, err := strconv.Atoi(numeroListaString)

	// Valido que exista el número de lista
	if err != nil || !listaPartidos.ValidarNumeroLista(numeroLista) {
		err := new(errores.ErrorAlternativaInvalida)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	if listaDNIsYaVotaron.PadronFraudulento(colaVotantes.VerPrimero().LeerDNI()) {
		err := new(errores.ErrorVotanteFraudulento)
		err.Dni = colaVotantes.Desencolar().LeerDNI()
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	colaVotantes.VerPrimero().Votar(tipoVoto, numeroLista)

	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}
