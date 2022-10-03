package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	n := len(vector)

	if n <= 0 {
		return -1
	}

	posicion_max := 0

	for i := 0; i < n; i++ {
		if vector[i] > vector[posicion_max] {
			posicion_max = i
		}
	}
	return posicion_max
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	n1, n2 := len(vector1), len(vector2)

	for i := 0; i < n1 && i < n2; i++ {
		switch {
		case vector1[i] > vector2[i]:
			return 1
		case vector2[i] > vector1[i]:
			return -1
		}
	}

	switch {
	case n1 > n2:
		return 1
	case n2 > n1:
		return -1
	}

	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	n := len(vector)
	var posicion_max int

	for i := (n - 1); i >= 0; i-- {
		posicion_max = Maximo(vector[:i+1])
		Swap(&vector[i], &vector[posicion_max])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	n := len(vector)

	if n == 0 {
		return 0
	}
	suma := vector[n-1] + Suma(vector[:(n-1)])
	return suma
}

// EsPalindromo devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsPalindromo(cadena string) bool {
	largoCadena := len(cadena)

	if largoCadena < 2 {
		return true
	}

	inicio := cadena[0]
	final := cadena[largoCadena-1]

	if inicio == final && EsPalindromo(cadena[1:largoCadena-1]) {
		return true
	}

	return false
}
