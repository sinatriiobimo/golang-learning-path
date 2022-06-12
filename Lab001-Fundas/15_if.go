package main

import "fmt"

func main() {
	name := "Sinatrio Bimo"

	if name == "Sinatrio Bimo" {
		fmt.Println("Hello Bimo")
	} else if name == "Ridwan Kamil" {
		fmt.Println("Hello RK")
	} else if name == "Budi Sudarsono" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
