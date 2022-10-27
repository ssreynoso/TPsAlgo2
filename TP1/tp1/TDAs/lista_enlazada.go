package TDAs

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	largo   int
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if !lista.EstaVacia() {
		return lista.primero.dato
	}
	panic("La lista esta vacia")
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if !lista.EstaVacia() {
		return lista.ultimo.dato
	}
	panic("La lista esta vacia")
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	if lista.EstaVacia() {
		lista.ultimo = nodo
	} else {
		nodo.siguiente = lista.primero
	}
	lista.primero = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.siguiente = nodo
	}
	lista.ultimo = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	nodoABorrar := lista.primero
	lista.primero = nodoABorrar.siguiente
	if lista.Largo() == 1 {
		lista.ultimo = lista.primero
	}
	lista.largo--
	return nodoABorrar.dato
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero
	for i := 0; i < lista.largo; i++ {
		if !visitar(nodoActual.dato) {
			return
		}
		nodoActual = nodoActual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := crearIteradorListaEnlazada(lista)
	return iter
}

// ------------ Iterador externo ------------

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func crearIteradorListaEnlazada[T any](lista *listaEnlazada[T]) IteradorLista[T] {
	iter := new(iterListaEnlazada[T])
	iter.lista = lista
	iter.actual = lista.primero
	return iter
}

func (iter iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
	return iter.anterior.dato
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	if iter.lista.EstaVacia() {
		iter.lista.primero = nodo
		iter.lista.ultimo = nodo
	} else {
		nodo.siguiente = iter.actual
	}
	if iter.anterior == nil {
		iter.lista.primero = nodo
	} else {
		iter.anterior.siguiente = nodo
	}
	if iter.actual == nil {
		iter.lista.ultimo = nodo
	}
	iter.actual = nodo
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iter.actual.dato
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return dato
}
