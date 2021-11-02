package main

import "fmt"

func main() {

	// For clasico
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// For Continuo
	// i := 0
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// For infinito

	// for {
	// 	println("Infinito papa")
	// }

	// For Range para maps, slices, y strings por defecto esta funcion regresa dos valores Index y value

	//numbers := []uint8{2, 4, 6, 8}
	//sport := map[string]string{"basketball": "🏀", "soccer": "⚽"}
	hello := "hello"

	// for index, value := range numbers {
	// 	fmt.Printf("Este es el index: %v, Este es el Value: %v \n", index, value)
	// }

	// for key, value := range sport {
	// 	fmt.Printf("Esta es la llave: %v, Este es el valor %v\n", key, value)
	// }

	for index, value := range hello {
		fmt.Printf("Este es el index: %v, Este es el Value: %v \n", index, string(value))
	}

}
