package sort

import "rerepolez/TDAs"

// Este Radix Sort utiliza countingsort con colas internamente

func RadixSort(arreglo []string, longitud int) []string {

	const (
		CANTIDAD_COUNTING int  = 10
		POSICION_ASCII    byte = 48
	)
	// Los valores que recibo tiene tamaño 8
	var (
		counting  [CANTIDAD_COUNTING]TDAs.Cola[string]
		character byte
	)

	// Inicializo todas las colas
	for i := 0; i < CANTIDAD_COUNTING; i++ {
		counting[i] = TDAs.CrearColaEnlazada[string]()
	}

	arregloOrdenado := arreglo

	// longitud = 8
	for i := (longitud - 1); i >= 0; i-- {
		for _, el := range arregloOrdenado {
			// Esto entra solo la primera vez porque después todos van a tener len = 8
			if len(el) < longitud {
				ceros := ""
				for j := 0; j < (longitud - len(el)); j++ {
					ceros = ceros + "0"
				}
				el = ceros + el
			}

			// Como la posición de los números [0 - 9] arranca en 48,
			// al restarle este número obtengo al número en cuestión
			character = el[i] - POSICION_ASCII

			counting[character].Encolar(el)
		}

		arregloOrdenado = []string{}

		for j := 0; j < CANTIDAD_COUNTING; j++ {
			var padron string
			for !(counting[j].EstaVacia()) {
				padron = counting[j].Desencolar()
				arregloOrdenado = append(arregloOrdenado, padron)
			}
		}
	}

	return arregloOrdenado
}
