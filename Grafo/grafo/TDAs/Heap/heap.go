package cola_prioridad

type heap[T comparable] struct {
	datos       []T
	cantidad    int
	funcion_cmp func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	const LARGO_INICIAL = 8
	heap_struct := new(heap[T])
	heap_struct.datos = make([]T, LARGO_INICIAL)
	heap_struct.funcion_cmp = funcion_cmp
	return heap_struct
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	const LARGO_INICIAL = 8
	var nuevo_arreglo = make([]T, len(arreglo))
	copy(nuevo_arreglo, arreglo)

	heapify(nuevo_arreglo, funcion_cmp)
	heap_struct := new(heap[T])
	heap_struct.funcion_cmp = funcion_cmp
	heap_struct.cantidad = len(nuevo_arreglo)

	if len(nuevo_arreglo) != 0 {
		heap_struct.datos = nuevo_arreglo
	} else {
		heap_struct.datos = make([]T, LARGO_INICIAL)
	}

	return heap_struct
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)

	ultRelativo := len(elementos) - 1

	for i := ultRelativo; i >= 0; i-- {
		elementos[i-ultRelativo], elementos[ultRelativo] = elementos[ultRelativo], elementos[i-ultRelativo]
		downheap(elementos, ultRelativo, i-ultRelativo, funcion_cmp)
		ultRelativo--
	}
}

func heapify[T comparable](arreglo []T, comparar func(T, T) int) {
	if len(arreglo) == 0 {
		return
	}

	inicio := len(arreglo) / 2

	for i := inicio; i >= 0; i-- {
		downheap(arreglo, len(arreglo), i, comparar)
	}
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == cap(heap.datos) {
		nueva_capacidad := (cap(heap.datos) * 2)
		heap.redimensionar(nueva_capacidad)
	}
	heap.datos[heap.cantidad] = elem
	upheap(heap.datos, heap.cantidad, heap.funcion_cmp)
	heap.cantidad++
}

func upheap[T comparable](arr []T, posHijo int, comparar func(T, T) int) {
	if posHijo == 0 {
		return
	}

	posPadre := calculaPosicionPadre(posHijo)
	comparacion := comparar(arr[posHijo], arr[posPadre])
	if comparacion > 0 {
		arr[posHijo], arr[posPadre] = arr[posPadre], arr[posHijo]
		upheap(arr, posPadre, comparar)
	}
}

func calculaPosicionPadre(posicionHijo int) int {
	posicionPadre := (posicionHijo - 1) / 2

	return posicionPadre
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	const PRIMERA_POSICION = 0

	dato := heap.datos[PRIMERA_POSICION]
	heap.datos[PRIMERA_POSICION], heap.datos[heap.cantidad-1] = heap.datos[heap.cantidad-1], heap.datos[PRIMERA_POSICION]
	downheap(heap.datos, heap.cantidad-1, PRIMERA_POSICION, heap.funcion_cmp)

	var valor_defecto T
	heap.datos[heap.cantidad-1] = valor_defecto
	heap.cantidad--

	if heap.cantidad != 0 && heap.cantidad <= (cap(heap.datos)/4) {
		nueva_capacidad := (cap(heap.datos) / 2)
		heap.redimensionar(nueva_capacidad)
	}

	return dato
}

// func (heap *heap[T]) downheap(posPadre int) {
func downheap[T comparable](arr []T, cantidad int, posPadre int, comparar func(T, T) int) {
	var (
		max         T
		posMax      int
		comparacion int
	)

	h_izq, h_der := calculaPosicionHijos(posPadre)

	if posPadre == cantidad-1 || h_izq >= cantidad {
		return
	}

	if h_der < cantidad {
		comparacion = comparar(arr[h_izq], arr[h_der])
	} else {
		comparacion = 1
	}

	if comparacion > 0 {
		posMax = h_izq
		max = arr[h_izq]
	} else {
		posMax = h_der
		max = arr[h_der]
	}

	comparacion = comparar(max, arr[posPadre])

	if comparacion > 0 {
		arr[posMax], arr[posPadre] = arr[posPadre], arr[posMax]
		downheap(arr, cantidad-1, posMax, comparar)
	}
}

func calculaPosicionHijos(posPadre int) (int, int) {
	h_izq := (2 * posPadre) + 1
	h_der := (2 * posPadre) + 2

	return h_izq, h_der
}

func (heap *heap[T]) redimensionar(nueva_capacidad int) {
	nuevos_datos := make([]T, nueva_capacidad)
	copy(nuevos_datos, heap.datos)
	heap.datos = nuevos_datos
}
