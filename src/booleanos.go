/*Los tipos booleanos solo tienen dos valores constantes false y true y su tipo se define como bool*/

package main

import "fmt"

var x bool

func main() {

	fmt.Println(x)
	x = true
	fmt.Println(x)

	y := 9
	z := 7

	fmt.Println(y != z)

}
