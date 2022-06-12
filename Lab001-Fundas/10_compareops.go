package main

import "fmt"

func main() {
	var name1 = "Bimo"
	var name2 = "Bambang"

	var resultEqual = name1 == name2
	var resultGreater = name1 > name2

	fmt.Println(resultEqual)
	fmt.Println(resultGreater)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
