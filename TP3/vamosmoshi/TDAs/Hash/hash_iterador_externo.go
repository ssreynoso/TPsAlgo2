package diccionarioHASH

import lista "vamosmoshi/TDAs/Lista"

// ------------------- Iterador Externo -------------------

type iteradorDiccionario[K comparable, V any] struct {
	dic             *DiccionarioHashAbierto[K, V]
	actual          *KeyValuePair[K, V]
	iterListaActual lista.IteradorLista[*KeyValuePair[K, V]]
	ultIndiceLista  int
}

func crearIteradorHash[K comparable, V any](dic *DiccionarioHashAbierto[K, V]) IterDiccionario[K, V] {
	iter := new(iteradorDiccionario[K, V])
	var iter2 lista.IteradorLista[*KeyValuePair[K, V]]

	iter.dic = dic

	for i := 0; i < dic.capacidad; i++ {
		if dic.datos[i] != nil {
			iter2 = dic.datos[i].Iterador()
			if iter2.HaySiguiente() {
				iter.actual = iter2.Siguiente()
				iter.iterListaActual = iter2
				iter.ultIndiceLista = i
				break
			}
		}
	}

	return iter
}

func (dic *DiccionarioHashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := crearIteradorHash(dic)
	return iter
}

func (iter iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.key, iter.actual.value
}

func (iter *iteradorDiccionario[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	var iter2 lista.IteradorLista[*KeyValuePair[K, V]]
	anterior := iter.actual
	i := iter.ultIndiceLista + 1

	// Sigo con el iterador de la lista hasta que termina.
	if iter.iterListaActual.HaySiguiente() {
		iter.actual = iter.iterListaActual.Siguiente()
	} else {
		// Busco la siguiente lista y su primer elemento
		for i < iter.dic.capacidad {
			if iter.dic.datos[i] != nil {
				iter2 = iter.dic.datos[i].Iterador()
				if iter2.HaySiguiente() {
					iter.actual = iter2.Siguiente()
					iter.iterListaActual = iter2
					iter.ultIndiceLista = i
					break
				}
			}
			i++
		}
	}

	if i == iter.dic.capacidad && anterior == iter.actual {
		iter.actual = nil
	}

	return anterior.key
}
