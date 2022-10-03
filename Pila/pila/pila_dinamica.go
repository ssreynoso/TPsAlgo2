package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	p := new(pilaDinamica[T])
	p.datos = make([]T, 8)
	// p.cantidad = 0 // Ya se iniciliza con el valor x default de los enteros
	return p
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		nueva_capacidad := (cap(pila.datos) * 2)
		pila.redimensionar((nueva_capacidad))
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	pila.cantidad--
	ultimo_elemento := pila.datos[pila.cantidad]

	if pila.cantidad != 0 && pila.cantidad <= (cap(pila.datos)/4) {
		nueva_capacidad := (cap(pila.datos) / 2)
		pila.redimensionar(nueva_capacidad)
	}

	return ultimo_elemento
}

func (pila *pilaDinamica[T]) redimensionar(nueva_capacidad int) {
	nuevos_datos := make([]T, nueva_capacidad)
	copy(nuevos_datos, pila.datos)
	pila.datos = nuevos_datos
}
