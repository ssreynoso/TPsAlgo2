package lista_test

import (
	"fmt"
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20}
	testListaVacia(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como estás"}
	testListaVacia(t, listaString, datosString)
}

func testListaVacia[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	require.Equal(t, true, lista.EstaVacia(), "La lista está vacía, devuelve true")
	lista.InsertarPrimero(datos[0])
	require.Equal(t, false, lista.EstaVacia(), "La lista no está vacía, devuelve false")
	lista.BorrarPrimero()
	require.Equal(t, true, lista.EstaVacia(), "La lista vuelve a estar vacía, devuelve true")
	lista.InsertarUltimo(datos[0])
	require.Equal(t, false, lista.EstaVacia(), "La lista no está vacía, devuelve false")
	lista.BorrarPrimero()
	require.Equal(t, true, lista.EstaVacia(), "La lista vuelve a estar vacía, devuelve true")
	lista.InsertarPrimero(datos[0])
	require.Equal(t, false, lista.EstaVacia(), "La lista no está vacía, devuelve false")
	lista.InsertarUltimo(datos[1])
	require.Equal(t, false, lista.EstaVacia(), "La lista no está vacía, devuelve false")
	lista.BorrarPrimero()
	require.Equal(t, false, lista.EstaVacia(), "La lista no está vacía, devuelve false")
	lista.BorrarPrimero()
	require.Equal(t, true, lista.EstaVacia(), "La lista vuelve a estar vacía, devuelve true")
}

func TestListaLargo(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20}
	testListaLargo(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como estás"}
	testListaLargo(t, listaString, datosString)
}

func testListaLargo[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	require.Equal(t, 0, lista.Largo(), "La lista está vacía, el largo es 0")
	lista.InsertarPrimero(datos[0])
	require.Equal(t, 1, lista.Largo(), "La lista no está vacía, el largo es 1")
	lista.BorrarPrimero()
	require.Equal(t, 0, lista.Largo(), "La lista está vacía, el largo es 0")
	lista.InsertarPrimero(datos[0])
	require.Equal(t, 1, lista.Largo(), "La lista no está vacía, el largo es 1")
	lista.InsertarUltimo(datos[1])
	require.Equal(t, 2, lista.Largo(), "La lista no está vacía, el largo es 2")
	lista.BorrarPrimero()
	require.Equal(t, 1, lista.Largo(), "La lista no está vacía, el largo es 1")
	lista.BorrarPrimero()
	require.Equal(t, 0, lista.Largo(), "La lista está vacía, el largo es 0")
}

func TestPrimerYUltimoElemento(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20, 30}
	testPrimerYUltimoElemento(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como", "estás"}
	testPrimerYUltimoElemento(t, listaString, datosString)
}

func testPrimerYUltimoElemento[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	lista.InsertarPrimero(datos[0])
	require.Equal(t, datos[0], lista.VerPrimero(), "El elemento que inserto primero es el primero")
	require.Equal(t, datos[0], lista.VerUltimo(), "Como antes no había elementos, también es el último.")
	lista.InsertarPrimero(datos[1])
	require.Equal(t, datos[1], lista.VerPrimero(), "El elemento que inserto primero es el primero")
	require.Equal(t, datos[0], lista.VerUltimo(), "El elemento que estaba antes queda último")
	lista.InsertarUltimo(datos[2])
	require.Equal(t, datos[1], lista.VerPrimero(), "El elemento que inserté primero sigue siendo el primero")
	require.Equal(t, datos[2], lista.VerUltimo(), "El elemento que inserto último es el último")
}

func TestBorrarPrimero(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20, 30, 40, 50}
	testBorrarPrimero(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como", "estás", "todo", "bien"}
	testBorrarPrimero(t, listaString, datosString)
}

func testBorrarPrimero[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	for i := 0; i < len(datos); i++ {
		lista.InsertarPrimero(datos[i])
		require.Equal(t, datos[i], lista.VerPrimero(), "El elemento que inserto primero es el primero")
		require.Equal(t, datos[i], lista.VerUltimo(), "Como antes no había elementos, también es el último.")
		require.Equal(t, datos[i], lista.BorrarPrimero(), "Borro el primer elemento, es el que inserté antes.")
		require.Equal(t, true, lista.EstaVacia(), "Como solo había un elemento ahora la lista está vacía.")
	}

	lista.InsertarPrimero(datos[0])
	lista.InsertarPrimero(datos[1])
	lista.InsertarPrimero(datos[2])
	lista.InsertarUltimo(datos[3])
	lista.InsertarUltimo(datos[4])
	require.Equal(t, datos[2], lista.BorrarPrimero(), "Respeto el orden de la lista")
	require.Equal(t, datos[1], lista.BorrarPrimero(), "Respeto el orden de la lista")
	require.Equal(t, datos[0], lista.BorrarPrimero(), "Respeto el orden de la lista")
	require.Equal(t, datos[3], lista.BorrarPrimero(), "Respeto el orden de la lista")
	require.Equal(t, datos[4], lista.BorrarPrimero(), "Respeto el orden de la lista")
}

