package main

import "fmt"

func main() {

	// los slice son apuntadores del array no obtiene una nueva referencia
	// set := [7]string{"🐯", "🐴", "🐮", "🦋", "🐦", "🛫", "🛬"}
	// _ = set[0:5]
	//fly := set[3:] // --> si no se especifica [init:final] final, go detecta que debe ir hasta el ultimo cuando se usa [:] llama todos los elementos del array

	//fly[0] = "🦅"

	// fmt.Println(set)
	// fmt.Println(animals)
	// fmt.Println(fly)

	// len(): # de elementos que contiene el slice
	// cap(): # de elementos del array apuntador apartir del indice

	// food := [6]string{"🌭", "🍓", "🍋", "🍔", "🍕", "🍎"}
	// fruits := food[1:3]
	// fruits = append(fruits, "🍍", "🍈", "🍉", "🥝")

	// clone := food
	// clone[4] = "🍐"

	// fmt.Println(food)
	// fmt.Println(clone)
	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	//Slices dinamicos de tamaño
	// fruits := []string{"🍓", "🍋"}
	// fruits = append(fruits, "🍎")

	// fmt.Println("Frutas: ", fruits)
	// fmt.Println("Tamaño: ", len(fruits))
	// fmt.Println("Capacidad: ", cap(fruits))

	// Puede crearse slice previamente definidos sin contenido

	// fruits := make([]string, 0, 1)
	// fruits = append(fruits, "🍍", "🍈", "🍉", "🥝", "🍎")

	// fmt.Println("Frutas: ", fruits)
	// fmt.Println("Tamaño: ", len(fruits))
	// fmt.Println("Capacidad: ", cap(fruits))

	//Slice con capacidad de cero

	var fruits []string //esto esta definido pero no inicializado es decir es == nil
	fruits = append(fruits, "🍍", "🍈", "🍉", "🥝", "🍎")

	fmt.Println("Frutas: ", fruits)
	fmt.Println("Tamaño: ", len(fruits))
	fmt.Println("Capacidad: ", cap(fruits))

	//food := []string{} // Esto esta definido e inicializado es decir no es nil

}
