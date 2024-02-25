//Constantes

package main

//Forma de llamar varias librerias
//o esta
//import "x"
//import "y"

import (
	"fmt"
	"math"
)

const h string = "Gopher"

func main() {
	fmt.Println(h)

	const n = 500
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
