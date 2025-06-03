package cola_prioridad

// ColaPrioridad ya definida en cola_prioridad.go

const _FACTOR_REDUCCION = 4
const _FACTOR_CRECIMIENTO = 2
const _CAPACIDAD_INICIAL = 8

// heap representa un heap binario máximo.
type heap[T any] struct {
	datos []T
	cmp   func(T, T) int
}

// panicHeap genera los panic cuando la cola esta vacía y no debería estarlo
func (h *heap[T]) panicHeap() {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (h *heap[T]) reducirCapacidad() {

	if len(h.datos) <= cap(h.datos)/_FACTOR_REDUCCION && cap(h.datos) > _CAPACIDAD_INICIAL {
		nuevaCapacidad := max(len(h.datos), _CAPACIDAD_INICIAL)
		nuevo := make([]T, len(h.datos), nuevaCapacidad)
		copy(nuevo, h.datos)
		h.datos = nuevo
	}
}

func (h *heap[T]) aumentarCapacidad() {

	if len(h.datos) == cap(h.datos) {
		nuevaCapacidad := max(cap(h.datos)*_FACTOR_CRECIMIENTO, _CAPACIDAD_INICIAL)
		nuevo := make([]T, len(h.datos), nuevaCapacidad)
		copy(nuevo, h.datos)
		h.datos = nuevo
	}
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		filtrarAbajo(arr, i, cmp)
	}
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: []T{}, cmp: cmp}
}

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {

	datos := make([]T, len(arreglo))
	copy(datos, arreglo)
	heapify(datos, cmp)
	h := &heap[T]{
		datos: datos,
		cmp:   cmp,
	}
	return h
}

func (h *heap[T]) EstaVacia() bool {
	return len(h.datos) == 0
}

func (h *heap[T]) Cantidad() int {
	return len(h.datos)
}

func (h *heap[T]) VerMax() T {
	h.panicHeap()
	return h.datos[0]
}

func (h *heap[T]) Encolar(elem T) {
	h.aumentarCapacidad()
	h.datos = append(h.datos, elem)
	h.filtrarArriba(len(h.datos) - 1)
}

func (h *heap[T]) Desencolar() T {
	h.panicHeap()
	max := h.datos[0]
	ultimo := len(h.datos) - 1
	h.datos[0] = h.datos[ultimo]
	h.datos = h.datos[:ultimo]
	filtrarAbajo(h.datos, 0, h.cmp)
	h.reducirCapacidad()
	return max
}

func (h *heap[T]) filtrarArriba(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if h.cmp(h.datos[pos], h.datos[padre]) <= 0 {
			break
		}
		h.datos[pos], h.datos[padre] = h.datos[padre], h.datos[pos]
		pos = padre
	}
}

func filtrarAbajo[T any](arr []T, pos int, cmp func(T, T) int) {
	ultimo := len(arr) - 1
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos

		if hijoIzq <= ultimo && cmp(arr[hijoIzq], arr[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer <= ultimo && cmp(arr[hijoDer], arr[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			break
		}
		arr[pos], arr[mayor] = arr[mayor], arr[pos]
		pos = mayor
	}
}

// Heapsort ordena un arreglo de acuerdo a la función de comparación.
// El orden resultante es ascendente (del menor al mayor) según cmp.

func HeapSort[T any](arreglo []T, cmp func(T, T) int) {

	heapify(arreglo, cmp)
	h := &heap[T]{datos: arreglo, cmp: cmp}

	largo := len(arreglo)
	for i := largo - 1; i > 0; i-- {
		arreglo[0], arreglo[i] = arreglo[i], arreglo[0]
		h.datos = arreglo[:i]
		filtrarAbajo(h.datos, 0, h.cmp)
	}
}
