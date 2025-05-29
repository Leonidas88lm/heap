package cola_prioridad_test

import (
	"cmp"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const CANTIDAD_PRUEBA_VOLUMEN = 100000

func cmpBool(a, b bool) int {
	if a == b {
		return 0
	}
	if !a && b {
		return -1 // false < true
	}
	return 1
}

func TestColaPrioridadVacia(t *testing.T) {

	heapInt := TDAHeap.CrearHeap[int](cmp.Compare)
	require.True(t, heapInt.EstaVacia())
	require.Equal(t, 0, heapInt.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapInt.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapInt.Desencolar() })

	heapString := TDAHeap.CrearHeap[string](strings.Compare)
	require.True(t, heapString.EstaVacia())
	require.Equal(t, 0, heapString.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapString.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapString.Desencolar() })

	heapBool := TDAHeap.CrearHeap[bool](cmpBool)
	require.True(t, heapBool.EstaVacia())
	require.Equal(t, 0, heapBool.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapBool.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapBool.Desencolar() })
}

func TestColaPrioridadUnElemento(t *testing.T) {

	entero := 5
	cadena := "a"
	booleano := true

	heapInt := TDAHeap.CrearHeap[int](cmp.Compare)
	heapInt.Encolar(entero)
	require.Equal(t, entero, heapInt.VerMax())
	heapInt.Desencolar()
	require.True(t, heapInt.EstaVacia())
	require.Equal(t, 0, heapInt.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapInt.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapInt.Desencolar() })

	heapString := TDAHeap.CrearHeap[string](strings.Compare)
	heapString.Encolar(cadena)
	require.Equal(t, cadena, heapString.VerMax())
	heapString.Desencolar()
	require.True(t, heapString.EstaVacia())
	require.Equal(t, 0, heapString.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapString.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapString.Desencolar() })

	heapBool := TDAHeap.CrearHeap[bool](cmpBool)
	heapBool.Encolar(booleano)
	require.Equal(t, booleano, heapBool.VerMax())
	heapBool.Desencolar()
	require.True(t, heapBool.EstaVacia())
	require.Equal(t, 0, heapBool.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapBool.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapBool.Desencolar() })

}

func TestColaPrioridadEncolarVariosElementosEntero(t *testing.T) {
	arregloEntero := []int{1, 2, 3, 4, 6, 8, 5, 9}           //arreglo que voy a usar para agregar uno a uno al heap
	arregloEnteroMax := []int{1, 2, 3, 4, 6, 8, 8, 9}        // arreglo de lo que tiene que devolver VerMax despues de cada encolar
	arregloEnteroDesencolar := []int{9, 8, 6, 5, 4, 3, 2, 1} // arreglo de como tienen que salir los elementos al desencolar
	heapInt := TDAHeap.CrearHeap[int](cmp.Compare)
	for i := 0; i < len(arregloEntero); i++ {
		heapInt.Encolar(arregloEntero[i])
		require.Equal(t, arregloEnteroMax[i], heapInt.VerMax())
	}
	for i := 0; i < len(arregloEntero); i++ {
		require.Equal(t, arregloEnteroDesencolar[i], heapInt.Desencolar())
	}
}

func TestColaPrioridadEncolarVariosElementosString(t *testing.T) {
	arregloString := []string{"1", "2", "3", "4", "6", "8", "5", "9"}           //arreglo que voy a usar para agregar uno a uno al heap
	arregloStringMax := []string{"1", "2", "3", "4", "6", "8", "8", "9"}        // arreglo de lo que tiene que devolver VerMax despues de cada encolar
	arregloStringDesencolar := []string{"9", "8", "6", "5", "4", "3", "2", "1"} // arreglo de como tienen que salir los elementos al desencolar
	heapString := TDAHeap.CrearHeap[string](cmp.Compare)
	for i := 0; i < len(arregloString); i++ {
		heapString.Encolar(arregloString[i])
		require.Equal(t, arregloStringMax[i], heapString.VerMax())
	}
	for i := 0; i < len(arregloString); i++ {
		require.Equal(t, arregloStringDesencolar[i], heapString.Desencolar())
	}
}

func TestColaPrioridadDesdeArregloEntero(t *testing.T) {
	arregloEntero := []int{1, 2, 3, 4, 6, 8, 5, 9}
	arregloEnteroDesencolar := []int{9, 8, 6, 5, 4, 3, 2, 1}
	heap := TDAHeap.CrearHeapArr(arregloEntero, cmp.Compare)
	for i := 0; i < len(arregloEntero); i++ {
		require.Equal(t, arregloEnteroDesencolar[i], heap.Desencolar())
	}
}

func TestColaPrioridadDesdeArregloString(t *testing.T) {
	arregloString := []string{"1", "2", "3", "4", "6", "8", "5", "9"}
	arregloStringDesencolar := []string{"9", "8", "6", "5", "4", "3", "2", "1"}
	heap := TDAHeap.CrearHeapArr(arregloString, strings.Compare)
	for i := 0; i < len(arregloString); i++ {
		require.Equal(t, arregloStringDesencolar[i], heap.Desencolar())
	}
}

func TestColaPrioridadVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmp.Compare)
	for i := 0; i < CANTIDAD_PRUEBA_VOLUMEN; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}
