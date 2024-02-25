package main

import "fmt"

func main() {
	//Primera forma de hacer un ciclo
	a := 1
	for a <= 3 {
		fmt.Println(a)
		a = a + 1
	}
	//Segunda forma de hacer un ciclo
	for c := 7; c <= 9; c++ {
		fmt.Println(c)
	}

}
