package main

import "fmt"

func main() {

	//Estructura

	// switch expression {
	// case condition:
	// case condition:

	// default:

	// }

	emoji := "🌵"

	switch emoji {
	case "🐕":
		fmt.Println("Es un perro")

	case "🐈":
		fmt.Println("es un gato")
	case "🐎":
		fmt.Println("es un caballo")

	default:
		fmt.Printf("es: %s\n", emoji)
	}

}
