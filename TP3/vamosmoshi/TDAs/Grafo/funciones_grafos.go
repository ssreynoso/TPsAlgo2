package grafo

import (
	TDACola "vamosmoshi/TDAs/Cola"
	TDAHash "vamosmoshi/TDAs/Hash"
	TDAHeap "vamosmoshi/TDAs/Heap"
	TDAPila "vamosmoshi/TDAs/Pila"
)

func EsConexo(g Grafo) bool {
	return CntComponentesConexas(g) == 1
}

func CntComponentesConexas(g Grafo) int {
	visitados := TDAHash.CrearHash[string, string]()
	cnt_conexos := 0

	for _, v := range g.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			cnt_conexos++
			visitados.Guardar(v, v)
			dfs_cnt_cmp_conexas(g, v, visitados)
		}
	}

	return cnt_conexos
}

func dfs_cnt_cmp_conexas(g Grafo, v string, visitados TDAHash.Diccionario[string, string]) {
	for _, w := range g.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(w) {
			visitados.Guardar(w, w)
			dfs_cnt_cmp_conexas(g, w, visitados)
		}
	}
}

func RecorridoBfs(g Grafo, verticeInicial string) (TDAHash.Diccionario[string, string], TDAHash.Diccionario[string, int]) {
	if !g.ExisteVertice(verticeInicial) {
		panic("El vertice no pertenece al grafo")
	}

	visitados := TDAHash.CrearHash[string, string]()
	padres := TDAHash.CrearHash[string, string]()
	ordenes := TDAHash.CrearHash[string, int]()

	padres.Guardar(verticeInicial, "")
	ordenes.Guardar(verticeInicial, 0)
	visitados.Guardar(verticeInicial, verticeInicial)

	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar(verticeInicial)

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, w := range g.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				padres.Guardar(w, v)
				ordenes.Guardar(w, ordenes.Obtener(v)+1)
				visitados.Guardar(w, w)
				cola.Encolar(w)
			}
		}
	}
	return padres, ordenes
}

func RecorridoDfs(g Grafo, verticeInicial string) (TDAHash.Diccionario[string, string], TDAHash.Diccionario[string, int]) {
	if !g.ExisteVertice(verticeInicial) {
		panic("El vertice no pertenece al grafo")
	}

	visitados := TDAHash.CrearHash[string, string]()
	padres := TDAHash.CrearHash[string, string]()
	ordenes := TDAHash.CrearHash[string, int]()

	padres.Guardar(verticeInicial, "")
	ordenes.Guardar(verticeInicial, 0)
	visitados.Guardar(verticeInicial, verticeInicial)

	dfs(g, verticeInicial, visitados, padres, ordenes)

	return padres, ordenes
}

func RecorridoDfsCompleto(g Grafo) (TDAHash.Diccionario[string, string], TDAHash.Diccionario[string, int]) {
	visitados := TDAHash.CrearHash[string, string]()
	padres := TDAHash.CrearHash[string, string]()
	ordenes := TDAHash.CrearHash[string, int]()

	for _, v := range g.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, v)
			padres.Guardar(v, "")
			ordenes.Guardar(v, 0)
			dfs(g, v, visitados, padres, ordenes)
		}
	}

	return padres, ordenes
}

func dfs(
	g Grafo,
	v string,
	visitados TDAHash.Diccionario[string, string],
	padres TDAHash.Diccionario[string, string],
	ordenes TDAHash.Diccionario[string, int],
) {
	for _, w := range g.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(w) {
			padres.Guardar(w, v)
			ordenes.Guardar(w, ordenes.Obtener(v)+1)
			visitados.Guardar(w, w)
			dfs(g, w, visitados, padres, ordenes)
		}
	}
}

