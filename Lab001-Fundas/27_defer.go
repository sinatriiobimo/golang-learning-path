//?? Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
//?? Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Running application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(0)
}
