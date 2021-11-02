package main

import "fmt"

//MAPS: son estructuras de llave valor

func main() {

	animals := make(map[string]string)
	animals["cat"] = "🐱"
	animals["dog"] = "🐶"

	fmt.Println(animals)

	// Mapa literal

	fruits := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
	}

	fmt.Println(fruits)

	// delete(): sirve para eliminar elemento del mapa

	delete(fruits, "banana")
	fmt.Println(fruits)

	//Cuando se consulta un valor en una mapa, este puede devolver dos valores de datos

	if _, exists := animals["gorilla"]; !exists {
		animals["gorilla"] = "🦍"
	}

	fmt.Println(animals)

}
