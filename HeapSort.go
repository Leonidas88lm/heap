package cola_prioridad

// Heapsort ordena un arreglo de acuerdo a la función de comparación.
// El orden resultante es ascendente (del menor al mayor) según cmp.

func HeapSort[T any](arreglo []T, cmp func(T, T) int) {

	h := &heap[T]{datos: arreglo, cmp: cmp}
	h.heapify()

	largo := len(arreglo)
	for i := largo - 1; i > 0; i-- {
		arreglo[0], arreglo[i] = arreglo[i], arreglo[0]
		h.datos = arreglo[:i]
		h.filtrarAbajo(0)
	}
}