func ReconstruirCamino(g Grafo, padres TDAHash.Diccionario[string, string], destino string) []string {
	recorrido := []string{}
	pila := TDAPila.CrearPilaDinamica[string]()
	_destino := destino

	for _destino != "" {
		pila.Apilar(_destino)
		_destino = padres.Obtener(_destino)
	}

	for !pila.EstaVacia() {
		recorrido = append(recorrido, pila.Desapilar())
	}

	return recorrido
}

func OrdenTopologico(g Grafo) []string {
	grados := TDAHash.CrearHash[string, int]()

	// Pongo en cero los grados de todos los vértices.
	for _, v := range g.ObtenerVertices() {
		grados.Guardar(v, 0)
	}

	// Por cada arista que apunte a un vértice le sumo un grado.
	for _, v := range g.ObtenerVertices() {
		for _, w := range g.ObtenerAdyacentes(v) {
			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}

	// Creo la cola que me va a definir el orden topológico.
	cola := TDACola.CrearColaEnlazada[string]()

	// Encolo solo los que tienen grado 0
	for _, v := range g.ObtenerVertices() {
		if grados.Obtener(v) == 0 {
			cola.Encolar(v)
		}
	}

	// Creo el slice que voy a devolver al final de la función
	orden := []string{}

	// Itero hasta que la cola esté vacía
	for !cola.EstaVacia() {

		// Desencolo al primer elemento y lo guardo en el slice
		v := cola.Desencolar()
		orden = append(orden, v)

		// Me fijo sus adyacentes y les resto un grado porque ya encolé a su padre
		for _, w := range g.ObtenerAdyacentes(v) {
			grados.Guardar(w, grados.Obtener(w)-1)

			// Si alguno tiene grado 0 (no tiene otro padre) lo encolo
			if grados.Obtener(w) == 0 {
				cola.Encolar(w)
			}
		}
	}

	return orden
}

type item struct {
	dato string
	peso int
}

func cmp_items(i1, i2 *item) int {
	if i1.peso < i2.peso {
		return 1
	}
	if i1.peso > i2.peso {
		return -1
	}
	return 0
}

// Dijkstra
func CaminoMinimoDijkstra(g Grafo, origen string) (TDAHash.Diccionario[string, string], TDAHash.Diccionario[string, int]) {
	padres := TDAHash.CrearHash[string, string]()
	distancias := TDAHash.CrearHash[string, int]()

	// Pongo todas las distancias con un número muy alto
	for _, v := range g.ObtenerVertices() {
		distancias.Guardar(v, 999999)
	}

	// El origen tiene distancia 0 y padre nulo ("" en este caso)
	padres.Guardar(origen, "")
	distancias.Guardar(origen, 0)

	el := new(item)
	el.dato = origen
	el.peso = 0

	// El heap contiene structs del tipo item que tienen un dato y un peso. Compara los pesos
	cola := TDAHeap.CrearHeap(cmp_items)
	cola.Encolar(el)

	// Puedo tener muchas veces el mismo elemento pero en los diccionarios siempre voy a tener las distancias mínimas.
	for !cola.EstaVacia() {
		v := cola.Desencolar().dato
		// Recorro todos los adyacentes
		for _, w := range g.ObtenerAdyacentes(v) {
			// Si la distancia es menor a la ya almacenada la reemplazo y cambio al padre, si no no hago nada.
			if distancias.Obtener(v)+g.ObtenerPesoArista(v, w) < distancias.Obtener(w) {
				distancias.Guardar(w, distancias.Obtener(v)+g.ObtenerPesoArista(v, w))
				padres.Guardar(w, v)
				it := new(item)
				it.dato = w
				it.peso = distancias.Obtener(w)
				cola.Encolar(it)
			}
		}
	}

	return padres, distancias
}

// Bellman-Ford
func CaminoMinimoBF(g Grafo, origen string) (TDAHash.Diccionario[string, string], TDAHash.Diccionario[string, int]) {
	padres := TDAHash.CrearHash[string, string]()
	distancias := TDAHash.CrearHash[string, int]()

	// Pongo todas las distancias con un número muy alto
	for _, v := range g.ObtenerVertices() {
		distancias.Guardar(v, 999999)
	}

	// El origen tiene distancia 0 y padre nulo ("" en este caso)
	padres.Guardar(origen, "")
	distancias.Guardar(origen, 0)

	el := new(item)
	el.dato = origen
	el.peso = 0

	// Puedo tener muchas veces el mismo elemento pero en los diccionarios siempre voy a tener las distancias mínimas.
	for i := 0; i < len(g.ObtenerVertices()); i++ {
		cambio := false
		for _, w := range g.ObtenerAristas() {
			// Si la distancia es menor a la ya almacenada la reemplazo y cambio al padre, si no no hago nada.
			if distancias.Obtener(w[0])+g.ObtenerPesoArista(w[0], w[1]) < distancias.Obtener(w[1]) {
				cambio = true
				padres.Guardar(w[1], w[0])
				distancias.Guardar(w[1], distancias.Obtener(w[0])+g.ObtenerPesoArista(w[0], w[1]))
			}
		}
		if !cambio {
			break
		}
	}

	for _, w := range g.ObtenerAristas() {
		if distancias.Obtener(w[0])+g.ObtenerPesoArista(w[0], w[1]) < distancias.Obtener(w[1]) {
			panic("Hay un ciclo negativo")
		}
	}

	return padres, distancias
}

// func CompConexCompletas(g Grafo) bool {
// 	comp_conexas := GetCompConex(g) // O(v + e)
// 	flg := true

// 	comp_conexas.Iterar(func(cmp string, cantidad_cmp int) bool { // O(v)
// 		adyacentes := g.ObtenerAdyacentes(cmp) //O(e_v)

// 		for _, v := range adyacentes {
// 			if len(g.ObtenerAdyacentes(v)) != cantidad_cmp-1 || len(adyacentes) != cantidad_cmp-1 {
// 				flg = false
// 				return false
// 			}
// 		}

// 		return true
// 	})

// 	return flg
// }

// func GetCompConex(g Grafo) TDAHash.Diccionario[string, int] {
// 	visitados := TDAHash.CrearHash[string, string]()
// 	cmp_conexas := TDAHash.CrearHash[string, int]()

// 	for _, v := range g.ObtenerVertices() { // O v + e
// 		if !visitados.Pertenece(v) {
// 			cmp_conexas.Guardar(v, 1)
// 			visitados.Guardar(v, v)
// 			dfs2(g, v, visitados, v, cmp_conexas)
// 		}
// 	}

// 	return cmp_conexas
// }

// func dfs2(g Grafo, v string, visitados TDAHash.Diccionario[string, string], id_cmp string, cmp_conexas TDAHash.Diccionario[string, int]) {
// 	for _, w := range g.ObtenerAdyacentes(v) {
// 		if !visitados.Pertenece(w) {
// 			visitados.Guardar(w, w)
// 			cmp_conexas.Guardar(id_cmp, cmp_conexas.Obtener(id_cmp)+1)
// 			dfs2(g, w, visitados, id_cmp, cmp_conexas)
// 		}
// 	}
// }

// func CompConexCompletas(g Grafo) bool {
// 	visitados := TDAHash.CrearHash[string, string]()
// 	padres := TDAHash.CrearHash[string, string]()

// 	for _, v := range g.ObtenerVertices() {
// 		if !visitados.Pertenece(v) {
// 			visitados.Guardar(v, v)
// 			padres.Guardar(v, "")
// 			dfs2(g, v, visitados, padres)
// 		}
// 	}
// }

// func dfs2(g Grafo, v string, visitados TDAHash.Diccionario[string, string], padres TDAHash.Diccionario[string, string]) {
// 	for _, w := range g.ObtenerAdyacentes(v) {
// 		if !visitados.Pertenece(w) {
// 			visitados.Guardar(w, w)
// 			padres.Guardar(w, v)
// 		}
// 	}
// }
