package cola_test

import (
	TDACola "cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.Equal(t, true, cola.EstaVacia(), "La cola está vacía, devuelve true")
	cola.Encolar(10)
	require.Equal(t, false, cola.EstaVacia(), "La cola no está vacía, devuelve false")
	cola.Desencolar()
	require.Equal(t, true, cola.EstaVacia(), "La cola vuelve a estar vacía, devuelve true")
	cola.Encolar(5)
	require.Equal(t, false, cola.EstaVacia(), "La cola no está vacía, devuelve false")
	cola.Encolar(5)
	require.Equal(t, false, cola.EstaVacia(), "La cola no está vacía, devuelve false")
	cola.Desencolar()
	require.Equal(t, false, cola.EstaVacia(), "La cola no está vacía, devuelve false")
	cola.Desencolar()
	require.Equal(t, true, cola.EstaVacia(), "La cola vuelve a estar vacía, devuelve true")
}

func TestUnElemento(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	datoEntero := 2

	colaStrings := TDACola.CrearColaEnlazada[string]()
	datoString := "Hola"

	colaColas := TDACola.CrearColaEnlazada[TDACola.Cola[int]]()
	datoCola := TDACola.CrearColaEnlazada[int]()

	testUnElemento(t, colaEnteros, datoEntero)
	testUnElemento(t, colaStrings, datoString)
	testUnElemento(t, colaColas, datoCola)
}

func testUnElemento[Type any](t *testing.T, cola TDACola.Cola[Type], dato Type) {
	cola.Encolar(dato)
	require.Equal(t, dato, cola.VerPrimero(), "El primero es el elemento que recien encolé")
	require.Equal(t, dato, cola.Desencolar(), "El elemento que desencolo es el que recien encolé")
}

func TestEncolarDesencolar(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	datosEnteros := []int{2, 3, 5, 7, 11, 13, 15}

	colaStrings := TDACola.CrearColaEnlazada[string]()
	datosStrings := []string{"Hola", "Como", "Estas", "Todo", "Bien", "Jeje", "Chau"}

	testEncolarDesencolar(t, colaEnteros, datosEnteros)
	testEncolarDesencolar(t, colaStrings, datosStrings)
}

func TestEncolarDesencolarMasivo(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()

	const inicio int = 0
	const fin int = 10000

	for i := inicio; i < fin; i++ {
		colaEnteros.Encolar(i)
		require.Equal(t, inicio, colaEnteros.VerPrimero(), "El primero sigue siendo el primero que encolé")
	}

	for i := inicio; i < fin; i++ {
		require.Equal(t, i, colaEnteros.Desencolar(), "Respeto la invariante FIFO de la cola")
	}
}

func testEncolarDesencolar[Type any](t *testing.T, cola TDACola.Cola[Type], datos []Type) {
	cola.Encolar(datos[0])
	require.Equal(t, datos[0], cola.VerPrimero(), "El primero es el elemento que recien encolé")
	require.Equal(t, datos[0], cola.Desencolar(), "El elemento que desencolo es el que recien encolé")
	cola.Encolar(datos[0])
	cola.Encolar(datos[1])
	cola.Encolar(datos[2])
	cola.Encolar(datos[3])
	require.Equal(t, datos[0], cola.VerPrimero(), "El primero es el elemento que encolé primero")
	require.Equal(t, datos[0], cola.Desencolar(), "El elemento que desencolo es el que encolé primero")
	require.Equal(t, datos[1], cola.VerPrimero(), "El primero es el elemento que encolé segundo")
	require.Equal(t, datos[1], cola.Desencolar(), "El elemento que desencolo es el que encolé segundo")
	require.Equal(t, datos[2], cola.VerPrimero(), "El primero es el elemento que encolé tercero")
	require.Equal(t, datos[2], cola.Desencolar(), "El elemento que desencolo es el que encolé tercero")
	require.Equal(t, datos[3], cola.VerPrimero(), "El primero es el elemento que encolé último")
	require.Equal(t, datos[3], cola.Desencolar(), "El elemento que desencolo es el que encolé último")
	cola.Encolar(datos[4])
	require.Equal(t, datos[4], cola.VerPrimero(), "El primero es el elemento que recien encolé")
	require.Equal(t, datos[4], cola.Desencolar(), "El elemento que desencolo es el que recien encolé")
	cola.Encolar(datos[5])
	require.Equal(t, datos[5], cola.VerPrimero(), "El primero es el elemento que recien encolé")
	require.Equal(t, datos[5], cola.Desencolar(), "El elemento que desencolo es el que recien encolé")
	cola.Encolar(datos[6])
	require.Equal(t, datos[6], cola.VerPrimero(), "El primero es el elemento que recien encolé")
	require.Equal(t, datos[6], cola.Desencolar(), "El elemento que desencolo es el que recien encolé")
}
