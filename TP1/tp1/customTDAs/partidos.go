package customTDAs

import "main/votos"

type ListaPartidos interface {

	// Sumar voto recine un tipo de voto y un voto.
	// En base a estos datos, suma el voto en la estructura para quien corresponda.
	SumarVoto(votos.Voto) error

	// ...
	ValidarNumeroLista(int) bool

	ImprimirResultados(votos.TipoVoto)
}
