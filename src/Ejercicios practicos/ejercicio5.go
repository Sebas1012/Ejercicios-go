package main

import "fmt"

type numero int

var x numero
var y int

func main() {
	fmt.Println("El valor de x es:", x)
	fmt.Printf("El tipo de x es: %T\n", x)

	x = 42

	fmt.Println("El nuevo valor de x es:", x)

	y = int(x)

	fmt.Println("El valor de y es:", y)
	fmt.Printf("El tipo de y es: %T\n", y)

}
