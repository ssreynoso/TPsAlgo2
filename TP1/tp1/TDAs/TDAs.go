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
