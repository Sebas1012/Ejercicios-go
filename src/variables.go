package main

import "fmt"

func main() {
	//Manera de declarar variables en Go
	//Tipos de variable: int, string, bool, const
	var a string = "Inicial"
	fmt.Println(a)

	var b, c int = 4, 5
	fmt.Println(b + c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	//Forma Resumida Variables
	f := "Short"
	fmt.Println(f)

	g := 5 + 5
	fmt.Println(g)

	//Forma Resumida Constantes
	f = "Long"
	fmt.Println(f)

	g = 6 + 6
	fmt.Println(g)

}
