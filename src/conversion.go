/*Para hacer una conversion de un tipo de dato a otro solo necesitamos crear una variable del tipo al que queremos convertir y aplicar la siguente sintaxis
a=tipo(b) y asi se hara la conversion de tipo*/
package main

import "fmt"

func main() {
	var a float32
	var b int

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	a = 15.8
	b = 7

	fmt.Println(a, b)

	b = int(a)

	fmt.Println(b)

}
