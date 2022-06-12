//?? Panic function adalah function yang bisa kita gunakan untuk menghentikan program
//?? Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
//?? Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Success: App Running")
}

func main() {
	runApp(true)
	fmt.Println("Panggil Main")
}
