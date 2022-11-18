package diccionario

import TDAPila "diccionario/pila"

type arbolBusquedaBinaria[K comparable, V any] struct {
	raiz        *nodoAbb[K, V]
	funcion_cmp func(K, K) int
	cantidad    int
}

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) *arbolBusquedaBinaria[K, V] {
	nuevoAbb := new(arbolBusquedaBinaria[K, V])
	nuevoAbb.raiz = nil
	nuevoAbb.funcion_cmp = funcion_cmp
	return nuevoAbb
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nuevoNodo := new(nodoAbb[K, V])
	nuevoNodo.clave = clave
	nuevoNodo.dato = dato
	return nuevoNodo
}

func (ab *arbolBusquedaBinaria[K, V]) Guardar(clave K, dato V) {
	guardar(&ab.raiz, &ab.cantidad, clave, dato, ab.funcion_cmp)
}

func guardar[K comparable, V any](ptrNodo **nodoAbb[K, V], cantidadArbol *int, clave K, dato V, comparar func(K, K) int) {
	if *ptrNodo == nil {
		*ptrNodo = crearNodoAbb(clave, dato)
		*cantidadArbol++
		return
	}

	comparacion := comparar(clave, (*ptrNodo).clave)
	switch {
	case comparacion < 0:
		guardar(&(*ptrNodo).izq, cantidadArbol, clave, dato, comparar)
	case comparacion > 0:
		guardar(&(*ptrNodo).der, cantidadArbol, clave, dato, comparar)
	default:
		(*ptrNodo).dato = dato // Actualizo el dato
	}
}

func (ab *arbolBusquedaBinaria[K, V]) Pertenece(clave K) bool {
	return pertenece(ab.raiz, clave, ab.funcion_cmp)
}

func pertenece[K comparable, V any](nodo *nodoAbb[K, V], clave K, comparar func(K, K) int) bool {
	if nodo == nil {
		return false
	}

	comparacion := comparar(clave, nodo.clave)
	switch {
	case comparacion < 0:
		return pertenece(nodo.izq, clave, comparar)
	case comparacion > 0:
		return pertenece(nodo.der, clave, comparar)
	default:
		return true
	}
}

func (ab *arbolBusquedaBinaria[K, V]) Obtener(clave K) V {
	if !ab.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	return obtener(ab.raiz, clave, ab.funcion_cmp)
}

func obtener[K comparable, V any](nodo *nodoAbb[K, V], clave K, comparar func(K, K) int) V {
	comparacion := comparar(clave, nodo.clave)
	switch {
	case comparacion < 0:
		return obtener(nodo.izq, clave, comparar)
	case comparacion > 0:
		return obtener(nodo.der, clave, comparar)
	default:
		return nodo.dato
	}
}

func (ab *arbolBusquedaBinaria[K, V]) Borrar(clave K) V {
	if !ab.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	ab.cantidad--
	return borrar(ab.raiz, clave, &ab.raiz, &ab.raiz, ab.funcion_cmp)
}

func borrar[K comparable, V any](
	nodo *nodoAbb[K, V],
	clave K,
	ptrHijoIzqDePadre **nodoAbb[K, V],
	ptrHijoDerDePadre **nodoAbb[K, V],
	comparar func(K, K) int,
) V {
	comparacion := comparar(clave, nodo.clave)
	switch {
	case comparacion < 0:
		return borrar(nodo.izq, clave, &nodo.izq, nil, comparar)
	case comparacion > 0:
		return borrar(nodo.der, clave, nil, &nodo.der, comparar)
	default:
		var conector **nodoAbb[K, V]

		if ptrHijoIzqDePadre != nil {
			conector = ptrHijoIzqDePadre
		} else {
			conector = ptrHijoDerDePadre
		}

		return borrarNodo(nodo, conector)
	}
}

func borrarNodo[K comparable, V any](nodo *nodoAbb[K, V], ptrHijoDePadre **nodoAbb[K, V]) V {
	cntHijos := 0

	if nodo.izq != nil {
		cntHijos++
	}

	if nodo.der != nil {
		cntHijos++
	}

	switch cntHijos {
	case 0:
		*ptrHijoDePadre = nil
	case 1:
		if nodo.izq != nil {
			*ptrHijoDePadre = nodo.izq
		}

		if nodo.der != nil {
			*ptrHijoDePadre = nodo.der
		}
	case 2:
		nodoReemplazante, _ := buscarMayor(nodo.izq, &nodo.izq)
		nodoReemplazante.der = nodo.der
		nodoReemplazante.izq = nodo.izq
		*ptrHijoDePadre = nodoReemplazante
	}

	return nodo.dato
}

func buscarMayor[K comparable, V any](nodo *nodoAbb[K, V], conector **nodoAbb[K, V]) (*nodoAbb[K, V], **nodoAbb[K, V]) {
	if nodo.der != nil {
		return buscarMayor(nodo.der, &nodo.der)
	}

	borrarNodo(nodo, conector)
	return nodo, conector
}