func TestIteradorInterno(t *testing.T) {
	const (
		RESULTADO_SUMA       int = 450
		CONTADOR_SUMA        int = 10
		RESULTADO_SUMA_MITAD int = 350
		CONTADOR_SUMA_MITAD  int = 5
	)

	var (
		resultado         int  = 0
		resultado_puntero *int = &resultado
		contador          int  = 0
		contador_puntero  *int = &contador
	)

	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i += 1 {
		lista.InsertarPrimero(i * 10)
	}

	funcionSumaElementos := func(dato int) bool {
		*resultado_puntero += dato
		*contador_puntero++
		return true
	}

	funcionSumaElementosCondicionada := func(dato int) bool {
		*resultado_puntero += dato
		*contador_puntero++
		return *contador_puntero < 5
	}

	lista.Iterar(funcionSumaElementos)
	require.Equal(t, CONTADOR_SUMA, contador, "Se recorrió la lista completa")
	require.Equal(t, RESULTADO_SUMA, resultado, "Se sumó bien la lista")

	resultado = 0
	contador = 0
	lista.Iterar(funcionSumaElementosCondicionada)
	fmt.Println(contador)
	require.Equal(t, CONTADOR_SUMA_MITAD, contador, "Se recorrió la lista por la mitad, funcionó el caso de corte")
	require.Equal(t, RESULTADO_SUMA_MITAD, resultado, "Se sumó bien la mitad de la lista")
}

func TestRecorrerIteradorExterno(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20, 30, 40, 50}
	testRecorrerIteradorExterno(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como", "estás", "todo", "bien"}
	testRecorrerIteradorExterno(t, listaString, datosString)
}

func testRecorrerIteradorExterno[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	var (
		i          int
		itemActual Type
	)

	for i = 0; i < len(datos); i++ {
		lista.InsertarUltimo(datos[i])
	}
	iteradorExterno := lista.Iterador()
	require.Equal(t, lista.VerPrimero(), iteradorExterno.VerActual(), "El primero al que apunta el iterador es el primero de la lista")
	require.Equal(t, datos[0], iteradorExterno.VerActual(), "El primero al que apunta el iterador es el primero del arreglo de datos")

	// Confirmo que recorro bien todo el TDA
	i = 0
	for iteradorExterno.HaySiguiente() {
		itemActual = iteradorExterno.VerActual()
		require.Equal(t, datos[i], itemActual, "El orden del iterador coincide con el insertado")
		i++
		iteradorExterno.Siguiente()
	}
}

func TestInsertarBorrarPrimeroConIterador(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20, 30, 40, 50}
	testInsertarBorrarPrimeroConIterador(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como", "estás", "todo", "bien"}
	testInsertarBorrarPrimeroConIterador(t, listaString, datosString)
}

func testInsertarBorrarPrimeroConIterador[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	iteradorExterno := lista.Iterador()
	for i := 0; i < len(datos); i++ {
		iteradorExterno.Insertar(datos[i])
		require.Equal(t, datos[i], iteradorExterno.VerActual(), "El primero al que apunta el iterador es el primero del arreglo de datos")
		require.Equal(t, lista.VerPrimero(), iteradorExterno.VerActual(), "El primero al que apunta el iterador es el primero de la lista")
	}

	for i := (len(datos) - 1); i >= 0; i-- {
		require.Equal(t, datos[i], iteradorExterno.Borrar(), "Elimino cada primero con el iterador")
	}
}

func TestInsertarBorrarUltimoConIterador(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	datosInt := []int{10, 20, 30, 40, 50}
	testInsertarBorrarUltimoConIterador(t, listaInt, datosInt)
	listaString := TDALista.CrearListaEnlazada[string]()
	datosString := []string{"Hola", "como", "estás", "todo", "bien"}
	testInsertarBorrarUltimoConIterador(t, listaString, datosString)
}

func testInsertarBorrarUltimoConIterador[Type any](t *testing.T, lista TDALista.Lista[Type], datos []Type) {
	iteradorExterno := lista.Iterador()
	for i := 0; i < len(datos); i++ {
		iteradorExterno.Insertar(datos[i])
		require.Equal(t, iteradorExterno.VerActual(), iteradorExterno.Siguiente(), "El método siguiente devuelve correctamente el elemento actual antes de avanzar")
	}

	// Si inserto a lo último el orden también se respeta
	for i := 0; i < len(datos); i++ {
		require.Equal(t, datos[i], iteradorExterno.Borrar(), "Elimino cada ultimo con el iterador")
	}
}
