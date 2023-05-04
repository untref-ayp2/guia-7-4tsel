package ejercicios

// Escriba un método recursivo para determinar si
// un elemento está en un arreglo utilizando búsqueda binaria.
func BusquedaBinaria(arreglo []int, elemento int) bool {

	medio := len(arreglo) / 2

	if arreglo[medio] == elemento {

		return true
	}

	if arreglo[medio] > elemento {

		return BusquedaBinaria(arreglo[:medio], elemento)
	}else if arreglo[medio] < elemento {

		return BusquedaBinaria(arreglo[medio:], elemento)
	}

	return false
}
