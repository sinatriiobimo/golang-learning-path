package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// fmt.Println(Pembagian(10, 10))
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
