/*Crear tipos es crear tipos de datos propios para quizas terner un formato diferente a los tipos de datos
que ya nos da go, en este gaso solo cero un tipo de dato llamado moneda que es int  y le puedo aplicar un formato diferente que al int
que ya me da go*/

package main

import "fmt"

type moneda int

var b moneda

func main() {
	fmt.Println("El valor de b es: ", b)
	fmt.Printf("El tipo de b es: %T\n", b)

}
