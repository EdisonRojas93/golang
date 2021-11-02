package main

import "fmt"

//Los tipos de datos son los siguiente en GO
//bool -> Tipo de datos boobleano
//string ->  tipo de datos sintrg
//uint -> Tipo de dato numerico unicamente para numeros positivos (tiene rangos definidos)
// int -> tipo de dato numerico positivo y negativo (Tiene rangos definidos)
// byte ->  alias for uint8
// rune -> alias int32 or unicode
// float -> numeros decimales, tiene rangos definidos como el int

func main() {

	// var a bool = true
	// var b string = "prueba"

	// fmt.Printf("Tipo: %T, Valor: %v ", a, a)
	// fmt.Printf("Tipo: %T Valor: %v", b, b)

	// var c uint8 = 255
	// var d uint16 = 20

	// var e byte = 255
	// var f rune = 'a'
	// fmt.Printf("Tipo: %T, Valor %v", f, f)

	//NOTA: Cuando hacemos operaciones con diferentes tipos de de datos numericos puede generar error

	// var number1 int = -200
	// var number2 uint = 300

	// result := number1 + number2
	// aqui lo mejor es convertir al tipo que se vaya a esperar, lo ideal es castear

	// result := number1 - int(number2)

	// fmt.Print(result)

	// PRobando el tipo de dato black para declarar variables que no se usaran por el momento

	_ = 255
	result := _ + 1
	fmt.Print(result)

}
