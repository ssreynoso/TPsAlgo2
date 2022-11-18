package votos

import (
	"rerepolez/errores"
	"rerepolez/pila"
)

// Voto tiene guardada la información de un voto emitido, por cada tipo de voto posible.
// Por ejemplo, en la posición GOBERNADOR, tendrá guardada la alternativa a Gobernador.
// Si vale 0, es un voto en blanco.
// Si Impugnado es 'true', entonces no hay que considerar ninguna de las alterantivas señaladas.
type Voto struct {
	VotoPorTipo [CANT_VOTACION]int
	Impugnado   bool
}

type movimiento struct {
	tipoVoto    TipoVoto
	alternativa int
}

type votanteImplementacion struct {
	DNI         int
	movimientos pila.Pila[*movimiento]
}

func CrearVotante(dni int) Votante {
	nuevoVotante := new(votanteImplementacion)
	nuevoVotante.DNI = dni
	nuevoVotante.movimientos = pila.CrearPilaDinamica[*movimiento]()
	return nuevoVotante
}

func crearMovimiento(tipo TipoVoto, alternativa int) *movimiento {
	nuevoMovimiento := new(movimiento)
	nuevoMovimiento.tipoVoto = tipo
	nuevoMovimiento.alternativa = alternativa
	return nuevoMovimiento
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.DNI
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	nuevoMovimiento := crearMovimiento(tipo, alternativa)

	votante.movimientos.Apilar(nuevoMovimiento)

	return nil
}

func (votante *votanteImplementacion) Deshacer() error {

	if votante.movimientos.EstaVacia() {
		err := new(errores.ErrorNoHayVotosAnteriores)
		return err
	}

	votante.movimientos.Desapilar()

	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {

	voto := new(Voto)

	for !votante.movimientos.EstaVacia() {
		movimiento := votante.movimientos.Desapilar()

		if movimiento.alternativa == 0 {
			voto.Impugnado = true
		}

		if voto.VotoPorTipo[movimiento.tipoVoto] == 0 {
			voto.VotoPorTipo[movimiento.tipoVoto] = movimiento.alternativa
		}
	}

	return *voto, nil
}
