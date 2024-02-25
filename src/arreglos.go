// Arreglos

// Declarar un array de 5 strings cada uno, los vamos a inicializar con 0.
// Declarar un segundo arreglo de 5 strings, los vamos a inicializar con los valores de strings.
// Asignar el segundo arreglo al primero y desplegar los resultados del primero.
// Mostrar el valor de la cadena y la dirección de cada elemento.

package main

import (
	"fmt"
	"runtime"
)

func main() {

	universidades := [3]string{"TdeA", "ITM", "UdeA"}

	fmt.Println(universidades)
	fmt.Println(runtime.GOOS)

}
