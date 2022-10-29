package customTDAs

import "rerepolez/TDAs"

type listaDNIsYaVotaron struct {
	dnis TDAs.Lista[int]
}

func CrearListaDNIsYaVotaron() ListaDNIs {
	nuevaLista := new(listaDNIsYaVotaron)
	nuevaLista.dnis = TDAs.CrearListaEnlazada[int]()
	return nuevaLista
}

func (lista *listaDNIsYaVotaron) InsertarDNI(dni int) {
	lista.dnis.InsertarUltimo(dni)
}

func (lista *listaDNIsYaVotaron) PadronFraudulento(dni int) bool {
	iter := lista.dnis.Iterador()
	flgPadronFraudulento := false

	for iter.HaySiguiente() && !flgPadronFraudulento {
		if iter.VerActual() == dni {
			flgPadronFraudulento = true
		}
		iter.Siguiente()
	}

	return flgPadronFraudulento
}
