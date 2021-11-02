package course

import "fmt"

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

func (c Course) printClasses() {

	for _, class := range c.Classes {
		fmt.Println(class)
	}
}

func (c *Course) updatePrice(newPrise float64) {
	c.Price = newPrise
	fmt.Println("price update successfull")
}
