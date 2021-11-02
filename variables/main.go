package main

import "fmt"

func main() {

	//Asiganacion en dos lineas
	var dog string
	dog = "🐩"

	//Aisgnacion en una linea
	var cat string = "😽"

	// Crear varias variables en una sola linea, solo separadas por una "," y a la vez poder asignar un  valor con la separacion de una ","

	var horse, chicken string = "🐴", "🐔"

	fmt.Println(dog, cat)
	fmt.Println(horse, chicken)

	//Tambien se puede declarar una variable sin un tipo, el identifica el tipo de forma dinamica

	var apple = "🍎"
	fmt.Println(apple)

	//Si no se quiere especificar var y tipo se puede usar un shorthand como este :=, este solo sirve si se va a inicializar la variable con un valor, de lo contrario no se puede usar ya que no se crea otra variable.

	cow := "🐮"
	fmt.Println(cow)

	//Casos especiales:

	// Crear una variable y a la vez asignarle un valor en conjunto a una variable ya existente

	cow, car := "🐄", "🚗"
	fmt.Println(cow, car)

	// cow := "vaca" Esto es un error ya que la variable ya ha sido declarada entonces le shorhand no se puede volver a usar

	// NOTA: No se puede cambiar el tipo de dato despues de la declaracion y asignacion por primera vez

}
