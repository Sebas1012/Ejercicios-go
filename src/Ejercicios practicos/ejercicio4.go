package main

import "fmt"

type an int

var x an

func main() {

	fmt.Println("El valor de x es:", x)
	fmt.Printf("El tipo de x es: %T\n", x)

	x = 42

	fmt.Println("El nuevo valor de x es:", x)

}
