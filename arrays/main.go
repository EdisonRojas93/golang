package main

import "fmt"

func main() {

	var foods [3]string
	foods[0] = "🍔"
	foods[1] = "🍕"
	foods[2] = "🌭"
	fmt.Println(foods)

	fruit := [3]string{"🍎", "🍌", "🍊"}
	fmt.Println(fruit)

	var transportation [3]string = [3]string{"🚗", "🛫", "🚲"}
	fmt.Println(transportation)

}
