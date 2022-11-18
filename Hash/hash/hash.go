package diccionario

import (
	"diccionario/lista"
)

type DiccionarioHashAbierto[K comparable, V any] struct {
	datos               []lista.Lista[*KeyValuePair[K, V]]
	capacidad           int
	cantidad            int
	hashFunction        func(K) uint64
	MAXIMO_FACTOR_CARGA int
}

func CrearHash[K comparable, V any]() *DiccionarioHashAbierto[K, V] {
	const MAXIMO_FACTOR_CARGA = 2
	const CAPACIDAD_INICIAL = 83

	dic := new(DiccionarioHashAbierto[K, V])
	dic.capacidad = CAPACIDAD_INICIAL
	dic.datos = make([]lista.Lista[*KeyValuePair[K, V]], CAPACIDAD_INICIAL)
	dic.hashFunction = InicializarHashFunction[K](CAPACIDAD_INICIAL)
	dic.MAXIMO_FACTOR_CARGA = MAXIMO_FACTOR_CARGA

	return dic
}

func (dic *DiccionarioHashAbierto[K, V]) Guardar(clave K, dato V) {
	index := dic.hashFunction(clave)
	newPair := CrearKeyValuePair(clave, dato)

	if dic.Pertenece(clave) {
		dic.datos[index].Iterar(func(kvp *KeyValuePair[K, V]) bool {
			if kvp.key == clave {
				kvp.value = dato
				return false
			}
			return true
		})

		return
	}

	if dic.datos[index] == nil {
		dic.datos[index] = lista.CrearListaEnlazada[*KeyValuePair[K, V]]()
	}

	dic.datos[index].InsertarUltimo(&newPair)

	dic.cantidad++

	factorCarga := float32(dic.cantidad) / float32(dic.capacidad)
	if factorCarga > float32(dic.MAXIMO_FACTOR_CARGA) {
		nuevaCapacidad := dic.capacidad * 4
		dic.redimensionar(nuevaCapacidad)
	}
}

func (dic *DiccionarioHashAbierto[K, V]) Pertenece(clave K) bool {
	index := dic.hashFunction(clave)

	if dic.datos[index] == nil {
		return false
	}

	pertenece := false
	pertenecePtr := &pertenece
	dic.datos[index].Iterar(func(kvp *KeyValuePair[K, V]) bool {
		if clave == kvp.key {
			*pertenecePtr = true
			return false
		}
		return true
	})

	return pertenece
}

func (dic *DiccionarioHashAbierto[K, V]) Obtener(clave K) V {
	if !dic.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	index := dic.hashFunction(clave)

	var dato V
	datoPtr := &dato
	dic.datos[index].Iterar(func(kvp *KeyValuePair[K, V]) bool {
		if clave == kvp.key {
			*datoPtr = kvp.value
			return false
		}
		return true
	})

	return dato
}

func (dic *DiccionarioHashAbierto[K, V]) Borrar(clave K) V {
	if !dic.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	index := dic.hashFunction(clave)
	iter := dic.datos[index].Iterador()
	var pair *KeyValuePair[K, V]
	var dato V

	for iter.HaySiguiente() {
		pair = iter.VerActual()
		if clave == pair.key {
			dato = iter.Borrar().value
			break
		} else {
			iter.Siguiente()
		}
	}

	dic.cantidad--
	factorCarga := float32(dic.cantidad) / float32(dic.capacidad)
	if factorCarga <= float32(dic.MAXIMO_FACTOR_CARGA/4) {
		nuevaCapacidad := dic.capacidad / 2
		dic.redimensionar(nuevaCapacidad)
	}

	return dato
}

func (dic DiccionarioHashAbierto[K, V]) Cantidad() int {
	return dic.cantidad
}

// ------------------- Iterador Interno -------------------

func (dic *DiccionarioHashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	var (
		iter lista.IteradorLista[*KeyValuePair[K, V]]
		pair *KeyValuePair[K, V]
	)

	for i := 0; i < dic.capacidad; i++ {
		if dic.datos[i] != nil {
			iter = dic.datos[i].Iterador()
			for iter.HaySiguiente() {
				pair = iter.Siguiente()
				if !visitar(pair.key, pair.value) {
					return
				}
			}
		}
	}
}

// ------------------- Redimensión -------------------

func (dic *DiccionarioHashAbierto[K, V]) redimensionar(nuevaCapacidad int) {
	// Creo un nuevo array de listas.
	datos_viejos := dic.datos
	capacidad_viejo := dic.capacidad

	nuevos_datos := make([]lista.Lista[*KeyValuePair[K, V]], nuevaCapacidad)

	dic.datos = nuevos_datos
	dic.capacidad = nuevaCapacidad
	dic.hashFunction = InicializarHashFunction[K](nuevaCapacidad)
	dic.cantidad = 0

	var (
		iter lista.IteradorLista[*KeyValuePair[K, V]]
		pair *KeyValuePair[K, V]
	)

	// Tengo que reubicar todos los pares kvp
	for i := 0; i < capacidad_viejo; i++ {
		if datos_viejos[i] != nil {
			iter = datos_viejos[i].Iterador()
			for iter.HaySiguiente() {
				pair = iter.Siguiente()
				dic.Guardar(pair.key, pair.value)
			}
		}
	}
}
