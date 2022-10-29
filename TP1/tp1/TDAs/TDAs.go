package TDAs

type Pila[T any] interface {

	// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
	EstaVacia() bool

	// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
	// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
	VerTope() T

	// Apilar agrega un nuevo elemento a la pila.
	Apilar(T)

	// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
	// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
	Desapilar() T
}

type Cola[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
	// "La cola esta vacia".
	VerPrimero() T

	// Encolar agrega un nuevo elemento a la cola, al final de la misma.
	Encolar(T)

	// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
	Desencolar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// Recibe un elemento y lo agrega en la primera posición de la lista
	InsertarPrimero(T)

	// Recibe un elemento y lo agrega en la última posición de la lista
	InsertarUltimo(T)

	// Borra al primer elemento de la lista. Si la lista tiene elementos, se saca al primer elemento y
	// se devuelve ese valor. Si está vacía entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// Devuelve el valor del primer elemento. Si la lista tiene elementos, se devuelve el valor del primer
	// elemento. Si está vacía entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// Devuelve el valor del último elemento. Si la lista tiene elementos, se devuelve el valor del último
	// elemento. Si está vacía entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Recorre la lista iterando elemento por elemento aplicando la función visitar en cada uno.
	// Se recorrerá la lista hasta que no haya más elementos (de primero a último) o hasta que
	// la función visitar devuelva false (lo que ocurra primero).
	Iterar(visitar func(T) bool)

	// Devuelve el iterador externo correspondiente al TDA.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// Devuelve el valor del elemento actual al que apunta el iterador. Si el iterador ya ha iterado
	// todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// Devuelve true si el elemento al que apunta el iterador tiene un elemento que le sigue.
	// En caso contrario devuelve false; es decir, estamos en el último elemento de la lista.
	HaySiguiente() bool

	// Avanza de posición en el iterador y devuelve el valor de ese nuevo elemento al que ahora apunta.
	// Si el iterador ya ha iterado todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Siguiente() T

	// Inserta un elemento en la posición actual a la que apunta el iterador de la lista. El elemento que ocupaba esa
	// posición (en caso de que haya) es reposicionado en la siguiente posición.
	Insertar(T)

	// Borra de la lista al elemento actual al que apunta el iterador y devuelve el valor de este. Al eliminarlo de la lista,
	// el elemento al que ahora apunta el iterador es el siguiente al que se acaba de eliminar.
	// Si el iterador ya ha iterado todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
