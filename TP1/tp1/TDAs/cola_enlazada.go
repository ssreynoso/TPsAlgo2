package TDAs

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nodo := nodoCrear(dato)
	if cola.EstaVacia() {
		cola.primero = nodo
	} else {
		cola.ultimo.siguiente = nodo
	}
	cola.ultimo = nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	nodo := cola.primero
	cola.primero = nodo.siguiente

	if cola.ultimo == nodo {
		cola.ultimo = cola.primero
	}

	return nodo.dato
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T]) // nodo := &nodoCola{} -> nodo va a ser de tipo *nodoCola[T]
	nodo.dato = dato
	return nodo
}
