package main

func main() {

	// Operadores basicos: (), *, +, -, %, /

	_ = 4 + 2

	// Operadores de asignacion: =, +=, -=, *=, /=, %=

	var fullName string = "Edison"

	fullName += " Rojas"
	//fmt.Println(fullName)

	var number int = 5
	number %= 2
	//fmt.Println(number)

	//Declaracion post-incremento y post-decremento: ++, --
	// (No son una expresión sino una declaracion)
	c := 3
	c++
	println(c)

}
