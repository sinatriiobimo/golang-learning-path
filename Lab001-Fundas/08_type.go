package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpEko NoKTP = "1111111111111"
	var marriedStatusEko Married = true

	fmt.Println(ktpEko)
	fmt.Println(NoKTP("222222222222"))
	fmt.Println(marriedStatusEko)
}
