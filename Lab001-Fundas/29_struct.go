/**
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field
Sederhananya struct adalah kumpulan dari field
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var bimo Customer
	bimo.Name = "Sinatrio Bimo"
	bimo.Address = "Jakarta"
	bimo.Age = 21

	fmt.Println(bimo)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)
}
