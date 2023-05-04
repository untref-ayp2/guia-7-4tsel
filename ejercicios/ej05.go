package ejercicios

// import "ejercicios/stack"
// Escribir un método que calcule recursivamente cuántos
// elementos hay en una pila y no altere el contenido de
// la misma. La pila sólo tiene los métodos Push, Pop y isEmpty.
func CantidadDeElementos(pila Stack) int {

	if pila.IsEmpty(){

		return 0
	}

	objeto, _ := pila.Pop()
	tamano := CantidadDeElementos(pila) + 1
	
	pila.Push(objeto)

	return tamano
}
