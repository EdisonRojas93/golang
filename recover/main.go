package main

import "fmt"

/*
	Recover: Permite recuper despues de un panic para continuar la ejecucion del programa

*/

func main() {

	validateOld(20)
	validateOld(25)
	validateOld(50)
	validateOld(17)
	validateOld(29)

}

func validateOld(old int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ya me pasaron platica para que puedas pasar \n")
		}
	}()

	if old >= 18 {
		fmt.Println("Hola tu edad es aceptada")
	} else {
		panic("😬 ups! no puedes estar aqui")
	}
}
