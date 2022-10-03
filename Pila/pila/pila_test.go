package pila_test

import (
	TDAPila "pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.Equal(t, true, pila.EstaVacia(), "La pila está vacía, devuelve true")
	pila.Apilar(5)
	require.Equal(t, false, pila.EstaVacia(), "La pila no está vacía, devuelve false")
	pila.Desapilar()
	require.Equal(t, true, pila.EstaVacia(), "La pila vuelve a estar vacía, devuelve true")
	pila.Apilar(5)
	require.Equal(t, false, pila.EstaVacia(), "La pila no está vacía, devuelve false")
	pila.Apilar(5)
	require.Equal(t, false, pila.EstaVacia(), "La pila no está vacía, devuelve false")
	pila.Desapilar()
	require.Equal(t, false, pila.EstaVacia(), "La pila no está vacía, devuelve false")
	pila.Desapilar()
	require.Equal(t, true, pila.EstaVacia(), "La pila vuelve a estar vacía, devuelve true")
}

func TestUnElemento(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	datoEntero := 2

	pilaStrings := TDAPila.CrearPilaDinamica[string]()
	datoString := "Hola"

	pilaPilas := TDAPila.CrearPilaDinamica[TDAPila.Pila[int]]()
	datoPila := TDAPila.CrearPilaDinamica[int]()

	testUnElemento(t, pilaEnteros, datoEntero)
	testUnElemento(t, pilaStrings, datoString)
	testUnElemento(t, pilaPilas, datoPila)
}

func TestApilarDesapilar(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	datosEnteros := []int{2, 3, 5, 7, 11, 13, 15}

	pilaStrings := TDAPila.CrearPilaDinamica[string]()
	datosStrings := []string{"Hola", "Como", "Estas", "Todo", "Bien", "Jeje", "Chau"}

	testApilarDesapilar(t, pilaEnteros, datosEnteros)
	testApilarDesapilar(t, pilaStrings, datosStrings)

	for i := 0; i < 1000; i++ {
		pilaEnteros.Apilar(i)
	}

	for i := 999; i >= 0; i-- {
		require.Equal(t, i, pilaEnteros.Desapilar(), "Respeto el orden de desapilado")
	}

}

func testUnElemento[Type any](t *testing.T, pila TDAPila.Pila[Type], dato Type) {
	pila.Apilar(dato)
	require.Equal(t, dato, pila.VerTope(), "El tope es el elemento que recien apilé")
	require.Equal(t, dato, pila.Desapilar(), "El elemento que desapilo es el que recien apilé")
}

func testApilarDesapilar[Type any](t *testing.T, pila TDAPila.Pila[Type], datos []Type) {

	apilarMucho := func(dato Type) {
		for i := 0; i < 1000; i++ {
			pila.Apilar(dato)
			require.Equal(t, dato, pila.VerTope(), "El tope es el elemento que recien apilé")
		}
	}

	desapilarMucho := func(dato Type) {
		for i := 0; i < 1000; i++ {
			require.Equal(t, dato, pila.Desapilar(), "El elemento que desapilo igual al esperado")
		}
	}

	// tipoDato := fmt.Sprintf("%T", datos[0])

	pila.Apilar(datos[0])
	require.Equal(t, datos[0], pila.VerTope(), "El tope es el elemento que recien apilé")
	require.Equal(t, datos[0], pila.Desapilar(), "El elemento que desapilo es el que recien apilé")
	pila.Apilar(datos[0])
	pila.Apilar(datos[1])
	require.Equal(t, datos[1], pila.VerTope(), "El tope es el elemento que recien apilé")
	require.Equal(t, datos[1], pila.Desapilar(), "El elemento que desapilo es el que recien apilé")
	require.Equal(t, datos[0], pila.VerTope(), "El tope es el elemento que apilé primero")
	require.Equal(t, datos[0], pila.Desapilar(), "El elemento que desapilo es el que apilé primero")
	pila.Apilar(datos[2])
	require.Equal(t, datos[2], pila.VerTope(), "El tope es el elemento que recien apilé")
	require.Equal(t, datos[2], pila.Desapilar(), "El elemento que desapilo es el que recien apilé")

	apilarMucho(datos[0])
	desapilarMucho(datos[0])
}
