package main

import "fmt"

func main() {
	// Estructuras permite guardar una collecion de campos

	type course struct {
		Name      string
		Professor string
		Country   string
	}

	var courses []course

	db := course{
		Name:      "Bases de datos",
		Professor: "Alexis",
		Country:   "Colombia",
	}

	courses = append(courses, db)
	fmt.Println(courses[0].Name)
	// fmt.Printf("%+v", db)
	// fmt.Println(courses)
}
