package main

import "fmt"

/*

	Panic, permite finalizar el programa, cuando es llamado finaliza el programa

*/
func main() {
	division(12, 2)
	division(100, 4)
	division(20, 2)
	division(12, 0)
	division(50, 3)
}

func division(dividendo, divisor int) {
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("😱")
	}
}
