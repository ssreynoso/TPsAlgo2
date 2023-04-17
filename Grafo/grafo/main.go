package main

import (
	"fmt"
	TDAGrafo "grafo/grafo"
	"strings"
)

func main() {
	grafo := TDAGrafo.CrearGrafo(false)

	grafo.AgregarVertice("A")
	grafo.AgregarVertice("B")
	grafo.AgregarVertice("C")
	grafo.AgregarVertice("D")
	grafo.AgregarVertice("E")
	grafo.AgregarVertice("F")
	grafo.AgregarVertice("G")
	grafo.AgregarVertice("H")
	grafo.AgregarVertice("I")
	grafo.AgregarVertice("J")

	// componente conexa de 5 elementos
	grafo.AgregarArista("A", "B", 1)
	// grafo.AgregarArista("A", "C", 1)
	// grafo.AgregarArista("A", "D", 1)
	grafo.AgregarArista("A", "E", 1)
	// grafo.AgregarArista("B", "C", 1)
	grafo.AgregarArista("B", "D", 1)
	// grafo.AgregarArista("B", "E", 1)
	grafo.AgregarArista("C", "D", 1)
	// grafo.AgregarArista("C", "E", 1)
	// grafo.AgregarArista("D", "E", 1)

	// componente conexa de 2 elementos
	grafo.AgregarArista("F", "G", 1)

	// componente conexa de 3 elementos
	grafo.AgregarArista("H", "I", 1)
	grafo.AgregarArista("H", "J", 1)
	// grafo.AgregarArista("I", "J", 1)

	// for _, v := range grafo.ObtenerVertices() {
	// 	fmt.Printf("\nVértice: %s \n \t Aristas:\n", v)
	// 	for _, w := range grafo.ObtenerAdyacentes(v) {
	// 		fmt.Printf("\t \t %s, peso: %d\n", w, grafo.ObtenerPesoArista(v, w))
	// 	}
	// }

	// padres, ordenes := TDAGrafo.RecorridoBfs(grafo, "Santi")
	// fmt.Println()
	// padres.Iterar(func(v, p string) bool {
	// 	fmt.Printf("Vertice: %s, padre: %s, orden: %d\n", v, p, ordenes.Obtener(v))
	// 	return true
	// })

	// fmt.Println(TDAGrafo.CntComponentesConexas(grafo))

	// fmt.Printf("Es conexo: %t\n", TDAGrafo.EsConexo(grafo))

	// TIENE QUE SER UN GRAFO DIRIGIDO!
	// fmt.Println(TDAGrafo.OrdenTopologico(grafo))

	// padres, distancias := TDAGrafo.CaminoMinimoDijkstra(grafo, "E")
	padres, distancias := TDAGrafo.RecorridoBfs(grafo, "E")

	padres.Iterar(func(v, p string) bool {
		fmt.Printf("Vertice: %s, padre: %s, distancia: %d\n", v, p, distancias.Obtener(v))
		return true
	})

	recorrido := TDAGrafo.ReconstruirCamino(grafo, padres, "C")

	fmt.Println(strings.Join(recorrido, " -> "))

	// fmt.Println(TDAGrafo.CompConexCompletas(grafo))
}
