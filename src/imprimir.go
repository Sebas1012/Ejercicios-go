/* Se pueden usar3 formas para imprimir, pero segun la forma hace cosas diferentes
fmt.Print() = Imprime en una sola linea
fmt.Println() = Agrega un salto de linea automatico
fmt.Printf() = Permite imprimir usando caracteres de escape como \n y como requisito nos pide dar un formato al tipo que usamos
y segun esto debemos usar un verbo por ejemplo %v es el formato por defecto para cualquier tipo de dato*/
package main

import "fmt"

var a int
var b string = "TdeA"
var c bool

func main() {
	fmt.Println("El valor de a es:", a)
	fmt.Println("El valor de b es:", b)
	fmt.Println("El valor de c es:", c)

	fmt.Printf("El tipo de a es: %T\n", a)
	fmt.Printf("El tipo de b es: %T\n", b)
	fmt.Printf("El tipo de c es: %T\n", c)

}
