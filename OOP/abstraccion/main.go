package main

import "fmt"

// Abstraccion por medio de estructuras
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	Users   []User
	Classes map[uint]string
}

type User struct {
	Name    string
	Country string
}

func main() {

	classes := map[uint]string{
		1: "Arrays",
		2: "Loops",
		3: "Conditionals",
	}

	Go := Course{
		Name:    "Edison",
		Price:   60000,
		IsFree:  false,
		Users:   []User{},
		Classes: classes,
	}

	fmt.Println(Go.Name)
	Go.PrintClasses()
	Go.UpdatePrice(130000)
	fmt.Println(Go)
}

//Agregar metodos a nuestra abstaccion

//function normal
// func printClasses(c Course) {

// 	for _, class := range c.Classes {
// 		fmt.Println(class)
// 	}
// }

// function como metodo de una estructura
/*

Es necesario cambiar la estructura para que se pueda utilizar una funcion como metodo

func (receptor estructura)nombreMetodo(no es necesario ya que se definio el receptor ){
	contenido
}
*/

func (c Course) PrintClasses() {

	for _, class := range c.Classes {
		fmt.Println(class)
	}
}

func (c *Course) UpdatePrice(newPrise float64) {
	c.Price = newPrise
	fmt.Println("price update successfull")
}
