package main

import "fmt"

func main() {
	x := "Sebastian"
	y := `Sebastian
	henao`

	fmt.Println(x)
	fmt.Println(y)

	z := []byte(x)

	fmt.Println(z)
}
