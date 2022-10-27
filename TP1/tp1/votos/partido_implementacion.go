package votos

type partidoImplementacion struct {
	nombre     string
	candidatos [CANT_VOTACION]string
	votos      [CANT_VOTACION]TipoVoto
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]TipoVoto
}

func CrearPartido(nombre string, candidatos [3]string) Partido {
	nuevoPartido := new(partidoImplementacion)
	nuevoPartido.nombre = nombre
	nuevoPartido.candidatos = candidatos
	return nuevoPartido
}

func CrearVotosEnBlanco() Partido {
	nuevoPartido := new(partidoEnBlanco)
	return nuevoPartido
}

// Recibe un entero, el tipo de voto al que le tengo que sumar uno.
func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {

	existe := verificarTipoVoto(tipo)

	if existe {
		partido.votos[tipo] += 1
	}
}

// Retorna la cantidad de votos que tiene el tipo de voto solicitado.
func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return ""
}

// Recibe un entero, el tipo de voto al que le tengo que sumar uno.
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {

	existe := verificarTipoVoto(tipo)

	if existe {
		blanco.votos[tipo] += 1
	}
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}

func verificarTipoVoto(tipoBuscado TipoVoto) bool {
	tiposValidos := []TipoVoto{PRESIDENTE, GOBERNADOR, INTENDENTE}
	for _, tipo := range tiposValidos {
		if tipoBuscado == tipo {
			return true
		}
	}
	return false
}
