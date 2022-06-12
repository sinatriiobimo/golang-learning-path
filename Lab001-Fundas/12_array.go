package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Sinatrio"
	names[1] = "Bimo"
	names[2] = "Wahyudi"

	fmt.Println(names)

	var values = [3]int{10, 20, 30}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[2])
	values[2] = 100
	fmt.Println(values)
}
