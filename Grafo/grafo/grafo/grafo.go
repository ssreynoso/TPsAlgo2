package grafo

import (
	TDAHash "grafo/TDAs/Hash"
)

type Grafo interface {
	ExisteVertice(string) bool
	AgregarVertice(string)
	EliminarVertice(string)
	AgregarArista(string, string, int)
	EliminarArista(string, string)
	SonAdyacentes(string, string) bool
	ObtenerPesoArista(string, string) int
	ObtenerAdyacentes(string) []string
	ObtenerVertices() []string
	ObtenerAristas() [][]string
}

type grafo struct {
	esDirigido bool
	vertices   TDAHash.Diccionario[string, TDAHash.Diccionario[string, int]]
}

func CrearGrafo(esDirigido bool) Grafo {
	grafo := new(grafo)
	grafo.vertices = TDAHash.CrearHash[string, TDAHash.Diccionario[string, int]]()
	grafo.esDirigido = esDirigido
	return grafo
}

func (g grafo) ExisteVertice(vertice string) bool {
	return g.vertices.Pertenece(vertice)
}

func (g grafo) AgregarVertice(vertice string) {
	if vertice == "" {
		panic("No se puede ingresar un vertice nulo")
	}

	if !g.ExisteVertice(vertice) {
		g.vertices.Guardar(vertice, TDAHash.CrearHash[string, int]())
	}
}

func (g grafo) EliminarVertice(vertice string) {
	if !g.ExisteVertice(vertice) {
		panic("El vertice no pertenece al grafo")
	}

	g.vertices.Iterar(func(v string, w_s TDAHash.Diccionario[string, int]) bool {
		if w_s.Pertenece(vertice) {
			w_s.Borrar(vertice)
		}
		return true
	})

	g.vertices.Borrar(vertice)
}

func (g grafo) AgregarArista(verticeInicial, verticeFinal string, peso int) {
	if !g.ExisteVertice(verticeInicial) || !g.ExisteVertice(verticeFinal) {
		panic("Al menos un vertice no pertenece al grafo")
	}

	g.vertices.Obtener(verticeInicial).Guardar(verticeFinal, peso)

	if !g.esDirigido {
		g.vertices.Obtener(verticeFinal).Guardar(verticeInicial, peso)
	}
}

func (g grafo) EliminarArista(verticeInicial, verticeFinal string) {
	if !g.ExisteVertice(verticeInicial) {
		panic("El vertice no pertenece al grafo")
	}

	if !g.SonAdyacentes(verticeInicial, verticeFinal) {
		panic("La arista no pertenece al vertice")
	}

	g.vertices.Obtener(verticeInicial).Borrar(verticeFinal)
}

func (g grafo) SonAdyacentes(verticeInicial, verticeFinal string) bool {
	if !g.ExisteVertice(verticeInicial) {
		panic("El vertice no pertenece al grafo")
	}
	return g.vertices.Obtener(verticeInicial).Pertenece(verticeFinal)
}

func (g grafo) ObtenerPesoArista(verticeInicial, verticeFinal string) int {
	if !g.SonAdyacentes(verticeInicial, verticeFinal) {
		panic("Los vertices no son adyacentes")
	}
	return g.vertices.Obtener(verticeInicial).Obtener(verticeFinal)
}

func (g grafo) ObtenerAdyacentes(vertice string) []string {
	if !g.ExisteVertice(vertice) {
		panic("El vertice no pertenece al grafo")
	}

	adyacentes := []string{}

	g.vertices.Obtener(vertice).Iterar(func(v string, peso int) bool {
		adyacentes = append(adyacentes, v)
		return true
	})

	return adyacentes
}

func (g grafo) ObtenerVertices() []string {
	vertices := []string{}

	g.vertices.Iterar(func(v string, adyacentes TDAHash.Diccionario[string, int]) bool {
		vertices = append(vertices, v)
		return true
	})

	return vertices
}

func (g grafo) ObtenerAristas() [][]string {
	aristas := [][]string{}

	g.vertices.Iterar(func(v string, dato TDAHash.Diccionario[string, int]) bool {
		dato.Iterar(func(w string, peso int) bool {
			aristas = append(aristas, []string{v, w})
			return true
		})
		return true
	})

	return aristas
}