func (ab *arbolBusquedaBinaria[K, V]) Cantidad() int {
	return ab.cantidad
}

// ------------------- Iterador interno -------------------

func (ab *arbolBusquedaBinaria[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterar(ab.raiz, visitar)
}

func iterar[K comparable, V any](nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if !iterar(nodo.izq, visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	if !iterar(nodo.der, visitar) {
		return false
	}

	return true
}

// ------------------- Iterador externo -------------------

type iteradorDiccionario[K comparable, V any] struct {
	pilaNodos TDAPila.Pila[*nodoAbb[K, V]]
}

func crearIteradorABB[K comparable, V any](abb *arbolBusquedaBinaria[K, V]) IterDiccionario[K, V] {
	iter := new(iteradorDiccionario[K, V])
	iter.pilaNodos = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()

	apilarHijosIzq(abb.raiz, iter.pilaNodos)
	return iter
}

func apilarHijosIzq[K comparable, V any](nodo *nodoAbb[K, V], pilaNodos TDAPila.Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}

	pilaNodos.Apilar(nodo)
	apilarHijosIzq(nodo.izq, pilaNodos)
}

func (abb *arbolBusquedaBinaria[K, V]) Iterador() IterDiccionario[K, V] {
	iter := crearIteradorABB(abb)
	return iter
}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	tope := iter.pilaNodos.VerTope()
	return tope.clave, tope.dato
}

func (iter *iteradorDiccionario[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	tope := iter.pilaNodos.Desapilar()
	apilarHijosIzq(tope.der, iter.pilaNodos)
	return tope.clave
}

func (iter *iteradorDiccionario[K, V]) HaySiguiente() bool {
	return !iter.pilaNodos.EstaVacia()
}

// ------------------- Iteración por rangos -------------------

// Iterador interno

func (ab arbolBusquedaBinaria[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if desde == nil && hasta == nil {
		ab.Iterar(visitar)
		return
	}

	iterarRango(ab.raiz, desde, hasta, ab.funcion_cmp, visitar)
}

func iterarRango[K comparable, V any](nodo *nodoAbb[K, V], desde *K, hasta *K, comparar func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	switch {
	case (desde != nil && comparar(nodo.clave, *desde) < 0):
		iterarRango(nodo.der, desde, hasta, comparar, visitar)
	case (hasta != nil && comparar(nodo.clave, *hasta) > 0):
		iterarRango(nodo.izq, desde, hasta, comparar, visitar)
	default:
		if !iterarRango(nodo.izq, desde, hasta, comparar, visitar) {
			return false
		}
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
		if !iterarRango(nodo.der, desde, hasta, comparar, visitar) {
			return false
		}
	}

	return true
}

// Iterador externo

type iteradorDiccionarioRangos[K comparable, V any] struct {
	pilaNodos   TDAPila.Pila[*nodoAbb[K, V]]
	desde       *K
	hasta       *K
	funcion_cmp func(K, K) int
}

func (abb *arbolBusquedaBinaria[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := crearIteradorABBRangos(abb, desde, hasta)
	return iter
}

func crearIteradorABBRangos[K comparable, V any](abb *arbolBusquedaBinaria[K, V], desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iteradorDiccionarioRangos[K, V])
	iter.pilaNodos = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	iter.funcion_cmp = abb.funcion_cmp

	apilarHijosIzqRangos(abb.raiz, iter.pilaNodos, desde, hasta, abb.funcion_cmp)
	return iter
}

func apilarHijosIzqRangos[K comparable, V any](
	nodo *nodoAbb[K, V],
	pilaNodos TDAPila.Pila[*nodoAbb[K, V]],
	desde *K,
	hasta *K,
	comparar func(K, K) int,
) {
	if nodo == nil {
		return
	}

	switch {
	case (desde != nil && comparar(nodo.clave, *desde) < 0):
		apilarHijosIzqRangos(nodo.der, pilaNodos, desde, hasta, comparar)
	case (hasta != nil && comparar(nodo.clave, *hasta) > 0):
		apilarHijosIzqRangos(nodo.izq, pilaNodos, desde, hasta, comparar)
	default:
		pilaNodos.Apilar(nodo)
		apilarHijosIzqRangos(nodo.izq, pilaNodos, desde, hasta, comparar)
	}
}

func (iter iteradorDiccionarioRangos[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	tope := iter.pilaNodos.VerTope()
	return tope.clave, tope.dato
}

func (iter *iteradorDiccionarioRangos[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	tope := iter.pilaNodos.Desapilar()
	apilarHijosIzqRangos(tope.der, iter.pilaNodos, iter.desde, iter.hasta, iter.funcion_cmp)
	return tope.clave
}

func (iter *iteradorDiccionarioRangos[K, V]) HaySiguiente() bool {
	return !iter.pilaNodos.EstaVacia()
}
