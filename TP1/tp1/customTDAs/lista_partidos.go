package customTDAs

import (
	"fmt"
	"main/utils"
	"main/votos"
	"os"
)

type listaPartidos struct {
	partidos []votos.Partido
}

func CrearListaPartidos(ruta string) (ListaPartidos, error) {
	lista, errorLeerArchivo := utils.ProcesarArchivoBoletas(ruta)
	if errorLeerArchivo != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errorLeerArchivo.Error())
		return nil, errorLeerArchivo
	}

	partidoEnBlanco := votos.CrearVotosEnBlanco()

	nuevaListaPartidos := new(listaPartidos)
	nuevaListaPartidos.partidos = append(nuevaListaPartidos.partidos, partidoEnBlanco)
	nuevaListaPartidos.partidos = append(nuevaListaPartidos.partidos, lista...)

	return nuevaListaPartidos, nil
}

func (lista *listaPartidos) SumarVoto(voto votos.Voto) error {

	alternativaPresidente := voto.VotoPorTipo[votos.PRESIDENTE]
	lista.partidos[alternativaPresidente].VotadoPara(votos.PRESIDENTE)

	alternativaGobernador := voto.VotoPorTipo[votos.GOBERNADOR]
	lista.partidos[alternativaGobernador].VotadoPara(votos.GOBERNADOR)

	alternativaIntendente := voto.VotoPorTipo[votos.INTENDENTE]
	lista.partidos[alternativaIntendente].VotadoPara(votos.INTENDENTE)

	return nil
}

func (lista *listaPartidos) ValidarNumeroLista(numeroLista int) bool {
	return (numeroLista >= 0 && numeroLista <= len(lista.partidos))
}

func (lista *listaPartidos) ImprimirResultados(tipo votos.TipoVoto) {
	for _, el := range lista.partidos {
		fmt.Fprintf(os.Stdout, "%s\n", el.ObtenerResultado(tipo))
	}
}
