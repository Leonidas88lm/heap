package cola_prioridad

// Heapsort ordena un arreglo de acuerdo a la función de comparación.
// El orden resultante es ascendente (del menor al mayor) según cmp.

func heapSort[T any](arreglo []T, cmp func(T, T) int) {

	heapify(arreglo, cmp)
	h := &heap[T]{datos: arreglo, cmp: cmp}

	largo := len(arreglo)
	for i := largo - 1; i > 0; i-- {
		arreglo[0], arreglo[i] = arreglo[i], arreglo[0]
		h.datos = arreglo[:i]
		filtrarAbajo(h.datos, 0, h.cmp)
	}
}
