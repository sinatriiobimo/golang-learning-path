//?? Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
//?? Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
//?? Method adalah function

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is", customer.Name)
}

func main() {
	rully := Customer{Name: "Rully"}
	rully.sayHello()
}
