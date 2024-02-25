/*Los tipos numericos se componen segun la arquitectura de procesador que usemos, tambien podemos limitar el numero de bits que puede recibir
usando base2 ejemplo uint8 y su base2 seria uint16 el numero define la cantidad de bits que puede recibir, uint8 recibe de 0 a 255 y uint16 recibe 65535

Tambien muetro un ejemplo de como usar la libreria runtime para saber la arquitectura y el sistema operativo que usa la maquina que lo compila
GOOS = Sistema operativo
GOARCH = Arquitectura*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 12
	y := 12.54

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
