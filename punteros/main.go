package main

import "fmt"

//Los punteros: son referencia a memoria, esto permite gestionar o copiar informacion pensando en la gestion de la memoria, es decir en vez de copiar informacion almaceno la posicion en memoria y hago referencia a ella, pero si todas las referencias que hayan de ese espacio en memoria, el valor de esa referencia cambia, todas quedan con el nuevo valor.

func main() {

	fruit := "🍎"
	var puntero *string = &fruit
	var puntero2 **string = &puntero

	fmt.Printf("Tipo: %T, Valor: %v, Direccion: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferensacion %v \n", puntero, puntero, *puntero)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferensacion %v \n", puntero2, puntero2, *puntero2)

	fruit = "🍌"
	fmt.Printf("Tipo: %T, Valor: %v, Direccion: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferensacion %v \n", puntero, puntero, *puntero)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferensacion %v \n", puntero2, puntero2, *puntero2)

}
