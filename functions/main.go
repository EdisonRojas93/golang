package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// hello("edison", "Rojas")
	// emoji := "😽"
	// changeValue(emoji)
	// fmt.Println(emoji)

	// fruit := "🍌"
	// changeReference(&fruit)
	// fmt.Println(fruit)

	// fmt.Println(sum(1, 5))

	// lower, upper := convert("Edison eres lo mejor")
	// fmt.Println(lower, upper)

	// readFile()

	// result, err := dividir(10, 2)

	// if err != nil {
	// 	fmt.Print("error", err)
	// 	return
	// }

	// fmt.Println("dividiendo", result, err)

	// nums := []int{1, 5, 3, 54, 454, 2, 100}

	// result := filter(nums, func(value int) bool {
	// 	if value < 10 {
	// 		return true
	// 	}

	// 	return false
	// })

	// fmt.Println(result)

	// result := sum(10)
	// fmt.Println(result(2))

	fmt.Println(sumar(10, 2, 2, 2, 2))

	//Funciones anonimas

	x := func() {
		fmt.Println("Hola funciones anonimas")
	}
	x()

	// Funciones anonimas autoejecutadas

	func(name string) {
		fmt.Println("hello ", name)
	}("Edison")

}

// func hello(firstName string, lastName string) {

// 	fmt.Printf("hello %v %v", firstName, lastName)
// }

// //parametros por valor
// func changeValue(value string) {
// 	value = "🐶"
// }

// //Parametros por regerencia
// func changeReference(value *string) {
// 	*value = "🍎"
// }

//Funciones que retornar valores
// func sum(numberA, numberB int) int {
// 	return numberA + numberB
// }

// Funciones que retornar mas de un valor
// func convert(text string) (string, string) {
// 	return strings.ToLower(text), strings.ToUpper(text)
// }

//Control de errores

func readFile() {
	content, err := ioutil.ReadFile("./things.txt")

	if err != nil {
		fmt.Printf("Ocurrio un error %v", err)
	} else {
		fmt.Println(content)
	}

}

// func dividir(dividendo, divisor int) (result float32, err error) {

// 	if divisor == 0 {
// 		err = errors.New("Error, divisor no puede ser cero")
// 		return
// 	}

// 	result = float32(dividendo / divisor)
// 	return

// }

//Funciones que reciben funciones
func filter(numbers []int, callback func(int) bool) []int {

	result := []int{}

	for _, value := range numbers {
		if callback(value) {
			result = append(result, value)
		}
	}

	return result
}

//Funciones que retornan funciones
func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Funciones variaticas
func sumar(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}

	return result
}
