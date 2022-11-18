package cola_prioridad_test

import (
	TDAColaPrioridad "cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmpInts(a, b int) int {
	if a > b {
		return 1
	}

	if a < b {
		return -1
	}

	return 0
}

func TestColasPrioridadVacia(t *testing.T) {
	cola_prioridad := TDAColaPrioridad.CrearHeap(cmpInts)

	arreglo := []int{}
	cola_prioridad2 := TDAColaPrioridad.CrearHeapArr(arreglo, cmpInts)

	require.True(t, cola_prioridad.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad.VerMax() })
	require.True(t, cola_prioridad2.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad2.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad2.VerMax() })
}

func TestEncolarYDesencolar(t *testing.T) {
	//creo la cola y encolo elementos
	cola_prioridad := TDAColaPrioridad.CrearHeap(cmpInts)

	arreglo := []int{32, 8, 16, 2, 4}
	cola_prioridad2 := TDAColaPrioridad.CrearHeapArr(arreglo, cmpInts)

	cola_prioridad.Encolar(32)
	cola_prioridad.Encolar(8)
	cola_prioridad.Encolar(16)
	cola_prioridad.Encolar(2)
	cola_prioridad.Encolar(4)

	//me fijo que hayan quedado en el orden deseado
	require.Equal(t, 32, cola_prioridad.Desencolar())
	require.Equal(t, 16, cola_prioridad.Desencolar())
	require.Equal(t, 8, cola_prioridad.Desencolar())
	require.Equal(t, 4, cola_prioridad.Desencolar())
	require.Equal(t, 2, cola_prioridad.Desencolar())

	require.Equal(t, 32, cola_prioridad2.Desencolar())
	require.Equal(t, 16, cola_prioridad2.Desencolar())
	require.Equal(t, 8, cola_prioridad2.Desencolar())
	require.Equal(t, 4, cola_prioridad2.Desencolar())
	require.Equal(t, 2, cola_prioridad2.Desencolar())
}

func TestVerPrimero(t *testing.T) {
	//verifica que al encolar y desencolar siempre va a quedar como primer elemento
	cola_prioridad := TDAColaPrioridad.CrearHeap(cmpInts)
	arreglo := []int{0, 4, 3, 6, 2, 8, 9, 7, 5, 1}
	cola_prioridad2 := TDAColaPrioridad.CrearHeapArr(arreglo, cmpInts)

	for i := 0; i < 10; i++ {
		cola_prioridad.Encolar(i)
		require.Equal(t, i, cola_prioridad.VerMax())
		cola_prioridad.Desencolar()
	}

	for i := 0; i < 10; i++ {
		cola_prioridad.Encolar(i)
	}

	for i := 9; i >= 0; i-- {
		require.Equal(t, i, cola_prioridad.VerMax())
		cola_prioridad.Desencolar()
	}

	for i := 9; i >= 0; i-- {
		require.Equal(t, i, cola_prioridad2.VerMax())
		cola_prioridad2.Desencolar()
	}
}

func TestColaPrioridadVolumen(t *testing.T) {
	cola_prioridad := TDAColaPrioridad.CrearHeap(cmpInts)
	arreglo := []int{}
	cola_prioridad2 := TDAColaPrioridad.CrearHeapArr(arreglo, cmpInts)

	for i := 10000; i > 0; i-- {
		cola_prioridad.Encolar(i)
	}
	for j := 10000; j > 0; j-- {
		require.Equal(t, j, cola_prioridad.Cantidad())
		require.Equal(t, j, cola_prioridad.Desencolar())
	}

	for i := 1; i <= 10000; i++ {
		cola_prioridad.Encolar(i)
		require.Equal(t, i, cola_prioridad.Cantidad())
	}
	for j := 10000; j > 0; j-- {
		require.Equal(t, j, cola_prioridad.Cantidad())
		require.Equal(t, j, cola_prioridad.Desencolar())
	}

	for i := 10000; i > 0; i-- {
		cola_prioridad2.Encolar(i)
	}

	for j := 10000; j > 0; j-- {
		require.Equal(t, j, cola_prioridad2.Cantidad())
		require.Equal(t, j, cola_prioridad2.Desencolar())
	}

	for i := 1; i <= 10000; i++ {
		cola_prioridad2.Encolar(i)
		require.Equal(t, i, cola_prioridad2.Cantidad())
	}
	for j := 10000; j > 0; j-- {
		require.Equal(t, j, cola_prioridad2.Cantidad())
		require.Equal(t, j, cola_prioridad2.Desencolar())
	}

	require.True(t, cola_prioridad.EstaVacia())
	require.True(t, cola_prioridad2.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad2.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola_prioridad2.VerMax() })
}

// Testeo

func cmpInts2(a, b int) int {
	if a < b {
		return 1
	}

	if a > b {
		return -1
	}

	return 0
}

func TestHeapSort(t *testing.T) {
	arreglo := []int{9, 13, 1, 5, 25, 347, 133, 20, 45}
	arregloOrdenado := []int{1, 5, 9, 13, 20, 25, 45, 133, 347}

	TDAColaPrioridad.HeapSort(arreglo, cmpInts)

	for i, value := range arregloOrdenado {
		require.Equal(t, value, arreglo[i])
	}

	arreglo2 := []int{50, 1, 2, 3, 30, 10, 80}
	arregloOrdenado2 := []int{80, 50, 30, 10, 3, 2, 1}

	TDAColaPrioridad.HeapSort(arreglo2, cmpInts2)

	for i, value := range arregloOrdenado2 {
		require.Equal(t, value, arreglo2[i])
	}
}
