package main

import (
	"fmt"
	"os"
)

/**

Defer: Permite aplazar una instrucion que se ejecuta dentro de una funcion hasta que finalice dicha funcion o retorne un valor

- El defer funciona como una pila que ejecuta desde el ultimo elemento que ingreso, hasta el primero

- Cuando se envia a defer una expresion, los valores se alamacenan como se evaluaron en su momento


USOS:

1. Para cerrar conexiones
2.

*/

func main() {

	// defer fmt.Println("3")
	// defer fmt.Println("1")
	// fmt.Println("2")

	// fmt.Println(sum(10,2))

	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("Ocurrio un error al crear", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello Edison"))

	if err != nil {
		fmt.Println("Ocurrio un error al escribir", err)
		return
	}
}

func sum(a, b int) int {

	var result int
	// defer result = total * 3
	// total := a + b

	return result
}
