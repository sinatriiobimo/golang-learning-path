//?? Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
//?? Sebuah interface berisikan definisi-definisi method
//?? Biasanya interface digunakan sebagai kontrak

package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	person := Person{Name: "Bimo"}
	animal := Animal{Name: "Push"}
	sayHello(person)
	sayHello(animal)
}
