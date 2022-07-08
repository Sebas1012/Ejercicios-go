package main

import "fmt"

func main() {
	/*En Go hay diferentes formas de crear variables, donde unas son mas explicitas que otras a la hora de definir el tipo de dato que se
	almacenara en la variale.

	Dos maneras en que se pueden definir variables son:

	var nombre = expresion
	var nombre tipo = expresion

	*/

	var edad_1 = 22
	var edad_2 int = 21

	fmt.Println(edad_1)
	fmt.Println(edad_2)

	/*Tambien es importante tener en cuenta que se pueden generar declaraciones de multiples variables de manera "resumida" de las siguientes maneras

	var nombre_1, nombre_2, nombre_3 tipo
	var(
		nombre_1 tipo
		nombre_2 tipo
		nombre_3 tipo
	)
	var(
		nombre_1 = expresion
		nombre_2 = expresion
		nombre_3 = expresion
	)
	*/

	var a, b, c int
	var (
		nombre string
		edad_3 int
		acceso bool
	)
	var (
		genero   = "M"
		acceso_2 = true
	)

	fmt.Println(a, b, c)
	fmt.Println(nombre, edad_3, acceso)
	fmt.Println(genero, acceso_2)

	/*Por ultimo tenemos la asignacion corta de variables, la cual nos permite ahorrarnos el var y el tipo ya que automaticamente interpreta el tipo de dato.
	Esta declaracion tiene una particularidad, y es que solo puede ser usada dentro de funciones y no fuera ya que de esa manera generara un error.

	La asignacion se puede hacer de la siguiente manera:

	Individual:

	nombre := expresion

	Multiple:

	nombre_1, nombre_2, nombre_3 := expresion_1, expresion_2, expresion_3
	*/

	color := "Rojo"
	peso, valor := 12.5, 300

	fmt.Println(color, peso, valor)
}
