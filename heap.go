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

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: []T{}, cmp: cmp}
}

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {

	datos := make([]T, len(arreglo))
	copy(datos, arreglo)
	h := &heap[T]{
		datos: datos,
		cmp:   cmp,
	}
	h.heapify()
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
	h.filtrarAbajo(0)
	h.reducirCapacidad()
	return max
}

func (h *heap[T]) heapify() {
	for i := len(h.datos)/2 - 1; i >= 0; i-- {
		h.filtrarAbajo(i)
	}
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

func (h *heap[T]) filtrarAbajo(pos int) {
	ultimo := len(h.datos) - 1
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos

		if hijoIzq <= ultimo && h.cmp(h.datos[hijoIzq], h.datos[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer <= ultimo && h.cmp(h.datos[hijoDer], h.datos[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			break
		}
		h.datos[pos], h.datos[mayor] = h.datos[mayor], h.datos[pos]
		pos = mayor
	}
}
