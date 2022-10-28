package votos

import "fmt"

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
	partido.votos[tipo] += 1
}

// Retorna la cantidad de votos que tiene el tipo de voto solicitado.
func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {

	var txtVoto string

	nombrePartido := partido.nombre
	candidatoPartido := partido.candidatos[tipo]
	cntVotos := partido.votos[tipo]

	if cntVotos == 0 || cntVotos > 1 {
		txtVoto = "votos"
	} else {
		txtVoto = "voto"
	}

	txt := fmt.Sprintf("%s - %s: %d %s", nombrePartido, candidatoPartido, cntVotos, txtVoto)

	return txt
}

// Recibe un entero, el tipo de voto al que le tengo que sumar uno.
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo] += 1
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	var txtVoto string

	cntVotos := blanco.votos[tipo]

	if cntVotos == 0 || cntVotos > 1 {
		txtVoto = "votos"
	} else {
		txtVoto = "voto"
	}

	txt := fmt.Sprintf("Votos en Blanco: %d %s", cntVotos, txtVoto)

	return txt
}
