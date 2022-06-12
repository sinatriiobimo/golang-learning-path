package main

import "fmt"

func main() {
	name := "Bimo"
	counter := 0

	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
