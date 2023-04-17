package customTDAs

import (
	Heap "algogram/TDAs/Heap"
	"math"
)

type Usuario interface {

	// AgregarPostAlFeed recibe un puntero a un post y lo encola en el feed.
	AgregarPostAlFeed(*Post)

	// GetID devuelve el id de un usuario.
	GetID() string

	// Devuelve true si quedan posts para ver; en caso contrario, false.
	HayPostsParaVer() bool

	// VerSiguientePost devuelve el puntero al post con mayor prioridad en el feed.
	VerSiguientePost() *Post

	// Devuelve la posición del usuario en el archivo de usuarios proveído al inicio del programa.
	GetIndex() int
}

type usuario struct {
	id    string
	index int
	feed  Heap.ColaPrioridad[*Post]
}

func CrearUsuario(id string, index int) Usuario {
	u := new(usuario)
	u.id = id
	u.index = index

	compararPrioridades := func(post1, post2 *Post) int {
		distancia_1 := math.Abs(float64(index - (*post1).GetUsuario().GetIndex()))
		distancia_2 := math.Abs(float64(index - (*post2).GetUsuario().GetIndex()))

		switch {
		case distancia_1 < distancia_2:
			return 1
		case distancia_1 > distancia_2:
			return -1
		default:
			if (*post1).GetFecha().Before((*post2).GetFecha()) {
				return 1
			} else {
				return -1
			}
		}
	}

	u.feed = Heap.CrearHeap(compararPrioridades)
	return u
}

func (u usuario) AgregarPostAlFeed(post *Post) {
	u.feed.Encolar(post)
}

func (u usuario) GetID() string {
	return u.id
}

func (u usuario) HayPostsParaVer() bool {
	return !u.feed.EstaVacia()
}

func (u usuario) VerSiguientePost() *Post {
	return u.feed.Desencolar()
}

func (u usuario) GetIndex() int {
	return u.index
}
